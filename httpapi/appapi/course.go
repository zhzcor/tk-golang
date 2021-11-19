package appapi

import (
	"context"
	"database/sql"
	"encoding/json"
	sql2 "entgo.io/ent/dialect/sql"
	"github.com/gin-gonic/gin"
	"gserver/httpapi/appapi/request"
	"gserver/httpapi/appapi/response"
	"gserver/internal/app"
	"gserver/internal/cache"
	"gserver/internal/config"
	"gserver/internal/errorno"
	"gserver/internal/store"
	"gserver/internal/store/ent"
	"gserver/internal/store/ent/advertise"
	"gserver/internal/store/ent/kcclass"
	"gserver/internal/store/ent/kccourse"
	"gserver/internal/store/ent/kccoursesmallcategory"
	"gserver/internal/store/ent/kccourseteacher"
	"gserver/internal/store/ent/kcsmallcategoryattachment"
	"gserver/internal/store/ent/kcsmallcategoryexampaper"
	"gserver/internal/store/ent/kcsmallcategoryquestion"
	"gserver/internal/store/ent/kcuserclass"
	"gserver/internal/store/ent/kcusercourse"
	"gserver/internal/store/ent/message"
	"gserver/internal/store/ent/teacher"
	"gserver/internal/store/ent/tkexampaper"
	"gserver/internal/store/ent/tkuserexamscorerecord"
	"gserver/internal/store/ent/tkuserquestionbankrecord"
	"gserver/internal/store/ent/useraskanswer"
	"gserver/internal/store/ent/usercourseappraise"
	"gserver/internal/store/ent/videorecord"
	"gserver/pkg/baijiayun"
	app2 "gserver/pkg/requestparam"
	"math"
	"sort"
	"strconv"
	"strings"
	"time"
)

//1：视频，2：音频，3：直播，4：资料下载
const (
	VideoSmall        = 1
	AudioSmall        = 2
	LiveSmall         = 3
	DataDownloadSmall = 4
)

//1:普通课程，2:直播课程，3:直播公开课，4:录播公开课
const (
	GeneralCourses    = 1
	LiveCourses       = 2
	LiveOpenCourses   = 3
	LiveRecordCourses = 4
)

const (
	Live         = 1 //直播
	LivePlayBack = 2 //回放
	LiveDemand   = 3 //点播
)

//"直播状态1.未开始，2:直播中，3：直播结束，4：回放生成中，5：已生成回放，6回放生成失败，7：已上传录像"
const (
	LiveNotStart            = 1
	LiveIng                 = 2
	LiveEnd                 = 3
	LivePlayBackIng         = 4
	LivePlayBackGenerated   = 5
	LivePlayBackBuildFailed = 6
	LiveUploaded            = 7
)

//直播课程日历 已经购买的课程
func CourseCalendar(ctx *gin.Context) (interface{}, error) {
	ca := cache.QuestionCache
	uid, _ := ctx.Get("uid")
	smsKey := cache.UserCalendarKey + strconv.Itoa(uid.(int))
	if data, ok := ca.Get(smsKey); ok {
		return data, nil
	}
	s := store.WithContext(ctx)
	com := app.Common{}

	date := time.Now().Format("2006-01-02")
	t, _ := time.ParseInLocation("2006-01-02", date, time.Local)

	/*	start, end := com.GetMonthStartEnd(t)
	 */
	/*	dateCount := com.MonthCount(t.Year(), int(t.Month()))
	 */
	//获取周结束时间
	start1, end := com.WeekDayDiff(t)

	end1 := end.Unix() + 86399

	end = time.Unix(end1, 0)

	start := com.GetZeroTime(start1)

	/*endTimeTmp := t.Unix() + 86399
	endTime := time.Unix(endTimeTmp, 0)*/
	start2 := com.GetZeroTime(t)

	dateList := com.GetBetweenDates(start1, end)

	//获取班级报名当月直播课程
	ClassInfoList := s.KcUserClass.Query().Where(kcuserclass.UserID(uid.(int))).
		WithClass(func(query *ent.KcClassQuery) {
			query.SoftDelete().Select("id", "class_title", "class_cover_img_id").WithAttachment(func(query *ent.AttachmentQuery) {
				query.SoftDelete().Select("id", "filename")
			}).WithKcClassCourses(func(query2 *ent.KcCourseQuery) {
				query2.SoftDelete().Select("id", "course_type", "course_name", "question_bank_id", "course_cover_img_id").WithAttachment(func(query *ent.AttachmentQuery) {
					query.SoftDelete().Select("id", "filename")
				}).WithCourseTeachers(func(query *ent.KcCourseTeacherQuery) {
					query.WithTeacher(func(query *ent.TeacherQuery) {
						query.SoftDelete().WithAttachment(func(query *ent.AttachmentQuery) {
							query.SoftDelete().Select("id", "filename")
						})
					})
				}).WithQuestionBank(func(query *ent.TkQuestionBankQuery) {
					query.SoftDelete().Select("id").WithExamPapers(func(query *ent.TkExamPaperQuery) {
						query.SoftDelete().Select("id", "name", "exam_question_type", "exam_type", "score", "duration", "start_at", "end_at", "question_bank_id").Where(tkexampaper.EnableStatus(1), tkexampaper.ExamQuestionType(1), tkexampaper.StartAtGTE(start), tkexampaper.EndAtLTE(end)).WithUserExamPapers(func(query2 *ent.TkUserExamScoreRecordQuery) {
							query2.Where(tkuserexamscorerecord.UserID(uid.(int)))
						})
					}).WithUserQuestionBank(func(query *ent.TkUserQuestionBankRecordQuery) {
						query.SoftDelete().Select("record_count", "question_bank_id").Where(tkuserquestionbankrecord.UserID(uid.(int)))
					})
				}).WithCourseSmallCategorys(func(query3 *ent.KcCourseSmallCategoryQuery) {
					query3.SoftDelete().Where(kccoursesmallcategory.LiveSmallStartGTE(start), kccoursesmallcategory.LiveSmallStartLTE(end))
				})
			})
		}).
		AllX(ctx)

	//获取单个课程报名的直播课
	courseUser := s.KcUserCourse.Query().Where(kcusercourse.UserID(uid.(int))).WithCourse(func(query *ent.KcCourseQuery) {
		query.SoftDelete().Select("id", "course_type", "course_name", "question_bank_id", "course_cover_img_id").WithAttachment(func(query *ent.AttachmentQuery) {
			query.SoftDelete().Select("id", "filename")
		}).WithCourseTeachers(func(query *ent.KcCourseTeacherQuery) {
			query.WithTeacher(func(query *ent.TeacherQuery) {
				query.SoftDelete().WithAttachment(func(query *ent.AttachmentQuery) {
					query.SoftDelete().Select("id", "filename")
				})
			})
		}).WithQuestionBank(func(query *ent.TkQuestionBankQuery) {
			query.Select("id").SoftDelete().WithExamPapers(func(query *ent.TkExamPaperQuery) {
				query.SoftDelete().Select("id", "name", "exam_question_type", "exam_type", "score", "duration", "start_at", "end_at", "question_bank_id").Where(tkexampaper.EnableStatus(1), tkexampaper.ExamQuestionType(1), tkexampaper.StartAtGTE(start), tkexampaper.EndAtLTE(end)).
					WithUserExamPapers(func(query2 *ent.TkUserExamScoreRecordQuery) {
						query2.Where(tkuserexamscorerecord.UserID(uid.(int)))
					})
			}).WithUserQuestionBank(func(query *ent.TkUserQuestionBankRecordQuery) {
				query.Select("record_count", "question_bank_id").SoftDelete().Where(tkuserquestionbankrecord.UserID(uid.(int)))
			})
		}).WithAttachment(func(query *ent.AttachmentQuery) {
			query.SoftDelete().Select("id", "filename")
		}).WithCourseSmallCategorys(func(query *ent.KcCourseSmallCategoryQuery) {
			query.SoftDelete().Where(kccoursesmallcategory.LiveSmallStartGTE(start), kccoursesmallcategory.LiveSmallStartLTE(end))
		})
	}).AllX(ctx)

	var courese = map[string][]response.CourseCalendar{}
	var examList = map[string][]response.CourseExam{}

	var res = []response.CourseCalendar{}
	var resExamCalendar = []response.CourseCalendar{}

	var userCourseList = []response.UserCourseList{}

	//模拟考试和试卷
	if len(ClassInfoList) > 0 || len(courseUser) > 0 {
		resLive := []response.GetDateLiveSmallCourse{}
		resExam := []response.ExamQuestionTypeList{}
		Now := time.Now().Unix()
		a := app.Common{}
		c := app.Course{}
		var examPaperIds []int
		var courseSmallIds []int
		for _, v := range ClassInfoList {
			uc := response.UserCourseList{}
			uc.Name = v.Edges.Class.ClassTitle
			uc.ClassId = v.Edges.Class.ID
			if v.Edges.Class.Edges.Attachment != nil {
				uc.Img = app2.GetOssHost() + v.Edges.Class.Edges.Attachment.Filename
			}
			if len(v.Edges.Class.Edges.KcClassCourses) > 0 {
				uc.CourseCount = len(v.Edges.Class.Edges.KcClassCourses)
				for _, v1 := range v.Edges.Class.Edges.KcClassCourses {
					if v1.Edges.QuestionBank != nil {
						if len(v1.Edges.QuestionBank.Edges.UserQuestionBank) > 0 {
							count := v1.Edges.QuestionBank.Edges.UserQuestionBank[0].RecordCount
							uc.QuestionCount = uc.QuestionCount + count
						}

					}
					//是否有试卷
					if v1.Edges.QuestionBank != nil {
						if v1.Edges.QuestionBank.Edges.ExamPapers != nil {
							for _, v3 := range v1.Edges.QuestionBank.Edges.ExamPapers {
								ex := response.CourseExam{}
								if !v3.StartAt.IsZero() {
									ex.Date = v3.StartAt.Format("2006-01-02")
								}
								ex.ExamId = v3.ID
								examList[ex.Date] = append(examList[ex.Date], ex)

								//试卷详情
								m := &response.ExamQuestionTypeList{}
								m.Id = v3.ID
								m.CourseName = v1.CourseName
								m.BankId = v3.QuestionBankID
								m.TypeId = int(v3.ExamQuestionType)
								m.ExamUserStudyCount = v3.AnsweredUserCount
								m.ExamName = v3.Name
								m.QuestionCount = v3.QuestionCount
								m.PracticeExamAt = v3.Duration
								m.Score = v3.Score
								m.ExamType = int(v3.ExamType)
								m.PassScore = v3.PassScore
								m.Difficulty = a.GetDifficultyMap(int(v3.Difficulty))
								m.Desc = v3.Desc
								if len(v3.Edges.UserExamPapers) > 0 {
									op := v3.Edges.UserExamPapers[0].WrongCount + v3.Edges.UserExamPapers[0].RightCount
									if op > 0 {
										m.Accuracy, _ = strconv.ParseFloat(strconv.FormatFloat(float64(v3.Edges.UserExamPapers[0].RightCount*100)/float64(op), 'f', 1, 64), 64)
									}
									m.UserQuestionCount = v3.Edges.UserExamPapers[0].TotalCount
									m.OrderStatus = int(v3.Edges.UserExamPapers[0].OrderStatus)
									m.IsDo = 1
								}
								if !v3.StartAt.IsZero() {
									m.Time = v3.StartAt.Format("1月2日") + " " + v3.StartAt.Format("15:04") + "-" + v3.EndAt.Format("15:04")
									m.StartAt2 = v3.StartAt
								}
								//	ExamStart      = 1 //考试开始
								//	ExamEnd        = 2 //考试结束
								//	ExamViewResult = 3 //查看成绩
								//	ExamNot        = 4 //考试未开始
								if Now >= v3.StartAt.Unix() && Now <= v3.EndAt.Unix() {
									m.SimulationExamStatus = ExamStart
								} else if Now < v3.StartAt.Unix() {
									m.SimulationExamStatus = ExamNot
								} else if Now > v3.EndAt.Unix() {
									m.SimulationExamStatus = ExamEnd
								}

								if len(v3.Edges.UserExamPapers) > 0 {
									m.SimulationExamStatus = ExamViewResult
								}
								if !app2.IsContainInt(examPaperIds, v3.ID) {
									examPaperIds = append(examPaperIds, v3.ID)

									if v3.StartAt.After(start2) {
										resExam = append(resExam, *m)
									}
								}

							}
						}
					}

					if v1.CourseType == LiveCourses || v1.CourseType == LiveOpenCourses { //直播课才记录
						if len(v1.Edges.CourseSmallCategorys) > 0 {
							for _, v2 := range v1.Edges.CourseSmallCategorys {
								cr := response.CourseCalendar{}
								if !v2.LiveSmallStart.IsZero() {
									cr.Date = v2.LiveSmallStart.Format("2006-01-02")
								}
								courese[cr.Date] = append(courese[cr.Date], cr)

								//直播小节信息
								re := response.GetDateLiveSmallCourse{}
								re.LiveName = v2.SmallName
								re.LiveFinishType = int(v2.FinishType)

								if v2.LiveSmallTime > 0 {
									end := v2.LiveSmallStart.Unix() + int64(v2.LiveSmallTime)
									ends := time.Unix(end, 0)
									re.LiveStarAt = v2.LiveSmallStart.Format("01月02日 15:04") + "-" + ends.Format("15:04")
								}

								re.LiveStarAt2 = v2.LiveSmallStart.Format("1月2日")
								re.LiveStarAt3 = v2.LiveSmallStart
								if v1.Edges.Attachment != nil {
									re.CourseImg = app2.GetOssHost() + v1.Edges.Attachment.Filename
								}

								re.LiveSmallId = v2.ID
								status := c.CheckLiveType(int(v1.CourseType), int(v2.LiveSmallStatus), v2.LiveRoomID)
								re.LiveStatus = status.LiveStatus
								if len(v1.Edges.CourseTeachers) > 0 {
									if v1.Edges.CourseTeachers[0].Edges.Teacher != nil {
										if v1.Edges.CourseTeachers[0].Edges.Teacher != nil {
											re.TeacherName = v1.Edges.CourseTeachers[0].Edges.Teacher.Name
										}
									}

									if v1.Edges.CourseTeachers[0].Edges.Teacher.Edges.Attachment != nil {
										if v1.Edges.CourseTeachers[0].Edges.Teacher.Edges.Attachment.Filename != "" {
											re.TeacherAvatar = app2.GetOssHost() + v1.Edges.CourseTeachers[0].Edges.Teacher.Edges.Attachment.Filename
										}
									}
								}

								if !app2.IsContainInt(courseSmallIds, v2.ID) {
									courseSmallIds = append(courseSmallIds, v2.ID)
									if v2.LiveSmallStart.After(start2) {
										resLive = append(resLive, re)
									}
								}
							}
						}
					}

				}

			}

			userCourseList = append(userCourseList, uc)
		}

		for _, v := range courseUser {
			uc := response.UserCourseList{}
			if v.Edges.Course != nil {
				uc.Name = v.Edges.Course.CourseName
				uc.CourseId = v.Edges.Course.ID
				uc.CourseCount = uc.CourseCount + 1
				if v.Edges.Course.Edges.Attachment != nil {
					uc.Img = app2.GetOssHost() + v.Edges.Course.Edges.Attachment.Filename
				}
				if v.Edges.Course.Edges.QuestionBank != nil {
					if len(v.Edges.Course.Edges.QuestionBank.Edges.UserQuestionBank) > 0 {
						count := v.Edges.Course.Edges.QuestionBank.Edges.UserQuestionBank[0].RecordCount
						uc.QuestionCount = uc.QuestionCount + count
					}
				}

				userCourseList = append(userCourseList, uc)
				//是否有试卷
				if v.Edges.Course.Edges.QuestionBank != nil {
					if v.Edges.Course.Edges.QuestionBank.Edges.ExamPapers != nil {
						for _, v3 := range v.Edges.Course.Edges.QuestionBank.Edges.ExamPapers {
							ex := response.CourseExam{}
							if !v3.StartAt.IsZero() {
								ex.Date = v3.StartAt.Format("2006-01-02")
							}
							ex.ExamId = v3.ID
							examList[ex.Date] = append(examList[ex.Date], ex)

							m := &response.ExamQuestionTypeList{}
							m.Id = v3.ID
							m.BankId = v3.QuestionBankID
							m.CourseName = v.Edges.Course.CourseName
							m.TypeId = int(v3.ExamQuestionType)
							m.ExamUserStudyCount = v3.AnsweredUserCount
							m.ExamName = v3.Name
							m.QuestionCount = v3.QuestionCount
							m.PracticeExamAt = v3.Duration
							m.Score = v3.Score
							m.ExamType = int(v3.ExamType)
							m.PassScore = v3.PassScore
							m.Difficulty = a.GetDifficultyMap(int(v3.Difficulty))
							m.Desc = v3.Desc
							if len(v3.Edges.UserExamPapers) > 0 {
								op := v3.Edges.UserExamPapers[0].WrongCount + v3.Edges.UserExamPapers[0].RightCount
								if op > 0 {
									m.Accuracy, _ = strconv.ParseFloat(strconv.FormatFloat(float64(v3.Edges.UserExamPapers[0].RightCount*100)/float64(op), 'f', 1, 64), 64)
								}
								m.UserQuestionCount = v3.Edges.UserExamPapers[0].TotalCount
								m.OrderStatus = int(v3.Edges.UserExamPapers[0].OrderStatus)
								m.IsDo = 1
							}
							if !v3.StartAt.IsZero() {
								m.Time = v3.StartAt.Format("1月2日") + " " + v3.StartAt.Format("15:04") + "-" + v3.EndAt.Format("15:04")
								m.StartAt2 = v3.StartAt
							}
							//	ExamStart      = 1 //考试开始
							//	ExamEnd        = 2 //考试结束
							//	ExamViewResult = 3 //查看成绩
							//	ExamNot        = 4 //考试未开始
							if Now >= v3.StartAt.Unix() && Now <= v3.EndAt.Unix() {
								m.SimulationExamStatus = ExamStart
							} else if Now < v3.StartAt.Unix() {
								m.SimulationExamStatus = ExamNot
							} else if Now > v3.EndAt.Unix() {
								m.SimulationExamStatus = ExamEnd
							}

							if len(v3.Edges.UserExamPapers) > 0 {
								m.SimulationExamStatus = ExamViewResult

								m.IsOrder = int(v3.Edges.UserExamPapers[0].OrderStatus)
							}

							if !app2.IsContainInt(examPaperIds, v3.ID) {
								examPaperIds = append(examPaperIds, v3.ID)

								if v3.StartAt.After(start2) {
									resExam = append(resExam, *m)
								}
							}
						}
					}
				}

				if v.Edges.Course.CourseType == LiveCourses || v.Edges.Course.CourseType == LiveOpenCourses { //直播课才记录
					if len(v.Edges.Course.Edges.CourseSmallCategorys) > 0 {
						for _, v1 := range v.Edges.Course.Edges.CourseSmallCategorys {
							cr := response.CourseCalendar{}
							if !v1.LiveSmallStart.IsZero() {
								cr.Date = v1.LiveSmallStart.Format("2006-01-02")
							}
							courese[cr.Date] = append(courese[cr.Date], cr)
							//直播小节信息
							re := response.GetDateLiveSmallCourse{}
							re.LiveName = v1.SmallName
							re.LiveFinishType = int(v1.FinishType)
							if v1.LiveSmallTime > 0 {
								end := v1.LiveSmallStart.Unix() + int64(v1.LiveSmallTime)
								ends := time.Unix(end, 0)
								re.LiveStarAt = v1.LiveSmallStart.Format("01月02日 15:04") + "-" + ends.Format("15:04")
							}
							re.LiveStarAt2 = v1.LiveSmallStart.Format("1月2日")
							re.LiveStarAt3 = v1.LiveSmallStart
							re.LiveSmallId = v1.ID
							status := c.CheckLiveType(int(v.Edges.Course.CourseType), int(v1.LiveSmallStatus), v1.LiveRoomID)
							re.LiveStatus = status.LiveStatus
							if v.Edges.Course.Edges.Attachment != nil {
								re.CourseImg = app2.GetOssHost() + v.Edges.Course.Edges.Attachment.Filename
							}

							if len(v.Edges.Course.Edges.CourseTeachers) > 0 {
								if v.Edges.Course.Edges.CourseTeachers[0].Edges.Teacher != nil {
									re.TeacherName = v.Edges.Course.Edges.CourseTeachers[0].Edges.Teacher.Name

									if v.Edges.Course.Edges.CourseTeachers[0].Edges.Teacher.Edges.Attachment != nil {
										if v.Edges.Course.Edges.CourseTeachers[0].Edges.Teacher.Edges.Attachment.Filename != "" {
											re.TeacherAvatar = app2.GetOssHost() + v.Edges.Course.Edges.CourseTeachers[0].Edges.Teacher.Edges.Attachment.Filename
										}
									}
								}
							}

							if !app2.IsContainInt(courseSmallIds, v1.ID) {
								courseSmallIds = append(courseSmallIds, v1.ID)

								if v1.LiveSmallStart.After(start2) {
									resLive = append(resLive, re)
								}
							}
						}
					}
				}
			}
		}
		for _, v := range dateList {
			re := response.CourseCalendar{}
			rx := response.CourseCalendar{}
			re.Date = v
			rx.Date = v
			if time.Now().Format("2006-01-02") == v {
				re.Today = 1
				rx.Today = 1
			}
			if _, ok := courese[v]; ok && len(courese[v]) > 0 {
				re.Ok = 1
			}

			if _, ok := examList[v]; ok && len(examList[v]) > 0 {
				rx.Ok = 1
			}
			resExamCalendar = append(resExamCalendar, rx)
			res = append(res, re)
		}

		sort.Slice(resLive, func(i, j int) bool {
			return resLive[i].LiveStarAt3.Before(resLive[j].LiveStarAt3) // 升序
		})

		sort.Slice(resExam, func(i, j int) bool {
			return resExam[i].StartAt2.Before(resExam[j].StartAt2) // 升序
		})
		var data = map[string]interface{}{
			"calendar_course_list": res,
			"calendar_exam_list":   resExamCalendar,
			"live_list":            resLive,
			"exam_list":            resExam,
			"my_course_list":       userCourseList,
		}

		ca.Set(smsKey, data)
		return data, nil

	}

	return nil, nil
}

//日历详情
func GetCalendarDetail(ctx *gin.Context) (interface{}, error) {
	var req request.CourseCalendar
	err := ctx.Bind(&req)
	if err != nil {
		return nil, errorno.NewParamErr(err)
	}
	uid, _ := ctx.Get("uid")

	var t time.Time
	if req.Date == "" {
		date := time.Now().Format("2006-01-02")
		t, _ = time.ParseInLocation("2006-01-02", date, time.Local)
	} else {
		t, _ = time.Parse("2006-01-02", req.Date)
	}

	s := store.WithContext(ctx)
	com := app.Common{}
	start, end := com.GetMonthStartEnd(t)

	dateList := com.GetBetweenDates(start, end)

	userClassCoruse := s.KcUserClass.Query().Select("class_id").Where(kcuserclass.UserID(uid.(int))).WithClass(func(query *ent.KcClassQuery) {
		query.SoftDelete().Select("id").WithKcClassCourses(func(query *ent.KcCourseQuery) {
			query.SoftDelete().Select("id", "course_type", "course_name").Where(kccourse.PushStatus(1)).Where(kccourse.CourseTypeIn(2, 3))
		})
	}).AllX(ctx)

	userCoruse := s.KcUserCourse.Query().Where(kcusercourse.UserID(uid.(int))).WithCourse(func(query *ent.KcCourseQuery) {
		query.SoftDelete().Select("id", "course_type", "course_name").Where(kccourse.PushStatus(1)).Where(kccourse.CourseTypeIn(2, 3))
	}).AllX(ctx)

	coureIds := []int{}

	for _, v := range userClassCoruse {
		for _, v1 := range v.Edges.Class.Edges.KcClassCourses {
			coureIds = append(coureIds, v1.ID)
		}
	}

	for _, v := range userCoruse {
		coureIds = append(coureIds, v.CourseID)
	}

	query := s.KcCourse.Query().SoftDelete().Select("id", "question_bank_id").Where(kccourse.IDIn(coureIds...))
	if req.Type == 1 { //课程
		query = query.WithCourseSmallCategorys(func(query *ent.KcCourseSmallCategoryQuery) {
			query.SoftDelete().Where(kccoursesmallcategory.LiveSmallStartGTE(start), kccoursesmallcategory.LiveSmallStartLT(end))
		})
	} else { //模拟考试
		query = query.WithQuestionBank(func(query *ent.TkQuestionBankQuery) {
			query.Select("id").SoftDelete().WithExamPapers(func(query *ent.TkExamPaperQuery) {
				query.SoftDelete().Select("id", "name", "exam_question_type", "exam_type", "score", "duration", "start_at", "end_at", "question_bank_id").Where(tkexampaper.ExamQuestionType(1), tkexampaper.StartAtGTE(start), tkexampaper.EndAtLTE(end))
			})
		})
	}

	list := query.AllX(ctx)

	var courese = map[string][]response.CourseCalendar{}
	var Exam = map[string][]response.CourseExam{}
	var res = []response.CourseCalendar{}

	if req.Type == 1 {
		for _, v := range list {
			if v.Edges.CourseSmallCategorys != nil {
				for _, v1 := range v.Edges.CourseSmallCategorys {
					re := response.CourseCalendar{}
					re.Date = v1.LiveSmallStart.Format("2006-01-02")
					courese[re.Date] = append(courese[re.Date], re)
				}
			}
		}
	} else if req.Type == 2 {
		for _, v := range list {
			if v.Edges.QuestionBank != nil {
				if v.Edges.QuestionBank.Edges.ExamPapers != nil {
					for _, v3 := range v.Edges.QuestionBank.Edges.ExamPapers {
						ex := response.CourseExam{}
						if !v3.StartAt.IsZero() {
							ex.Date = v3.StartAt.Format("2006-01-02")
						}
						ex.ExamId = v3.ID
						Exam[ex.Date] = append(Exam[ex.Date], ex)
					}
				}
			}
		}
	}

	for _, v := range dateList {
		re := response.CourseCalendar{}
		re.Date = v
		if time.Now().Format("2006-01-02") == v {
			re.Today = 1
		}
		if req.Type == 1 {
			if _, ok := courese[v]; ok && len(courese[v]) > 0 {
				re.Ok = 1
			}
		} else {
			if _, ok := Exam[v]; ok && len(Exam[v]) > 0 {
				re.Ok = 1
			}
		}
		res = append(res, re)
	}

	return res, nil
}

//获取指定日期的直播直播课程和模拟考试
func GetUserLiveCourse(ctx *gin.Context) (interface{}, error) {
	var req request.CourseCalendar
	err := ctx.Bind(&req)

	if err != nil {
		return nil, errorno.NewParamErr(err)
	}
	uid, _ := ctx.Get("uid")
	ca := cache.CommCache
	Key := cache.UserLiveCourseDateKey + strconv.Itoa(uid.(int)) + ":" + req.Date + "type:" + strconv.Itoa(req.Type)
	if data, ok := ca.Get(Key); ok {
		return data, nil
	}

	s := store.WithContext(ctx)
	com := app.Common{}

	var t time.Time
	if req.Date == "" {
		t = time.Now()
	} else {
		t, _ = time.Parse("2006-01-02", req.Date)
	}

	start, end := com.GetDateStartEnd(t)

	//查询用户报名的直播课
	userClassCoruse, err := s.KcUserClass.Query().Select("class_id").Where(kcuserclass.UserID(uid.(int))).WithClass(func(query *ent.KcClassQuery) {
		query.SoftDelete().Select("id").WithKcClassCourses(func(query *ent.KcCourseQuery) {
			query.SoftDelete().Select("id", "course_type").Where(kccourse.PushStatus(1)).Where(kccourse.CourseTypeIn(2, 3))
		})
	}).All(ctx)

	userCoruse, err := s.KcUserCourse.Query().Where(kcusercourse.UserID(uid.(int))).WithCourse(func(query *ent.KcCourseQuery) {
		query.SoftDelete().Select("id", "course_type").Where(kccourse.PushStatus(1)).Where(kccourse.CourseTypeIn(2, 3))
	}).All(ctx)

	if err != nil {
		return nil, errorno.NewInternalErr(err)
	}

	coureIds := []int{}

	for _, v := range userClassCoruse {
		for _, v1 := range v.Edges.Class.Edges.KcClassCourses {
			coureIds = append(coureIds, v1.ID)
		}
	}

	for _, v := range userCoruse {
		coureIds = append(coureIds, v.CourseID)
	}

	query := s.KcCourse.Query().SoftDelete().Where(kccourse.IDIn(coureIds...)).WithAttachment(func(query *ent.AttachmentQuery) {
		query.Select("id", "filename").SoftDelete()
	})
	if req.Type == 1 { //课程
		query = query.WithCourseSmallCategorys(func(query *ent.KcCourseSmallCategoryQuery) {
			query.SoftDelete().Where(kccoursesmallcategory.LiveSmallStartGTE(start), kccoursesmallcategory.LiveSmallStartLT(end))
		}).WithCourseTeachers(func(query1 *ent.KcCourseTeacherQuery) {
			query1.Select("teacher_id", "course_id").WithTeacher(func(query2 *ent.TeacherQuery) {
				query2.Select("avatar_id", "name").SoftDelete().WithAttachment(func(query3 *ent.AttachmentQuery) {
					query3.Select("id", "filename").SoftDelete()
				})
			})
		})

	} else {
		query = query.WithQuestionBank(func(query *ent.TkQuestionBankQuery) {
			query.Select("id").SoftDelete().WithExamPapers(func(query *ent.TkExamPaperQuery) {
				query.SoftDelete().Select("id", "name", "exam_question_type", "exam_type", "score", "duration", "start_at", "end_at", "question_bank_id").Where(tkexampaper.ExamQuestionType(1), tkexampaper.StartAtGTE(start), tkexampaper.EndAtLTE(end)).
					WithUserExamPapers(func(query2 *ent.TkUserExamScoreRecordQuery) {
						query2.Where(tkuserexamscorerecord.UserID(uid.(int)))
					})
			}).WithUserQuestionBank(func(query *ent.TkUserQuestionBankRecordQuery) {
				query.Select("record_count", "question_bank_id").SoftDelete().Where(tkuserquestionbankrecord.UserID(uid.(int)))
			})
		})

	}

	list := query.AllX(ctx)
	c := app.Course{}
	if req.Type == 1 { //直播课程
		res := []response.GetDateLiveSmallCourse{}
		for _, v1 := range list {
			for _, v := range v1.Edges.CourseSmallCategorys {
				re := response.GetDateLiveSmallCourse{}
				re.LiveName = v.SmallName
				re.LiveFinishType = int(v.FinishType)
				if v1.Edges.Attachment != nil {
					re.CourseImg = app2.GetOssHost() + v1.Edges.Attachment.Filename
				}
				if v.LiveSmallTime > 0 {
					end := v.LiveSmallStart.Unix() + int64(v.LiveSmallTime)
					ends := time.Unix(end, 0)
					re.LiveStarAt = v.LiveSmallStart.Format("01月02日 15:04") + "-" + ends.Format("15:04")
				}

				/*				re.LiveStarAt = v.LiveSmallStart.Format("15:04")
				 */re.LiveStarAt2 = v.LiveSmallStart.Format("1月2日")
				re.LiveStarAt3 = v.LiveSmallStart
				re.LiveSmallId = v.ID
				status := c.CheckLiveType(int(v1.CourseType), int(v.LiveSmallStatus), v.LiveRoomID)
				re.LiveStatus = status.LiveStatus
				if len(v1.Edges.CourseTeachers) > 0 {
					if v1.Edges.CourseTeachers[0].Edges.Teacher != nil {
						re.TeacherName = v1.Edges.CourseTeachers[0].Edges.Teacher.Name

						if v1.Edges.CourseTeachers[0].Edges.Teacher.Edges.Attachment != nil {
							re.TeacherAvatar = app2.GetOssHost() + v1.Edges.CourseTeachers[0].Edges.Teacher.Edges.Attachment.Filename
						}
					}

				}
				res = append(res, re)
			}
		}
		if len(res) > 0 {
			sort.Slice(res, func(i, j int) bool {
				return res[i].LiveStarAt3.Before(res[j].LiveStarAt3) // 降序
			})
			ca.Set(Key, res)

			return res, nil
		}
	} else {
		Now := time.Now().Unix()
		a := app.Common{}

		res := []response.ExamQuestionTypeList{}
		for _, v := range list {
			if v.Edges.QuestionBank != nil {
				if v.Edges.QuestionBank.Edges.ExamPapers != nil {
					for _, v1 := range v.Edges.QuestionBank.Edges.ExamPapers {
						m := &response.ExamQuestionTypeList{}
						m.Id = v1.ID
						m.CourseName = v.CourseName
						m.BankId = v1.QuestionBankID
						m.TypeId = int(v1.ExamQuestionType)
						m.ExamUserStudyCount = v1.AnsweredUserCount
						m.ExamName = v1.Name
						m.QuestionCount = v1.QuestionCount
						m.PracticeExamAt = v1.Duration
						m.Score = v1.Score
						m.ExamType = int(v1.ExamType)
						m.PassScore = v1.PassScore
						m.Difficulty = a.GetDifficultyMap(int(v1.Difficulty))
						m.Desc = v1.Desc
						if len(v1.Edges.UserExamPapers) > 0 {
							op := v1.Edges.UserExamPapers[0].WrongCount + v1.Edges.UserExamPapers[0].RightCount
							if op > 0 {
								m.Accuracy, _ = strconv.ParseFloat(strconv.FormatFloat(float64(v1.Edges.UserExamPapers[0].RightCount*100)/float64(op), 'f', 1, 64), 64)
							}
							m.UserQuestionCount = v1.Edges.UserExamPapers[0].TotalCount
							m.OrderStatus = int(v1.Edges.UserExamPapers[0].OrderStatus)
							m.IsDo = 1
						}
						if !v1.StartAt.IsZero() {
							m.Time = v1.StartAt.Format("1月2日") + " " + v1.StartAt.Format("15:04") + "-" + v1.EndAt.Format("15:04")
							m.StartAt2 = v1.StartAt
						}
						//	ExamStart      = 1 //考试开始
						//	ExamEnd        = 2 //考试结束
						//	ExamViewResult = 3 //查看成绩
						//	ExamNot        = 4 //考试未开始
						if Now >= v1.StartAt.Unix() && Now <= v1.EndAt.Unix() {
							m.SimulationExamStatus = ExamStart
						} else if Now < v1.StartAt.Unix() {
							m.SimulationExamStatus = ExamNot
						} else if Now > v1.EndAt.Unix() {
							m.SimulationExamStatus = ExamEnd
						}

						if len(v1.Edges.UserExamPapers) > 0 {
							m.SimulationExamStatus = ExamViewResult
						}

						if len(v1.Edges.UserExamPapers) > 0 {
							m.IsOrder = int(v1.Edges.UserExamPapers[0].OrderStatus)
						}
						res = append(res, *m)
					}
				}
			}
		}

		if len(res) > 0 {
			sort.Slice(res, func(i, j int) bool {
				return res[i].StartAt2.Before(res[j].StartAt2) // 降序
			})
			ca.Set(Key, res)

			return res, nil
		}
	}
	ca.Set(Key, nil)
	return nil, nil
}

//获取班级下课程
func GetClassCourseList(ctx *gin.Context) (interface{}, error) {
	var req request.ClassCourseList
	err := ctx.Bind(&req)
	if err != nil {
		return nil, errorno.NewParamErr(err)
	}
	s := store.WithContext(ctx)
	ca := cache.CommCache
	Key := cache.ClassCourseListKey + strconv.Itoa(req.ClassId)
	if data, ok := ca.Get(Key); ok {
		return data, nil
	}

	res := response.ClassCourseList{
		CourseList: []response.CourseList{},
	}

	list := s.KcClass.Query().SoftDelete().Where(kcclass.ID(req.ClassId), kcclass.IsDisplay(1)).WithKcClassCourses(func(query *ent.KcCourseQuery) {
		query.SoftDelete().Where(kccourse.PushStatus(1)).WithAttachment(func(query *ent.AttachmentQuery) {
			query.SoftDelete().Select("id", "filename")
		}).WithCourseTeachers(func(query *ent.KcCourseTeacherQuery) {
			query.Where(kccourseteacher.TeacherIDNotNil()).WithTeacher(func(query *ent.TeacherQuery) {
				query.SoftDelete().WithAttachment(func(query *ent.AttachmentQuery) {
					query.SoftDelete().Select("id", "filename")
				})
			})
		}).WithCourseChapters(func(query *ent.KcCourseChapterQuery) {
			query.WithChapterSections()
		}).WithCourseSmallCategorys(func(query *ent.KcCourseSmallCategoryQuery) {
			query.SoftDelete().Select("id","course_id")
		})
	}).AllX(ctx)

	if len(list) > 0 {
		for _, v := range list {
			res.ClassName = v.ClassTitle
			res.ClassId = v.ID
			res.ClassCourseCount = v.CourseCount
			for _, v1 := range v.Edges.KcClassCourses {
				re := response.CourseList{}
				re.CourseName = v1.CourseName
				re.CourseId = v1.ID
				re.SmallCourseCount = len(v1.Edges.CourseSmallCategorys)
				for _, v2 := range v1.Edges.CourseChapters {
					re.ChapterCount = re.ChapterCount + 1
					for _ = range v2.Edges.ChapterSections {
						re.SecCount = re.SecCount + 1
					}
				}

				if v1.Edges.Attachment != nil {
					re.ImgCover = app2.GetOssHost() + v1.Edges.Attachment.Filename
				}

				if len(v1.Edges.CourseTeachers) > 0 {
					if v1.Edges.CourseTeachers[0].Edges.Teacher != nil {
						re.TeacherName = v1.Edges.CourseTeachers[0].Edges.Teacher.Name

						if v1.Edges.CourseTeachers[0].Edges.Teacher.Edges.Attachment != nil {
							re.TeacherAvatar = app2.GetOssHost() + v1.Edges.CourseTeachers[0].Edges.Teacher.Edges.Attachment.Filename
						}
					}
				}

				res.CourseList = append(res.CourseList, re)
			}
		}
		ca.Set(Key, res)

		return res, nil
	}
	ca.Set(Key, nil)

	return nil, nil
}

//获取课程详情目录课时
func GetCourseDetails(ctx *gin.Context) (interface{}, error) {
	var req request.CourseDetailsById
	err := ctx.Bind(&req)
	if err != nil {
		return nil, errorno.NewParamErr(err)
	}

	ca := cache.CommCache
	Key := cache.CourseChapterKey + strconv.Itoa(req.CourseId)
	if data, ok := ca.Get(Key); ok {
		return data, nil
	}

	c := app.Course{}

	list, err := c.GetCourseChapterList(ctx, req.CourseId)

	if err != nil {

		return nil, errorno.NewInternalErr(err)
	}

	ca.Set(Key, list)
	return list, nil

}

//获取课时为资料的列表
func GetCourseFile(ctx *gin.Context) (interface{}, error) {
	var req request.CourseDetailsById
	err := ctx.Bind(&req)
	if err != nil {
		return nil, errorno.NewParamErr(err)
	}
	c := app.Course{}

	data, err := c.GetCourseFileList(ctx, req.CourseId)

	if err != nil {
		return nil, errorno.NewInternalErr(err)
	}

	return data, nil
}

//小节课程详情
func GetCateSmallDetail(ctx *gin.Context) (interface{}, error) {
	var req request.CourseSmallDetail
	err := ctx.Bind(&req)
	if err != nil {
		return nil, errorno.NewParamErr(err)
	}
	s := store.WithContext(ctx)

	smallInfo, err := s.KcCourseSmallCategory.Query().SoftDelete().Where(kccoursesmallcategory.ID(req.SmallId)).WithCourse(func(query *ent.KcCourseQuery) {
		query.SoftDelete().Select("id", "course_type", "course_name").WithCourseTeachers(func(query *ent.KcCourseTeacherQuery) {
			query.Select("id", "teacher_id", "course_id").Where(kccourseteacher.ShowStatus(1)).WithTeacher(func(query *ent.TeacherQuery) {
				query.SoftDelete().Select("id", "avatar_id", "name").WithAttachment(func(query *ent.AttachmentQuery) {
					query.SoftDelete().Select("filename", "id")
				})
			})
		})
	}).First(ctx)

	if err != nil {
		if ent.IsNotFound(err) {
			return nil, nil
		}
		return nil, errorno.NewInternalErr(err)
	}
	uid, _ := ctx.Get("uid")
	if uid == nil{
		uid = 0
	}

	bjy := baijiayun.BjyCloud{}
	//设置视频转码回调地址
	var smallDetail response.SmallDetail
	//获取token
	switch smallInfo.Edges.Course.CourseType {
	case LiveCourses, LiveOpenCourses: //直播小节课程
		//判断直播状态
		var statusArr = []int{
			LivePlayBackGenerated, //已上传录像
			LiveUploaded,          //已生成回放
		}
		if smallInfo.LiveSmallStatus == LiveNotStart {
			// 1.未开始，2:直播中，3：直播结束，4：回放
			smallDetail.LiveStatus = 1
		}

		if smallInfo.LiveSmallStatus == LiveEnd {
			smallDetail.LiveStatus = 3
		}
		smallDetail.SmallType = Live

		if app2.IsContainInt(statusArr, int(smallInfo.LiveSmallStatus)) { //是否是回放
			if smallInfo.LiveRoomID > 0 {
				data, err := bjy.GetPlayBackToken(smallInfo.LiveRoomID)
				smallDetail.SmallType = 1
				if err != nil {
					return nil, err
				}
				// 1.未开始，2:直播中，3：直播结束，4：回放
				smallDetail.LiveStatus = 4
				smallDetail.SmallType = LivePlayBack
				smallDetail.Token = data.Token
				smallDetail.VideoId = data.VideoId
			}
		}

		//直播中获取学生参加码
		if smallInfo.LiveSmallStatus == LiveIng {
			Code, err := bjy.GetStudentCode(smallInfo.LiveRoomID, uid.(int), "")
			if err != nil {
				return nil, err
			}
			// 1.未开始，2:直播中，3：直播结束，4：回放
			smallDetail.LiveStatus = 2
			smallDetail.StudentCode = Code.StudentCode
		}
	case LiveRecordCourses, GeneralCourses: //普通课程 和录播
		smallDetail.SmallType = LiveDemand
		if smallInfo.OrderVideoID > 0 {
			data, err := bjy.GetOnDemandToken(smallInfo.OrderVideoID) //获取点播token
			if err != nil {
				return nil, err
			}
			smallDetail.Token = data.Token
			smallDetail.VideoId = smallInfo.OrderVideoID
		}
	}

	//记录观看记录
	if uid !=0{
		query := s.VideoRecord.Query().Where(videorecord.UserID(uid.(int)), videorecord.SmallID(req.SmallId))
		data := query.FirstX(ctx)
		if data != nil {
			s.VideoRecord.UpdateOneID(data.ID).SetViewAt(time.Now()).SaveX(ctx)
		} else {
			s.VideoRecord.Create().SetUserID(uid.(int)).SetViewAt(time.Now()).SetSmallID(req.SmallId).SaveX(ctx)
		}
	}

	smallDetail.Id = smallInfo.ID
	smallDetail.CourseType = int(smallInfo.Edges.Course.CourseType)
	if len(smallInfo.Edges.Course.Edges.CourseTeachers) > 0 {
		if smallInfo.Edges.Course.Edges.CourseTeachers[0].Edges.Teacher != nil {
			smallDetail.TeacherName = smallInfo.Edges.Course.Edges.CourseTeachers[0].Edges.Teacher.Name
		}

		if smallInfo.Edges.Course.Edges.CourseTeachers[0].Edges.Teacher != nil {
			if smallInfo.Edges.Course.Edges.CourseTeachers[0].Edges.Teacher.Edges.Attachment != nil {
				if smallInfo.Edges.Course.Edges.CourseTeachers[0].Edges.Teacher.Edges.Attachment.Filename != "" {
					smallDetail.Avatar = app2.GetOssHost() + smallInfo.Edges.Course.Edges.CourseTeachers[0].Edges.Teacher.Edges.Attachment.Filename
				}
			}
		}
	}

	smallDetail.DomainUrl = config.ServerConfig.Baijiayun.NoHttpDomain
	smallDetail.FileType = int(smallInfo.Type)
	/*	smallDetail.KcCourseSmallCategory.OrderVideoID = smallInfo.OrderVideoID
	 */
	smallDetail.PartnerKey = config.ServerConfig.Baijiayun.PartnerKey
	smallDetail.LiveRoomID = smallInfo.LiveRoomID
	smallDetail.ChapterID = smallInfo.ChapterID
	smallDetail.SmallName = smallInfo.SmallName
	smallDetail.SectionID = smallInfo.SectionID
	smallDetail.CourseID = smallInfo.CourseID
	smallDetail.CourseType = int(smallInfo.Edges.Course.CourseType)
	smallDetail.AttachmentCount = smallInfo.AttachmentCount
	smallDetail.LiveSmallStart = smallInfo.LiveSmallStart.Format("2006-01-02 15:04:05")
	smallDetail.LiveSmallTime = smallInfo.LiveSmallTime
	smallDetail.LiveSmallStatus = int(smallInfo.LiveSmallStatus)
	smallDetail.QuestionBankId = smallInfo.Edges.Course.QuestionBankID

	return smallDetail, nil
}

//获取随堂练习
func GetPracticeInClass(ctx *gin.Context) (interface{}, error) {
	var req request.CourseSmallDetail
	err := ctx.Bind(&req)
	if err != nil {
		return nil, errorno.NewParamErr(err)
	}
	res := response.SmallCourseExam{
		ExamList:      []response.QuestionBankExam{},
		LianxiList:    []int{},
		WorkList:      []response.QuestionBankExam{},
		MaterialsList: []response.SmallMaterials{},
	}
	uid, _ := ctx.Get("uid")
	s := store.WithContext(ctx)
	//试卷
	exma := s.KcSmallCategoryExamPaper.Query().SoftDelete().Where(kcsmallcategoryexampaper.SmallCategoryID(req.SmallId)).WithExamPaper(func(query *ent.TkExamPaperQuery) {
		query.SoftDelete().Where(tkexampaper.EnableStatus(1)).WithUserExamPapers(func(query *ent.TkUserExamScoreRecordQuery) {
			query.SoftDelete().Where(tkuserexamscorerecord.UserID(uid.(int)))
		})
	}).AllX(ctx)
	//练习
	quesInfo := s.KcSmallCategoryQuestion.Query().SoftDelete().Where(kcsmallcategoryquestion.SmallCategoryID(req.SmallId)).AllX(ctx)
	//资料
	zilioa := s.KcSmallCategoryAttachment.Query().SoftDelete().Where(kcsmallcategoryattachment.SmallCategoryID(req.SmallId)).WithAttachment(func(query *ent.AttachmentQuery) {
		query.SoftDelete()
	}).AllX(ctx)
	a := app.Common{}

	for _, v := range exma {
		m := &response.QuestionBankExam{}
		if v.Edges.ExamPaper != nil {
			m.Id = v.Edges.ExamPaper.ID
			m.BankId = v.Edges.ExamPaper.QuestionBankID
			m.TypeId = int(v.Edges.ExamPaper.ExamQuestionType)
			m.ExamUserStudyCount = v.Edges.ExamPaper.AnsweredUserCount
			m.ExamName = v.Edges.ExamPaper.Name
			m.QuestionCount = v.Edges.ExamPaper.QuestionCount
			m.PracticeExamAt = v.Edges.ExamPaper.Duration
			m.Score = v.Edges.ExamPaper.Score
			m.ExamType = int(v.Edges.ExamPaper.ExamType)
			m.PassScore = v.Edges.ExamPaper.PassScore
			m.Difficulty = a.GetDifficultyMap(int(v.Edges.ExamPaper.Difficulty))
			m.Desc = v.Edges.ExamPaper.Desc
			if len(v.Edges.ExamPaper.Edges.UserExamPapers) > 0 {
				op := v.Edges.ExamPaper.Edges.UserExamPapers[0].WrongCount + v.Edges.ExamPaper.Edges.UserExamPapers[0].RightCount
				if op > 0 {
					m.Accuracy, _ = strconv.ParseFloat(strconv.FormatFloat(float64(v.Edges.ExamPaper.Edges.UserExamPapers[0].RightCount*100)/float64(op), 'f', 1, 64), 64)
				}

				m.OrderStatus = int(v.Edges.ExamPaper.Edges.UserExamPapers[0].OrderStatus)
				m.IsDo = 1
			}

			// 分类，1：试卷，2：作业
			if v.Type == 1 {
				res.ExamList = append(res.ExamList, *m)
			} else {
				res.WorkList = append(res.WorkList, *m)
			}
		}
	}

	for _, v := range quesInfo {
		res.LianxiList = append(res.LianxiList, v.QuestionID)
	}

	for _, v := range zilioa {
		m := response.SmallMaterials{}
		if v.Edges.Attachment != nil {
			m.UserID = v.Edges.Attachment.UserID
			m.Filename = app2.GetOssHost() + v.Edges.Attachment.Filename
			m.FileSize = v.Edges.Attachment.FileSize
			m.MimeType = v.Edges.Attachment.MimeType
			res.MaterialsList = append(res.MaterialsList, m)
		}

	}

	return res, nil

}

//获取老师列表
func GetTeacherList(ctx *gin.Context) (interface{}, error) {
	ca := cache.CommCache
	Key := cache.TeacherListKey
	if data, ok := ca.Get(Key); ok {
		return data, nil
	}
	uid, _ := ctx.Get("uid")
	//获取已报名课程老师
	s := store.WithContext(ctx)
	classList := s.KcUserClass.Query().Where(kcuserclass.UserID(uid.(int))).WithClass(func(query *ent.KcClassQuery) {
		query.SoftDelete().Where(kcclass.Status(1)).Select("id", "class_head_master_id").WithMasterTeachers(func(query *ent.TeacherQuery) {
			query.SoftDelete()
		}).WithKcClassCourses(func(query *ent.KcCourseQuery) {
			query.SoftDelete().Where(kccourse.PushStatus(1)).Select("id").WithCourseTeachers(func(query *ent.KcCourseTeacherQuery) {
				query.Select("id", "teacher_id", "course_id").Where(kccourseteacher.ShowStatus(1))
			})
		})
	}).AllX(ctx)

	courseList := s.KcUserCourse.Query().SoftDelete().Where(kcusercourse.UserID(uid.(int))).WithCourse(func(query *ent.KcCourseQuery) {
		query.SoftDelete().Where(kccourse.PushStatus(1)).WithCourseTeachers(func(query *ent.KcCourseTeacherQuery) {
			query.Select("id", "teacher_id", "course_id").Where(kccourseteacher.ShowStatus(1))
		})
	}).AllX(ctx)

	if len(courseList) > 0 {
		teacherIds := []int{}
		com := app.Common{}

		res := response.TeacherAnswerList{
			ClassTeacher: []response.TeacherDetail{},
			TeacherList:  []response.TeacherDetail{},
		}

		var classTeacherId int

		for _, v := range classList {
			if v.Edges.Class != nil {
				if v.Edges.Class.Edges.MasterTeachers != nil {
					classTeacherId = v.Edges.Class.Edges.MasterTeachers.ID

					if classTeacherId > 0 {
						teacherIds = append(teacherIds, classTeacherId)
					}
				}

				if v.Edges.Class.Edges.KcClassCourses != nil {
					for _, v1 := range v.Edges.Class.Edges.KcClassCourses {
						for _, v2 := range v1.Edges.CourseTeachers {
							teacherIds = append(teacherIds, v2.TeacherID)
						}
					}
				}
			}
		}

		for _, v := range courseList {
			if v.Edges.Course != nil {
				if v.Edges.Course.Edges.CourseTeachers != nil {
					for _, v1 := range v.Edges.Course.Edges.CourseTeachers {
						teacherIds = append(teacherIds, v1.TeacherID)
					}
				}
			}
		}

		Ids := com.RmDuplicate(teacherIds)

		teacherList := s.Teacher.Query().Where(teacher.IDIn(Ids...), teacher.Status(1)).WithTeacherCourses(func(query *ent.KcCourseTeacherQuery) {
			query.Where(kccourseteacher.ShowStatus(1)).WithCourse(func(query *ent.KcCourseQuery) {
				query.SoftDelete().Select("id", "course_name")
			})
		}).WithMajors(func(query *ent.MajorQuery) {
			query.SoftDelete().Select("id", "name")
		}).WithTeacherTags(func(query *ent.TeacherTagQuery) {
			query.SoftDelete().Select("id", "name", "teacher_id")
		}).WithAttachment(func(query *ent.AttachmentQuery) {
			query.SoftDelete().Select("id", "filename")
		}).AllX(ctx)

		for _, v := range teacherList {
			re := response.TeacherDetail{
				Majors:    []string{},
				Tags:      []string{},
				CourseArr: []string{},
			}
			re.Name = v.Name
			re.Nickname = v.Nickname
			re.TeacherId = v.ID
			if v.Edges.Attachment != nil {
				re.Avatar = app2.GetOssHost() + v.Edges.Attachment.Filename
			}
			re.TeachingAge = int(v.TeachingAge)
			if len(v.Edges.Majors) > 0 {
				for _, v1 := range v.Edges.Majors {
					re.Majors = append(re.Majors, v1.Name)
				}
			}

			if len(v.Edges.TeacherTags) > 0 {
				for _, v2 := range v.Edges.TeacherTags {
					re.Tags = append(re.Tags, v2.Name)
				}
			}

			if len(v.Edges.TeacherTags) > 0 {
				for _, v1 := range v.Edges.TeacherCourses {
					if v1.Edges.Course != nil {
						re.CourseArr = append(re.CourseArr, v1.Edges.Course.CourseName)
					}
				}
			}

			//如果是班主任
			if v.ID == classTeacherId {
				res.ClassTeacher = append(res.ClassTeacher, re)
			} else {
				res.TeacherList = append(res.TeacherList, re)
			}
		}
		ca.Set(Key, res)
		return res, nil
	}
	ca.Set(Key, nil)
	return nil, nil
}

//学圆提问
func UserAskAnswer(ctx *gin.Context) (interface{}, error) {
	var req request.UserTeacherAsk
	err := ctx.Bind(&req)
	if err != nil {
		return nil, errorno.NewParamErr(err)
	}

	imgIds := ctx.PostForm("img_ids") //兼容ios

	if imgIds != "" {
		err = json.Unmarshal([]byte(imgIds), &req.ImgIds)
	}

	if len(req.ImgIds)==0 && req.AskDesc == ""{
		return nil,errorno.NewParamErr(err)
	}
	uid, _ := ctx.Get("uid")
	err = store.WithTx(ctx, func(ctx context.Context) error {
		s := store.WithContext(ctx)

		info := s.UserAskAnswer.Create().SetUserID(uid.(int)).SetAskDesc(req.AskDesc).SetTeacherID(req.TeacherId).SaveX(ctx)

		if req.ImgIds != nil && len(req.ImgIds) > 0 { //插入图片id
			numb := make([]*ent.UserAskAnswerAttachmentCreate, len(req.ImgIds))
			for i, v := range req.ImgIds {
				query := s.UserAskAnswerAttachment.
					Create().
					SetAttachmentID(v).
					SetAskID(info.ID)

				numb[i] = query
			}
			s.UserAskAnswerAttachment.CreateBulk(numb...).SaveX(ctx)
		}
		return nil
	})

	return nil, nil
}

//提交评价
func AddUserCourseAppraise(ctx *gin.Context) (interface{}, error) {
	var req request.AddUserCourseAppraise
	err := ctx.Bind(&req)
	if err != nil {
		return nil, errorno.NewParamErr(err)
	}
	uid, _ := ctx.Get("uid")
	s := store.WithContext(ctx)
	info := s.UserCourseAppraise.Query().SoftDelete().Where(usercourseappraise.CourseIDIn(req.CourseId), usercourseappraise.SmallCateID(req.SmallId), usercourseappraise.UserID(uid.(int))).ExistX(ctx)
	if info {
		return nil, errorno.NewErr(errorno.Errno{
			Code:    20002,
			Message: "您的评价已收到",
		})
	}

	req.CompositeScore = float64(math.Floor((req.TeachAtmosphereScore+req.TeachContentScore+req.TeachAttitudeScore)/3 + 0.5))
		/*strconv.ParseFloat(fmt.Sprintf("%.1f",(req.TeachAtmosphereScore+req.TeachContentScore+req.TeachAttitudeScore)/3),64)*/

	err = store.WithTx(ctx, func(ctx context.Context) error {
		s.UserCourseAppraise.Create().
			SetUserID(uid.(int)).
			SetType(uint8(req.Type)).
			SetDesc(req.Desc).
			SetShowStatus(1).
			SetSmallCateID(req.SmallId).
			SetCourseID(req.CourseId).
			SetTeachAtmosphereScore(req.TeachAtmosphereScore).
			SetTeacherImpression(req.TeacherImpression).
			SetCompositeScore(req.CompositeScore).
			SetTeachAttitudeScore(req.TeachAttitudeScore).
			SetTeachContentScore(req.TeachContentScore).
			SaveX(ctx)
		return nil
	})

	if err != nil {
		return nil, errorno.NewInternalErr(err)
	}

	return nil, nil
}

//获取课程评价
func GetUserCourseAppraiseList(ctx *gin.Context) (interface{}, error) {
	var req request.GetUserCourseAppraise
	var res []response.UserCourseAppraiseList
	err := ctx.Bind(&req)
	if err != nil {
		return nil, errorno.NewParamErr(err)
	}
	s := store.WithContext(ctx)
	uid, _ := ctx.Get("uid")

	info := s.UserCourseAppraise.Query().SoftDelete().Where(usercourseappraise.CourseID(req.CourseId), usercourseappraise.SmallCateID(req.SmallId), usercourseappraise.UserID(uid.(int))).ExistX(ctx)

	query := s.UserCourseAppraise.Query().SoftDelete().Where(usercourseappraise.CourseID(req.CourseId), usercourseappraise.ShowStatus(1)).WithUser(func(query *ent.UserQuery) {
		query.SoftDelete().Select("id", "username", "nickname", "avatar")
	}).WithSmallCate(func(query *ent.KcCourseSmallCategoryQuery) {
		query.SoftDelete().WithCourse(func(query *ent.KcCourseQuery) {
			query.SoftDelete().Select("id", "course_name", "people_num", "course_type")
		})
	})

	count, err := query.Count(ctx)
	var page response.PageResponse
	page.Total = count
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, errorno.NewInternalErr(err)
	}
	PageSize := 20
	list := query.ForPage(&req.Page, &PageSize).Order(ent.Desc("id")).AllX(ctx)

	if len(list) > 0 {
		for _, v := range list {
			re := response.UserCourseAppraiseList{}
			re.Id = v.ID
			re.CompositeScore = v.CompositeScore
			re.TeachAttitudeScore = v.TeachAttitudeScore
			re.TeachContentScore = v.TeachContentScore
			re.TeachAtmosphereScore = v.TeachAtmosphereScore
			re.Desc = v.Desc
			re.UserId = v.UserID
			re.SmallId = v.SmallCateID
			if !app2.IsNil(v.TeacherImpression) {
				split := strings.Split(v.TeacherImpression, ",")
				for _, v := range split {
					re.TeacherImpression = append(re.TeacherImpression, v)
				}
			}

			if v.Edges.SmallCate != nil {
				if v.Edges.SmallCate.Edges.Course != nil {
					re.CourseName = v.Edges.SmallCate.Edges.Course.CourseName
					re.CourseId = v.Edges.SmallCate.Edges.Course.ID
				}
			}

			if v.Edges.User != nil {
				re.UserName = v.Edges.User.Username
				re.Avatar = v.Edges.User.Avatar
			}
			re.Time = v.CreatedAt.Format("2006-01-02 15:04:05")


			res = append(res, re)

		}

		var data = map[string]interface{}{
			"is_report":info,//是否评价过 1 是 0否
			"list": res,
			"total": page.Total,
		}

		return data, nil
	}

	return nil, nil
}

//公开课
func GetLiveOpenCourses(ctx *gin.Context) (interface{}, error) {
	ca := cache.CommCache
	Key := cache.OpenCourseKey
	if data, ok := ca.Get(Key); ok {
		return data, nil
	}

	s := store.WithContext(ctx)
	loc, _ := time.LoadLocation("Local")
	oldTime := time.Now().AddDate(0, 0, -30) //过去30天
	newTime := time.Now().AddDate(0, 0, 7)   //未来7天
	today := time.Now().Format("2006-01-02")
	fuday7 := newTime.Format("2006-01-02")
	old30 := oldTime.Format("2006-01-02")

	//4.拼接成当天23点时间字符串
	endDate := today + " 23:59:59"
	oldDate := old30 + " 23:59:59"
	newDate := fuday7 + " 23:59:59"
	//得到23点日期 2021-04-24 23:59:59 +0800 CST
	endTime, _ := time.ParseInLocation("2006-01-02 15:04:05", endDate, loc)

	oldTime, _ = time.ParseInLocation("2006-01-02 15:04:05", oldDate, loc) //过去30天 结束时间
	newTime, _ = time.ParseInLocation("2006-01-02 15:04:05", newDate, loc) //未来7天结束时间

	com := app.Common{}
	todayStart := com.GetZeroTime(time.Now())

	openLiveList := s.KcCourse.Query().SoftDelete().Where(kccourse.PushStatus(1), kccourse.CourseType(LiveOpenCourses)).WithCourseSmallCategorys(func(query *ent.KcCourseSmallCategoryQuery) {
		query.SoftDelete().Where(kccoursesmallcategory.LiveSmallStartGTE(oldTime), kccoursesmallcategory.LiveSmallStartLTE(newTime))
	}).WithAttachment(func(query *ent.AttachmentQuery) {
		query.SoftDelete().Select("id", "filename")
	}).WithCourseTeachers(func(query *ent.KcCourseTeacherQuery) {
		query.Select("id", "teacher_id", "course_id").WithTeacher(func(query *ent.TeacherQuery) {
			query.WithAttachment(func(query *ent.AttachmentQuery) {
				query.SoftDelete().Select("id", "filename")
			})
		})
	}).Order(func(s *sql2.Selector) {
		t := sql2.Table(kccoursesmallcategory.Table)
		s.Join(t).On(s.C(kccourse.FieldID), t.C(kccoursesmallcategory.FieldCourseID))
		s.OrderBy(t.C(sql2.Desc(kccoursesmallcategory.FieldLiveSmallStart)))
	}).AllX(ctx)

	resToday := []response.LiveDetail{} //今日
	resJinqi := []response.LiveDetail{} //近7天
	resLast := []response.LiveDetail{}  //往期最近一个月
	c := app.Course{}

	if len(openLiveList) > 0 {
		for _, v := range openLiveList {
			if v.Edges.CourseSmallCategorys != nil && len(v.Edges.CourseSmallCategorys) > 0 {
				re := response.LiveDetail{}
				date := v.Edges.CourseSmallCategorys[0].LiveSmallStart.Format("2006-01-02")
				timeAt := v.Edges.CourseSmallCategorys[0].LiveSmallStart
				re.CourseName = v.CourseName
				re.LiveName = v.Edges.CourseSmallCategorys[0].SmallName
				status := c.CheckLiveType(int(v.CourseType), int(v.Edges.CourseSmallCategorys[0].LiveSmallStatus), v.Edges.CourseSmallCategorys[0].LiveRoomID)
				re.LiveStatus = status.LiveStatus
				re.LiveDesc = v.CourseDesc
				re.CourseId = v.ID
				re.IsEnroll = 0
				re.LiveSmallId = v.Edges.CourseSmallCategorys[0].ID
				if v.Edges.CourseSmallCategorys[0].LiveSmallTime > 0 {
					re.LiveTime = v.Edges.CourseSmallCategorys[0].LiveSmallStart
					end := v.Edges.CourseSmallCategorys[0].LiveSmallStart.Unix() + int64(v.Edges.CourseSmallCategorys[0].LiveSmallTime)
					ends := time.Unix(end, 0)
					re.LiveAt = v.Edges.CourseSmallCategorys[0].LiveSmallStart.Format("15:04") + "-" + ends.Format("15:04")
					re.LiveDate = v.Edges.CourseSmallCategorys[0].LiveSmallStart.Format("2006-01-02 ") + v.Edges.CourseSmallCategorys[0].LiveSmallStart.Format("15:04") + "-" + ends.Format("15:04")
					re.LiveDateAt = v.Edges.CourseSmallCategorys[0].LiveSmallStart.Format("01.02")
				}

				if v.Edges.Attachment != nil {
					re.ImgCover = app2.GetOssHost() + v.Edges.Attachment.Filename
				}
				if v.Edges.CourseTeachers != nil && len(v.Edges.CourseTeachers) > 0 {
					if v.Edges.CourseTeachers[0].Edges.Teacher != nil {
						re.TeacherName = v.Edges.CourseTeachers[0].Edges.Teacher.Name
						if v.Edges.CourseTeachers[0].Edges.Teacher.Edges.Attachment != nil {
							re.TeacherAvatar = app2.GetOssHost() + v.Edges.CourseTeachers[0].Edges.Teacher.Edges.Attachment.Filename
						}
					}
				}

				if date == today { //是否为今天
					resToday = append(resToday, re)
				} else if timeAt.After(endTime) && timeAt.Before(newTime) { //未来七天
					resJinqi = append(resJinqi, re)
				} else if timeAt.Before(todayStart) && timeAt.After(oldTime) { //过去30天
					resLast = append(resLast, re)
				}

			}
		}

		if len(resToday) > 0 || len(resJinqi) > 0 || len(resLast) > 0 {

			sort.Slice(resJinqi, func(i, j int) bool {
				return resJinqi[i].LiveTime.Before(resJinqi[j].LiveTime) // 升序
			})

			var data = map[string]interface{}{
				"today_list":  resToday,
				"lately_list": resJinqi,
				"last_list":   resLast,
			}

			ca.Set(Key, data)

			return data, nil
		}
	}
	//没有数据展示缓存nil 10秒过期
	ca.Set(Key, nil)
	return nil, nil

}

//精品课程
func GetRecordCourses(ctx *gin.Context) (interface{}, error) {
	var req request.PageInfo
	err := ctx.Bind(&req)
	if err != nil {
		return nil, errorno.NewParamErr(err)
	}
	ca := cache.CommCache
	Key := cache.RecommendCourse + strconv.Itoa(req.Page)
	if data, ok := ca.Get(Key); ok {
		return data, nil
	}

	s := store.WithContext(ctx)
	query := s.KcCourse.Query().SoftDelete().Where(kccourse.PushStatus(1), kccourse.CourseType(LiveRecordCourses)).WithCourseSmallCategorys(func(query *ent.KcCourseSmallCategoryQuery) {
		query.SoftDelete().Select("id", "course_id", "live_small_status", "live_room_id")
	}).WithCourseTeachers(func(query *ent.KcCourseTeacherQuery) {
		query.Select("id", "teacher_id", "course_id").Where(kccourseteacher.ShowStatus(1)).WithTeacher(func(query *ent.TeacherQuery) {
			query.SoftDelete().Select("id", "avatar_id", "name").WithAttachment(func(query *ent.AttachmentQuery) {
				query.SoftDelete().Select("filename", "id")
			})
		})
	}).WithAttachment(func(query *ent.AttachmentQuery) {
		query.SoftDelete().Select("id", "filename")
	})
	count, err := query.Count(ctx)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, errorno.NewInternalErr(err)
	}

	res := []response.RecommendCourse{}

	PageSize := 20
	list := query.ForPage(&req.Page, &PageSize).Order(ent.Desc("id")).AllX(ctx)
	c := app.Course{}

	for _, v := range list {
		re := response.RecommendCourse{}
		re.CourseId = v.ID
		re.CourseDesc = v.CourseDesc
		re.CourseName = v.CourseName
		re.SmallCount = len(v.Edges.CourseSmallCategorys)
		if v.Edges.Attachment != nil && v.Edges.Attachment.Filename != "" {
			re.ImgCover = app2.GetOssHost() + v.Edges.Attachment.Filename
		}

		if len(v.Edges.CourseSmallCategorys) > 0 {
			status := c.CheckLiveType(int(v.CourseType), int(v.Edges.CourseSmallCategorys[0].LiveSmallStatus), v.Edges.CourseSmallCategorys[0].LiveRoomID)
			re.LiveStatus = status.LiveStatus
		}

		if v.Edges.CourseTeachers != nil {
			if len(v.Edges.CourseTeachers) > 0 {
				if v.Edges.CourseTeachers != nil && len(v.Edges.CourseTeachers) > 0 {
					if v.Edges.CourseTeachers[0].Edges.Teacher != nil {
						re.TeacherName = v.Edges.CourseTeachers[0].Edges.Teacher.Name
						if v.Edges.CourseTeachers[0].Edges.Teacher.Edges.Attachment != nil {
							re.TeacherAvatar = app2.GetOssHost() + v.Edges.CourseTeachers[0].Edges.Teacher.Edges.Attachment.Filename
						}
					}
				}
			}
		}
		res = append(res, re)
	}

	if len(res) > 0 {
		var data = map[string]interface{}{
			"list":  res,
			"total": count,
		}
		ca.Set(Key, data)
		return data, nil
	}
	ca.Set(Key, nil)
	return nil, nil

}
//pc端首页
func GetPcIndex(ctx *gin.Context)(interface{},error){
	s := store.WithContext(ctx)
	bannerList := s.Advertise.Query().SoftDelete().Where(advertise.Status(1), advertise.Position(4)).WithAttachment(func(query *ent.AttachmentQuery) {
		query.SoftDelete().Select("id", "filename")
	}).AllX(ctx)

	var resBanner []*response.AppIndexBanner

	for _, v := range bannerList {
		rb := response.AppIndexBanner{}
		if v.Edges.Attachment != nil {
			if v.Edges.Attachment.Filename != "" {
				rb.Img = app2.GetOssHost() + v.Edges.Attachment.Filename
			}
		}
		rb.Url = v.AdURL
		rb.Name = v.Name

		resBanner = append(resBanner, &rb)
	}

	return resBanner,nil
}

//首页
func GetIndexInfo(ctx *gin.Context) (interface{}, error) {
	ca := cache.CommCache
	Key := cache.AppIndexKey
	if data, ok := ca.Get(Key); ok {
		return data, nil
	}
	s := store.WithContext(ctx)
	now := time.Now()
	com := app.Common{}
	start, _ := com.GetDateStartEnd(now)

	bannerList := s.Advertise.Query().SoftDelete().Where(advertise.Status(1), advertise.Position(1)).WithAttachment(func(query *ent.AttachmentQuery) {
		query.SoftDelete().Select("id", "filename")
	}).AllX(ctx)
	//直播公开课程
	courseList := s.KcCourse.Query().SoftDelete().Where(kccourse.PushStatus(1), kccourse.CourseType(LiveOpenCourses)).WithCourseSmallCategorys(func(query *ent.KcCourseSmallCategoryQuery) {
		query.SoftDelete().Where(kccoursesmallcategory.LiveSmallStartGTE(start)).Order(ent.Asc("live_small_start")).Limit(6)
	}).WithAttachment(func(query *ent.AttachmentQuery) {
		query.SoftDelete().Select("id", "filename")
	}).WithCourseTeachers(func(query *ent.KcCourseTeacherQuery) {
		query.Select("id", "teacher_id", "course_id").WithTeacher(func(query *ent.TeacherQuery) {
			query.WithAttachment(func(query *ent.AttachmentQuery) {
				query.SoftDelete().Select("id", "filename")
			})
		})
	}).Order(func(s *sql2.Selector) {
		t := sql2.Table(kccoursesmallcategory.Table)
		s.Join(t).On(s.C(kccourse.FieldID), t.C(kccoursesmallcategory.FieldCourseID))
		s.OrderBy(t.C(sql2.Asc(kccoursesmallcategory.FieldLiveSmallStart)))
	}).AllX(ctx)

	//录播公开课
	recodeCourseList := s.KcCourse.Query().SoftDelete().Where(kccourse.PushStatus(1), kccourse.CourseType(LiveRecordCourses)).WithCourseSmallCategorys(func(query *ent.KcCourseSmallCategoryQuery) {
		query.SoftDelete().Select("id", "course_id")
	}).WithAttachment(func(query *ent.AttachmentQuery) {
		query.SoftDelete().Select("id", "filename")
	}).WithCourseTeachers(func(query *ent.KcCourseTeacherQuery) {
		query.Select("id", "teacher_id", "course_id").WithTeacher(func(query *ent.TeacherQuery) {
			query.WithAttachment(func(query *ent.AttachmentQuery) {
				query.SoftDelete().Select("id", "filename")
			})
		})
	}).Order(ent.Desc("created_at")).Limit(3).AllX(ctx)

	//activity
	var resBanner = []*response.AppIndexBanner{}
	var resAppIndexCourse = []*response.AppIndexCourse{}
	var resAppIndexrecodeCourseList = []*response.AppIndexRecodeCourse{}

	if len(bannerList) > 0 || len(courseList) > 0 || len(recodeCourseList) > 0 {
		c := app.Course{}

		for _, v := range bannerList {
			rb := response.AppIndexBanner{}
			if v.Edges.Attachment != nil {
				if v.Edges.Attachment.Filename != "" {
					rb.Img = app2.GetOssHost() + v.Edges.Attachment.Filename
				}
			}
			rb.Url = v.AdURL
			rb.Name = v.Name

			resBanner = append(resBanner, &rb)
		}

		//公开课
		for _, v := range courseList {
			if len(v.Edges.CourseSmallCategorys) > 0 {
				rc := response.AppIndexCourse{}
				rc.CourseName = v.CourseName
				rc.LiveSmallId = v.Edges.CourseSmallCategorys[0].ID
				rc.StudyCount = v.PeopleNum
				if v.Edges.Attachment != nil {
					rc.ImgCover = app2.GetOssHost() + v.Edges.Attachment.Filename
				}
				if v.Edges.CourseTeachers != nil {
					if len(v.Edges.CourseTeachers) > 0 {
						if v.Edges.CourseTeachers != nil && len(v.Edges.CourseTeachers) > 0 {
							if v.Edges.CourseTeachers[0].Edges.Teacher != nil {
								rc.TeacherName = v.Edges.CourseTeachers[0].Edges.Teacher.Name

								if v.Edges.CourseTeachers[0].Edges.Teacher.Edges.Attachment != nil {
									rc.TeacherAvatar = app2.GetOssHost() + v.Edges.CourseTeachers[0].Edges.Teacher.Edges.Attachment.Filename
								}
							}
						}
					}
				}

				if v.Edges.CourseSmallCategorys[0].LiveSmallTime > 0 {
					end := v.Edges.CourseSmallCategorys[0].LiveSmallStart.Unix() + int64(v.Edges.CourseSmallCategorys[0].LiveSmallTime)
					ends := time.Unix(end, 0)
					rc.LiveAt = v.Edges.CourseSmallCategorys[0].LiveSmallStart.Format("2006-01-02 ") + v.Edges.CourseSmallCategorys[0].LiveSmallStart.Format("15:04") + "-" + ends.Format("15:04")
				}

				if len(v.Edges.CourseSmallCategorys) > 0 {
					status := c.CheckLiveType(int(v.CourseType), int(v.Edges.CourseSmallCategorys[0].LiveSmallStatus), v.Edges.CourseSmallCategorys[0].LiveRoomID)
					rc.LiveStatus = status.LiveStatus
				}

				resAppIndexCourse = append(resAppIndexCourse, &rc)
			}
		}

		for _, v := range recodeCourseList {
			re := response.AppIndexRecodeCourse{}
			re.CourseId = v.ID
			re.CourseName = v.CourseName

			if v.Edges.Attachment != nil && v.Edges.Attachment.Filename != "" {
				re.ImgCover = app2.GetOssHost() + v.Edges.Attachment.Filename
			}

			if v.Edges.CourseTeachers != nil {
				if len(v.Edges.CourseTeachers) > 0 {
					if v.Edges.CourseTeachers != nil && len(v.Edges.CourseTeachers) > 0 {
						if v.Edges.CourseTeachers[0].Edges.Teacher != nil {
							re.TeacherName = v.Edges.CourseTeachers[0].Edges.Teacher.Name
							if v.Edges.CourseTeachers[0].Edges.Teacher.Edges.Attachment != nil {
								re.TeacherAvatar = app2.GetOssHost() + v.Edges.CourseTeachers[0].Edges.Teacher.Edges.Attachment.Filename
							}
						}
					}
				}
			}

			if len(v.Edges.CourseSmallCategorys) > 0 {
				re.SmallCount = len(v.Edges.CourseSmallCategorys)

			}

			resAppIndexrecodeCourseList = append(resAppIndexrecodeCourseList, &re)

		}

		var data = map[string]interface{}{
			"banner_list":        resBanner,
			"open_course_list":   resAppIndexCourse,
			"recode_course_list": resAppIndexrecodeCourseList,
		}
		ca.Set(Key, data)

		return data, nil
	}

	ca.Set(Key, nil)

	return nil, nil
}

func GetMessInfo(ctx *gin.Context) (interface{}, error) {
	var req request.MessageAskType
	err := ctx.Bind(&req)
	if err != nil {
		return nil, errorno.NewParamErr(err)
	}
	uid, _ := ctx.Get("uid")
	s := store.WithContext(ctx)
	now := time.Now()
	com := app.Common{}
	PageSize := 20
	switch req.Type {
	case 1:
		res := []response.StyMessage{}
		//系统消息
		query := s.Message.Query().SoftDelete().Where(message.ClassIDIsNil(), message.CourseIDIsNil())
		count := query.CountX(ctx)
		list := query.ForPage(&req.Page, &PageSize).Order(ent.Desc("id")).AllX(ctx)
		for _, v := range list {
			re := response.StyMessage{}
			re.Desc = v.Detail
			re.Name = v.Name
			re.Date = v.CreatedAt.Format("2006-01-02 15:04:05")
			res = append(res, re)
		}

		var data = map[string]interface{}{
			"list":  res,
			"total": count,
		}

		return data, nil
	case 2:
		res := []response.AskMessage{}
		//问答消息
		query := s.UserAskAnswer.Query().SoftDelete().Where(useraskanswer.UserID(uid.(int)), useraskanswer.ShowStatus(1), useraskanswer.AnswerStatus(1)).
			WithAskAnswersAttachments(func(query *ent.UserAskAnswerAttachmentQuery) {
				query.WithAttachment(func(query *ent.AttachmentQuery) {
					query.SoftDelete().Select("id", "filename")
				})
			}).WithTeacher(func(query *ent.TeacherQuery) {
			query.SoftDelete().WithAttachment(func(query *ent.AttachmentQuery) {
				query.SoftDelete().Select("id", "filename")
			})
		})
		count, _ := query.Count(ctx)
		list := query.ForPage(&req.Page, &PageSize).Order(ent.Desc("id")).AllX(ctx)

		for _, v := range list {
			r := response.AskMessage{
				ReplyImg: []string{},
			}
			r.AnswerDesc = v.AskDesc
			r.ReplyDesc = v.AnswerDesc
			r.TimeAt = com.GetDiffDescDate(now, v.AnswerAt)
			if v.Edges.Teacher != nil {
				r.TeacherName = v.Edges.Teacher.Name
				if v.Edges.Teacher.Edges.Attachment != nil {
					r.TeacherAvatar = app2.GetOssHost() + v.Edges.Teacher.Edges.Attachment.Filename
				}
			}
			if len(v.Edges.AskAnswersAttachments) > 0 {
				for _, v1 := range v.Edges.AskAnswersAttachments {
					// 图片类型 1：学生提问，2：老师回复
					if v1.Edges.Attachment != nil {
						if int(v1.Type) == 1 {
							r.AnswerImg = app2.GetOssHost() + v1.Edges.Attachment.Filename
						} else if int(v1.Type) == 2 {
							r.ReplyImg = append(r.ReplyImg, app2.GetOssHost()+v1.Edges.Attachment.Filename)
						}
					}
				}
			}
			res = append(res, r)
		}

		var data = map[string]interface{}{
			"list":  res,
			"total": count,
		}

		return data, nil
	}

	return nil, nil
}

//小节课程观看记录
func AddVideoRecord(ctx *gin.Context) (interface{}, error) {
	var req request.AddVideoRecord
	err := ctx.Bind(&req)
	if err != nil {
		return nil, errorno.NewParamErr(err)
	}
	uid, _ := ctx.Get("uid")
	s := store.WithContext(ctx)
	query := s.VideoRecord.Query().Where(videorecord.UserID(uid.(int)), videorecord.SmallID(req.SmallId))
	data := query.FirstX(ctx)
	if data != nil {
		s.VideoRecord.UpdateOneID(data.ID).SetVideoTime(req.VideoTime).SetViewAt(time.Now()).SetVideoName(req.Name).SaveX(ctx)
	} else {
		s.VideoRecord.Create().SetUserID(uid.(int)).SetViewAt(time.Now()).SetSmallID(req.SmallId).SetVideoName(req.Name).SaveX(ctx)
	}

	return nil, nil
}

//观看记录
func GetVideoRecord(ctx *gin.Context) (interface{}, error) {
	var req request.PageInfo
	err := ctx.Bind(&req)
	if err != nil {
		return nil, errorno.NewParamErr(err)
	}
	uid, _ := ctx.Get("uid")
	s := store.WithContext(ctx)
	oldTime := time.Now().AddDate(0, 0, -30) //过去30天

	var res = []response.VideoRecord{}
	query := s.VideoRecord.Query().SoftDelete().Where(videorecord.UserID(uid.(int)), videorecord.ViewAtGTE(oldTime)).WithSmallCourse(func(query *ent.KcCourseSmallCategoryQuery) {
		query.SoftDelete().Select("id", "course_id", "small_name", "live_small_status", "live_room_id").WithCourse(func(query *ent.KcCourseQuery) {
			query.SoftDelete().Select("id", "course_cover_img_id", "course_type").WithAttachment(func(query *ent.AttachmentQuery) {
				query.SoftDelete().Select("id", "filename")
			}).WithCourseTeachers(func(query *ent.KcCourseTeacherQuery) {
				query.WithTeacher(func(query *ent.TeacherQuery) {
					query.SoftDelete().WithAttachment(func(query *ent.AttachmentQuery) {
						query.SoftDelete().Select("id", "filename")
					})
				})
			})
		})
	})
	count, _ := query.Count(ctx)
	PageSize := 20
	list := query.ForPage(&req.Page, &PageSize).Order(ent.Desc("view_at")).AllX(ctx)
	c := app.Course{}

	for _, v := range list {
		re := response.VideoRecord{}
		re.SmallId = v.SmallID
		re.ViewAt = v.ViewAt.Format("2006-01-02 15:04")
		if v.Edges.SmallCourse != nil {
			re.SmallName = v.Edges.SmallCourse.SmallName

			if v.Edges.SmallCourse.Edges.Course != nil {
				re.CourseId = v.Edges.SmallCourse.Edges.Course.ID

				status := c.CheckLiveType(int(v.Edges.SmallCourse.Edges.Course.CourseType), int(v.Edges.SmallCourse.LiveSmallStatus), v.Edges.SmallCourse.LiveRoomID)
				re.SmallType = status.SmallType
				re.LiveStatus = status.LiveStatus
				if v.Edges.SmallCourse.Edges.Course.Edges.Attachment != nil {
					if v.Edges.SmallCourse.Edges.Course.Edges.Attachment.Filename != "" {
						re.ImgCover = app2.GetOssHost() + v.Edges.SmallCourse.Edges.Course.Edges.Attachment.Filename
					}
				}

				if len(v.Edges.SmallCourse.Edges.Course.Edges.CourseTeachers) > 0 {
					if v.Edges.SmallCourse.Edges.Course.Edges.CourseTeachers[0].Edges.Teacher != nil {
						re.TeacherName = v.Edges.SmallCourse.Edges.Course.Edges.CourseTeachers[0].Edges.Teacher.Name
						if v.Edges.SmallCourse.Edges.Course.Edges.CourseTeachers[0].Edges.Teacher.Edges.Attachment != nil {
							if v.Edges.SmallCourse.Edges.Course.Edges.CourseTeachers[0].Edges.Teacher.Edges.Attachment.Filename != "" {
								re.TeacherAvatar = app2.GetOssHost() + v.Edges.SmallCourse.Edges.Course.Edges.CourseTeachers[0].Edges.Teacher.Edges.Attachment.Filename
							}
						}
					}
				}
			}
		}
		res = append(res, re)
	}

	if len(res) > 0 {
		var data = map[string]interface{}{
			"list":  res,
			"total": count,
		}
		return data, nil
	}

	return nil, nil
}

//清空消息
/*func DelMess(ctx *gin.Context) (interface{}, error) {
	var req request.MessageAskType
	err := ctx.Bind(&req)

	if err != nil {
		return nil, errorno.NewParamErr(err)
	}
	uid, _ := ctx.Get("uid")
	s := store.WithContext(ctx)

	if req.Type == 1{
		s.Message.Update().Where(Message.User) SoftDelete().
	}

}*/
