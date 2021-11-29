package appapi

import (
	"context"
	"database/sql"
	"encoding/json"
	sql2 "entgo.io/ent/dialect/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"math"
	"strconv"
	"time"
	"tkserver/httpapi/appapi/request"
	"tkserver/httpapi/appapi/response"
	"tkserver/internal/app"
	"tkserver/internal/cache"
	"tkserver/internal/errorno"
	"tkserver/internal/store"
	"tkserver/internal/store/ent"
	"tkserver/internal/store/ent/collection"
	"tkserver/internal/store/ent/groupcard"
	"tkserver/internal/store/ent/itemcategory"
	"tkserver/internal/store/ent/kcuserclass"
	"tkserver/internal/store/ent/kcusercourse"
	"tkserver/internal/store/ent/level"
	"tkserver/internal/store/ent/major"
	"tkserver/internal/store/ent/makeuserquestionrecord"
	"tkserver/internal/store/ent/tkchapter"
	"tkserver/internal/store/ent/tkexampaper"
	"tkserver/internal/store/ent/tkexampartitionquestionlink"
	"tkserver/internal/store/ent/tkquestion"
	"tkserver/internal/store/ent/tkquestionbank"
	"tkserver/internal/store/ent/tkquestionbankcity"
	"tkserver/internal/store/ent/tkquestionbankmajor"
	"tkserver/internal/store/ent/tkquestionsection"
	"tkserver/internal/store/ent/tkuserexamscorerecord"
	"tkserver/internal/store/ent/tkuserquestionbankrecord"
	"tkserver/internal/store/ent/tkuserquestionrecord"
	"tkserver/internal/store/ent/tkuserrandomexamrecode"
	"tkserver/internal/store/ent/tkuserwrongquestionrecode"
	"tkserver/pkg/asynctask"
	"tkserver/pkg/htmltojson"
	"tkserver/pkg/log"
	app2 "tkserver/pkg/requestparam"
)

const (
	SuiJiJuan  = 2 //随机卷
	GuDingJuan = 1 //固定卷
)

const (
	ExamStart      = 1 //考试开始
	ExamEnd        = 2 //考试结束
	ExamViewResult = 3 //查看成绩
	ExamNot        = 4 //考试未开始
)

//题库首页
func GetQuestionIndex(ctx *gin.Context) (interface{}, error) {
	var req request.QuestionIndex
	err := ctx.Bind(&req)
	if err != nil {
		return nil, errorno.NewParamErr(err)
	}

	s := store.WithContext(ctx)
	uid, _ := ctx.Get("uid")
	//班级课程
	infoList := s.KcUserClass.Query().Where(kcuserclass.UserID(uid.(int))).
		WithClass(func(query *ent.KcClassQuery) {
			query.WithKcClassCourses(func(query2 *ent.KcCourseQuery) {
				query2.SoftDelete().WithQuestionBank()
			})
		}).AllX(ctx)

	couseList := s.KcUserCourse.Query().SoftDelete().Where(kcusercourse.UserID(uid.(int))).WithCourse(func(query *ent.KcCourseQuery) {
		query.SoftDelete().WithQuestionBank()
	}).AllX(ctx)

	var list = []*response.QuestionClass{}
	var otherList = []*response.ClassCourse{}
	var quesIds []int

	for _, v := range infoList {
		i := &response.QuestionClass{
			ClassCourse: []response.ClassCourse{},
		}
		i.Name = v.Edges.Class.ClassTitle
		i.ClassId = v.Edges.Class.ID
		i.CourseCount = len(v.Edges.Class.Edges.KcClassCourses)
		for _, c := range v.Edges.Class.Edges.KcClassCourses {
			//课程题库
			o := response.ClassCourse{}
			o.CourseID = c.ID
			if c.Edges.QuestionBank != nil {
				o.QuestionNum = o.QuestionNum + c.Edges.QuestionBank.QuestionCount
				//题库名称
				o.QuestionBankName = c.Edges.QuestionBank.Name

				o.QuestionBankId = c.Edges.QuestionBank.ID
				if !app2.IsContainInt(quesIds, c.Edges.QuestionBank.ID) {
					quesIds = append(quesIds, c.Edges.QuestionBank.ID)
					i.ClassCourse = append(i.ClassCourse, o)

				}
			}
		}
		list = append(list, i)
	}

	for _, v := range couseList {
		if v.Edges.Course != nil {
			//课程题库
			if v.Edges.Course.Edges.QuestionBank != nil {
				o := response.ClassCourse{}
				o.CourseID = v.Edges.Course.ID
				o.QuestionNum = o.QuestionNum + v.Edges.Course.Edges.QuestionBank.QuestionCount
				o.QuestionBankName = v.Edges.Course.Edges.QuestionBank.Name

				o.QuestionBankId = v.Edges.Course.Edges.QuestionBank.ID
				if !app2.IsContainInt(quesIds, v.Edges.Course.Edges.QuestionBank.ID) {
					quesIds = append(quesIds, v.Edges.Course.Edges.QuestionBank.ID)
					otherList = append(otherList, &o)
				}
			}
		}
	}

	//查询题库user记录
	/*if len(quesIds) > 0 {
		userList, err := app.NewUserRecode().GetUserRecodeByIds(ctx, quesIds, uid.(int))
		if err == nil {
			for _, v := range list {
				for _, c := range v.ClassCourse {
					if _, ok := userList[c.QuestionBankId]; ok {
						recount := userList[c.QuestionBankId].RecordCount
						v.RecodeNum = v.RecodeNum + recount
					}
				}
			}
		}
	}*/

	if len(list) > 0 || len(otherList) > 0 {
		var data = map[string]interface{}{
			"class": list,
			"other": otherList,
		}
		return data, nil
	}

	return nil, nil
}

//tk专业题库标签
func GetQuestionBankTag(ctx *gin.Context) (interface{}, error) {
	ca := cache.CommCache
	Key := cache.QuestionBankTag
	if data, ok := ca.Get(Key); ok {
		return data, nil
	}

	s := store.WithContext(ctx)
	leveList := s.Level.Query().SoftDelete().Where(level.Status(2)).Order(ent.Asc(level.FieldSortOrder)).AllX(ctx)
	majorList := s.Major.Query().SoftDelete().Where(major.Status(2)).Order(ent.Asc(major.FieldSortOrder)).AllX(ctx)
	itemList := s.ItemCategory.Query().SoftDelete().Where(itemcategory.Status(2)).Order(ent.Asc(itemcategory.FieldSortOrder)).AllX(ctx)

	if len(leveList) == 0 && len(majorList) == 0 && len(itemList) == 0 {
		return nil, nil
	}

	res := response.QuestionMajorCheck{
		LevelList:        []response.LevelList{},
		MajorList:        []response.MajorList{},
		ItemCategoryList: []response.ItemCategoryList{},
	}

	for _, v := range leveList {
		var re response.LevelList
		re.LeveId = v.ID
		re.LeveName = v.Name
		res.LevelList = append(res.LevelList, re)
	}

	for _, v := range majorList {
		var re response.MajorList
		re.MajorId = v.ID
		re.MajorName = v.Name
		res.MajorList = append(res.MajorList, re)
	}

	for _, v := range itemList {
		var re response.ItemCategoryList
		re.ItemId = v.ID
		re.ItemName = v.Name
		res.ItemCategoryList = append(res.ItemCategoryList, re)
	}

	ca.Set(Key, res)

	return res, nil
}

//题库筛选
func GetMajorQuestionBankList(ctx *gin.Context) (interface{}, error) {
	var req request.MajorsQuestionBank
	err := ctx.Bind(&req)
	if err != nil {
		return nil, errorno.NewParamErr(err)
	}

	ca := cache.CommCache
	Key := cache.QuestionBankCheckTag +strconv.Itoa(req.LevelId)+strconv.Itoa(req.MajorId)+strconv.Itoa(req.ItemCategoryId)+strconv.Itoa(req.CityId)
	if data, ok := ca.Get(Key); ok {
		return data, nil
	}
	s := store.WithContext(ctx)

	CityIds := s.TkQuestionBankCity.Query().SoftDelete().Where(tkquestionbankcity.CityID(req.CityId)).AllX(ctx)

	bankIds := []int{}
	questionCityIds := []int{}
	for _, v := range CityIds {
		questionCityIds = append(questionCityIds, v.QuestionBankID)
	}
	bankIds = questionCityIds

	if req.MajorId > 0 && req.CityId > 0 {
		majorListIds := s.TkQuestionBankMajor.Query().SoftDelete().Where(tkquestionbankmajor.MajorID(req.MajorId)).AllX(ctx)
		questionMajorIds := []int{}
		for _, v := range majorListIds {
			questionMajorIds = append(questionMajorIds, v.QuestionBankID)
		}

		bankIds = app2.SliceIntersect(questionCityIds, questionMajorIds)

	}

	if len(bankIds) == 0 {
		return nil, nil
	}

	query := s.TkQuestionBank.Query().SoftDelete().Where(tkquestionbank.Status(1), tkquestionbank.IDIn(bankIds...))

	if req.ItemCategoryId > 0 {
		query = query.Where(tkquestionbank.ItemCategoryID(req.ItemCategoryId))
	}

	if req.LevelId > 0 {
		query = query.Where(tkquestionbank.LevelID(req.LevelId))
	}

	bankList := query.AllX(ctx)

	list := []response.BankListInfo{}
	for _, v := range bankList {
		res := response.BankListInfo{}
		res.QuestionBankId = v.ID
		res.QuestionBankName = v.Name

		list = append(list, res)
	}

	ca.Set(Key,list)
	return list, nil
}

//学习圈子
func GetGroupCardList(ctx *gin.Context)(interface{},error){
	ca := cache.CommCache
	Key := cache.GroupCardList
	if data, ok := ca.Get(Key); ok {
		return data, nil
	}
	s := store.WithContext(ctx)

	list := s.GroupCard.Query().SoftDelete().Where(groupcard.Status(2)).WithAttachment(func(query *ent.AttachmentQuery) {
		query.SoftDelete().Select("id", "filename")
	}).Order(ent.Asc(groupcard.FieldSortOrder)).AllX(ctx)

	if len(list)==0{
		return nil,nil
	}

	res := []response.GroupCardList{}

	for _,v:=range list{
		re := response.GroupCardList{}
		re.Id = v.ID
		re.Name = v.Title
		re.Desc = v.Desc
		if v.Edges.Attachment != nil {
			re.CodeUrl = app2.GetOssHost() + v.Edges.Attachment.Filename
		}

		res = append(res,re)
	}
	ca.Set(Key,res)

	return res,nil
}


//课程下的题库信息
func GetCourseQuestionBankInfo(ctx *gin.Context) (interface{}, error) {
	var req request.CourseQuestionBank
	err := ctx.Bind(&req)
	if err != nil {
		return nil, errorno.NewParamErr(err)
	}

	//判断是否登录
	isLoginUid, err := app.Common{}.IsUserLogin(ctx)
	uid := isLoginUid
	if err != nil {
		uid  = 0
	}

	if req.QuestionId == 0 {
		c := app.TkQuestionBank{}
		bankIds, _ := c.GetCityBankIds(ctx, req.CityId)
		if len(bankIds) > 0 {
			req.QuestionId = bankIds[0].ID
		} else {
			return nil, nil
		}
	}

	s := store.WithContext(ctx)

	//查询课程题库
	TypeList := &response.QuestionBankType{
		TkExamTypeList: []response.TkExamType{},
	}
	quType := s.TkQuestionBank.Query().Where(tkquestionbank.ID(req.QuestionId)).WithExamQuestionTypes().FirstX(ctx)

	if quType != nil {
		TypeList.BankId = quType.ID
		TypeList.QuestionNum = quType.QuestionCount
		TypeList.QuestionBankName = quType.Name

		if uid > 0 {
			//查询用户做已刷题数
			userBankRecode, _ := s.TkUserQuestionBankRecord.Query().Where(tkuserquestionbankrecord.UserID(uid)).
				Where(tkuserquestionbankrecord.QuestionBankID(req.QuestionId)).First(ctx)
			if userBankRecode != nil {
				TypeList.UserStudyCount = userBankRecode.RecordCount
				TypeList.Accuracy = math.Ceil(userBankRecode.CorrectRate * 100)
				TypeList.FinishRate, _ = strconv.ParseFloat(fmt.Sprintf("%.1f", TypeList.UserStudyCount/TypeList.QuestionNum*100), 64)
				/*strconv.ParseFloat(strconv.FormatFloat(float64(TypeList.UserStudyCount)/float64(TypeList.QuestionNum)*100, 'f', 1, 64), 64)*/
				/*			TypeList.FinishRate= (userBankRecode.FinishRate) * 100
				 */
				totalRecode := TypeList.FinishRate + TypeList.Accuracy
				TypeList.ForecastScore = math.Ceil(totalRecode / 2)
			}
		}

		//用户学习在该题库总数
		var typeSpec = make([]response.TkExamType, 5)
		for i := 0; i < 5; i++ {
			typeSpec[i].Type = i + 1
			typeSpec[i].QuestionCount = 0
			typeSpec[i].TotalCount = 0
			typeSpec[i].UserCount = 0
		}
		var typeCountMap = make([]int, 6)

		if uid > 0 {
			qulist := s.TkUserQuestionRecord.Query().SoftDelete().Select("id", "question_id", "exam_question_type").Where(tkuserquestionrecord.UserID(uid), tkuserquestionrecord.QuestionBankID(req.QuestionId)).WithQuestion(func(query *ent.TkQuestionQuery) {
				query.SoftDelete().Where(tkquestion.PidIsNil())
			}).AllX(ctx)

			for _, v := range qulist {
				if v.Edges.Question != nil {
					switch int(v.ExamQuestionType) {
					case 1:
						typeCountMap[0] = typeCountMap[0] + 1
					case 2:
						typeCountMap[1] = typeCountMap[1] + 1
					case 3:
						typeCountMap[2] = typeCountMap[2] + 1
					case 4:
						typeCountMap[3] = typeCountMap[3] + 1
					case 5:
						typeCountMap[4] = typeCountMap[4] + 1
					default:
						break
					}
				}

			}
		}

		if len(quType.Edges.ExamQuestionTypes) > 0 {
			for _, v := range quType.Edges.ExamQuestionTypes {
				//，1：模拟考试，2：考点练习，3：历年真题，4：通关必做300题，5：考前密押卷
				switch int(v.ExamQuestionType) {
				case 1:
					typeSpec[0].TotalCount = typeCountMap[0]
					typeSpec[0].UserCount = v.AnsweredUserCount
					typeSpec[0].QuestionCount = typeSpec[0].QuestionCount + v.QuestionCount
					typeSpec[0].Type = int(v.ExamQuestionType)
				case 2:
					typeSpec[1].TotalCount = typeCountMap[1]
					typeSpec[1].UserCount = v.AnsweredUserCount
					typeSpec[1].QuestionCount = typeSpec[1].QuestionCount + v.QuestionCount
					typeSpec[1].Type = int(v.ExamQuestionType)
				case 3:
					typeSpec[2].TotalCount = typeCountMap[2]
					typeSpec[2].UserCount = v.AnsweredUserCount

					typeSpec[2].QuestionCount = typeSpec[2].QuestionCount + v.QuestionCount
					typeSpec[2].Type = int(v.ExamQuestionType)
				case 4:
					typeSpec[3].TotalCount = typeCountMap[3]
					typeSpec[3].UserCount = v.AnsweredUserCount

					typeSpec[3].QuestionCount = typeSpec[3].QuestionCount + v.QuestionCount
					typeSpec[3].Type = int(v.ExamQuestionType)
				case 5:
					typeSpec[4].TotalCount = typeCountMap[4]
					typeSpec[4].UserCount = v.AnsweredUserCount

					typeSpec[4].QuestionCount = typeSpec[4].QuestionCount + v.QuestionCount
					typeSpec[4].Type = int(v.ExamQuestionType)
				}
			}

			//模拟考试时间
			moNiInfo, _ := s.TkExamPaper.Query().Select("start_at", "end_at").Where(tkexampaper.ExamQuestionType(1), tkexampaper.QuestionBankID(quType.ID)).Where(tkexampaper.StartAtGT(time.Now())).Order(ent.Asc("start_at")).First(ctx)
			if moNiInfo != nil && !moNiInfo.StartAt.IsZero() {
				typeSpec[0].Time = moNiInfo.StartAt.Format("1月2日") + " " + moNiInfo.StartAt.Format("15:04") + "-" + moNiInfo.EndAt.Format("15:04")
			}

		}

		for _, v := range typeSpec {
			TypeList.TkExamTypeList = append(TypeList.TkExamTypeList, v)
		}

		return TypeList, nil
	}

	return nil, nil
}

//题库的试卷分类接口
func GetQuestionBankExamList(ctx *gin.Context) (interface{}, error) {
	var req request.QuestionBankExamList
	err := ctx.Bind(&req)
	if err != nil {
		return nil, errorno.NewParamErr(err)
	}
	s := store.WithContext(ctx)

	tk := app.TkQuestionBank{}
	uid, _ := ctx.Get("uid")
	if uid == nil {
		return nil, errorno.NewInternalErr(errorno.TokenError)
	}

	a := app.Common{}

	Now := time.Now().Unix()
	//试卷题目分类，1：模拟考试，2：考点练习，3：历年真题，4：通关必做300题，5：考前密押卷
	switch req.TypeId {

	case 2:
		//章节练习
		chapterList, err := s.TkChapter.Query().SoftDelete().
			Where(tkchapter.QuestionBankID(req.QuestionBankId)).
			WithSections(func(query *ent.TkSectionQuery) {
				query.SoftDelete().WithTkSectionLinks()
			}).All(ctx)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, errorno.NewParamErr(errorno.Errno{
					Code:    0,
					Message: "没有数据",
				})
			}
			return nil, errorno.NewInternalErr(err)
		}
		var list []*response.QuestionBankChapter

		var secIds []int

		for _, v := range chapterList {
			chapter := response.QuestionBankChapter{
				SecList: []response.QuestionBankSection{},
			}
			chapter.ChapterId = v.ID
			chapter.ChapTerName = v.Name
			for _, c := range v.Edges.Sections {
				sec := response.QuestionBankSection{
					QuestionIds: []int{},
				}
				sec.SecId = c.ID
				sec.SecName = c.Name
				sec.QuestionCount = c.QuestionCount
				for _, lv := range c.Edges.TkSectionLinks {
					sec.QuestionIds = append(sec.QuestionIds, lv.QuestionID)
				}
				chapter.SecList = append(chapter.SecList, sec)

				secIds = append(secIds, c.ID)
			}
			list = append(list, &chapter)
		}

		data, err := app.NewUserRecode().GetUserChapterRecode(ctx, uid.(int), secIds)

		for k, v := range list {
			for i, c := range v.SecList {
				if len(data) > 0 {
					list[k].SecList[i].UserQuestionCount = data[c.SecId]
					if _, ok := data[c.SecId]; ok {
						list[k].SecList[i].IsDo = 1
					}
				}
			}
		}
		return list, nil
	case 1, 3, 4, 5:
		//必做300列表和密押卷
		list, err := tk.GetExamTypeListUserRecode(ctx, req.TypeId, req.QuestionBankId, uid.(int))
		if err != nil {
			return nil, errorno.NewInternalErr(err)
		}
		var tgList []*response.QuestionBankExam

		for _, v := range list {
			m := &response.QuestionBankExam{}
			m.Id = v.ID
			m.BankId = v.QuestionBankID
			m.TypeId = int(v.ExamQuestionType)
			m.ExamUserStudyCount = v.AnsweredUserCount
			m.ExamName = v.Name
			m.QuestionCount = v.QuestionCount
			m.PracticeExamAt = v.Duration
			m.Score = v.Score
			m.ExamType = int(v.ExamType)
			m.PassScore = v.PassScore
			m.Difficulty = a.GetDifficultyMap(int(v.Difficulty))
			m.Desc = v.Desc
			if len(v.Edges.UserExamPapers) > 0 {
				op := v.Edges.UserExamPapers[0].WrongCount + v.Edges.UserExamPapers[0].RightCount
				if op > 0 {
					m.Accuracy, _ = strconv.ParseFloat(fmt.Sprintf("%.1f", v.Edges.UserExamPapers[0].RightCount*100/op), 64)
					/*strconv.ParseFloat(strconv.FormatFloat(float64(v.Edges.UserExamPapers[0].RightCount*100)/float64(op), 'f', 1, 64), 64)*/
				}
				m.UserQuestionCount = v.Edges.UserExamPapers[0].TotalCount
				m.OrderStatus = int(v.Edges.UserExamPapers[0].OrderStatus)
				m.IsDo = 1
			}
			if !v.StartAt.IsZero() {
				/*	montu := int(v.StartAt.Month())
					montuDay := v.StartAt.Day()
					montuHours := v.StartAt.Hour()
					montuMinute := v.StartAt.Minute()
					montuEndHours := v.EndAt.Hour()
					montuEndMinute := v.EndAt.Minute()
					if montu != 0 {
						m.Time = strconv.Itoa(montu) + "月" + strconv.Itoa(montuDay) + "日 " + strconv.Itoa(montuHours) + ":" +
							strconv.Itoa(montuMinute) + "-" + strconv.Itoa(montuEndHours) + ":" + strconv.Itoa(montuEndMinute)
					}*/
				m.Time = v.StartAt.Format("1月2日") + " " + v.StartAt.Format("15:04") + "-" + v.EndAt.Format("15:04")

			}
			//	ExamStart      = 1 //考试开始
			//	ExamEnd        = 2 //考试结束
			//	ExamViewResult = 3 //查看成绩
			//	ExamNot        = 4 //考试未开始
			if req.TypeId == 1 { //模拟考试
				if Now >= v.StartAt.Unix() && Now <= v.EndAt.Unix() {
					m.SimulationExamStatus = ExamStart
				} else if Now < v.StartAt.Unix() {
					m.SimulationExamStatus = ExamNot
				} else if Now > v.EndAt.Unix() {
					m.SimulationExamStatus = ExamEnd
				}
			}

			if len(v.Edges.UserExamPapers) > 0 {
				m.SimulationExamStatus = ExamViewResult
			}

			tgList = append(tgList, m)
		}
		return tgList, nil
	}
	return nil, nil
}

func GetSimulationExamStatus(ctx *gin.Context) (interface{}, error) {
	var req request.SimulationExamStatus
	err := ctx.Bind(&req)
	if err != nil {
		return nil, errorno.NewParamErr(err)
	}
	uid, _ := ctx.Get("uid")
	s := store.WithContext(ctx)
	data := s.TkExamPaper.Query().SoftDelete().Select("id", "start_at", "end_at").Where(tkexampaper.ID(req.ExamId)).WithUserExamPapers(func(query *ent.TkUserExamScoreRecordQuery) {
		query.SoftDelete().Select("id", "exam_paper_id").Where(tkuserexamscorerecord.UserID(uid.(int)))
	}).FirstX(ctx)
	Now := time.Now().Unix()

	res := response.GetSimulationExamStatus{}

	if data != nil && !data.StartAt.IsZero() {
		//	ExamStart      = 1 //考试开始
		//	ExamEnd        = 2 //考试结束
		//	ExamViewResult = 3 //查看成绩
		//	ExamNot        = 4 //考试未开始
		if Now >= data.StartAt.Unix() && Now <= data.EndAt.Unix() {
			res.SurplusTime = int(data.EndAt.Unix() - Now)
			res.SimulationExamStatus = ExamStart
		} else if Now < data.StartAt.Unix() {
			res.SimulationExamStatus = ExamNot
		} else if Now > data.EndAt.Unix() {
			res.SimulationExamStatus = ExamEnd
		}
		if len(data.Edges.UserExamPapers) > 0 {
			res.SimulationExamStatus = ExamViewResult
		}
	}

	return res, nil
}

//试卷子部分详情
func GetQuestionList(ctx *gin.Context) (interface{}, error) {
	var req request.ExamQuestionList
	err := ctx.Bind(&req)
	if err != nil {
		return nil, errorno.NewParamErr(err)
	}
	uid, _ := ctx.Get("uid")
	ca := cache.CommCache
	smsKey := cache.ExamPartition + strconv.Itoa(req.Id)
	if data, ok := ca.Get(smsKey); ok {
		return data, nil
	}
	res := []response.ExamPaperPartition{}

	tk := app.TkQuestionBank{}

	/*	a := app.Common{}
	 */err = store.WithTx(ctx, func(ctx context.Context) error {
		s := store.WithContext(ctx)

		err := tk.TkExamPaperIdExist(ctx, req.Id)
		if err != nil {
			return err
		}

		if req.AnswerAgain == 1 {
			tk.DelUserExam(ctx, req.Id, uid.(int), 1)
		}

		query, err := s.TkExamPaper.Query().SoftDelete().Where(tkexampaper.ID(req.Id)).
			WithExamPartitions(func(query *ent.TkExamPaperPartitionQuery) {
				query.SoftDelete().WithExamPartitionLinks(func(query *ent.TkExamPartitionQuestionLinkQuery) {
					query.SoftDelete().WithQuestion(func(query *ent.TkQuestionQuery) {
						query.SoftDelete().WithCollectionQuestion(func(query *ent.CollectionQuery) {
							query.SoftDelete().Select("id", "value_id").Where(collection.UserID(uid.(int)), collection.Type(1))
						}).WithChildren(func(query *ent.TkQuestionQuery) {
							query.SoftDelete().WithCollectionQuestion(func(query *ent.CollectionQuery) {
								query.SoftDelete().Select("id", "value_id").Where(collection.UserID(uid.(int)), collection.Type(1))
							})
						})
					})
				}).WithExamPartitionScores(func(query *ent.TkExamPaperPartitionScoreQuery) {
					query.SoftDelete()
				})
			}).First(ctx)

		if err != nil {
			if err == sql.ErrNoRows {
				return nil
			}
			return err
		}

		suiJiQuestionIds := [6][]response.QuestionIdsCollection{}
		var userRandIds []int

		if int(query.ExamType) == SuiJiJuan {
			query1 := s.TkQuestion.Query().SoftDelete()
			if req.LookReport == "look_report" {
				randList := s.TkUserRandomExamRecode.Query().Where(tkuserrandomexamrecode.ExamID(req.Id), tkuserrandomexamrecode.UserID(uid.(int))).WithRandomExamQuestion(func(query1 *ent.TkQuestionQuery) {
					query1.Where(tkquestion.PidIsNil()).SoftDelete()
				}).FirstX(ctx)

				for _, v := range randList.Edges.RandomExamQuestion {
					userRandIds = append(userRandIds, v.ID)
				}

				if len(userRandIds) > 0 {
					query1 = query1.Where(tkquestion.IDIn(userRandIds...))
				}

			} else {
				query1 = query1.Where(tkquestion.QuestionBankID(query.QuestionBankID), tkquestion.PidIsNil())
			}

			list := query1.Select("id", "type").WithChildren(func(query *ent.TkQuestionQuery) {
				query.SoftDelete().WithCollectionQuestion(func(query *ent.CollectionQuery) {
					query.SoftDelete().Where(collection.UserID(uid.(int)), collection.Type(1))
				})
			}).WithCollectionQuestion(func(query *ent.CollectionQuery) {
				query.SoftDelete().Select("id", "value_id").Where(collection.UserID(uid.(int)), collection.Type(1))
			}).AllX(ctx)

			for _, v := range list {
				re := response.QuestionIdsCollection{
					MaterialsQuestion: []response.QuestionIdsCollection{},
				}
				re.Id = v.ID
				re.Type = int(v.Type)
				if len(v.Edges.CollectionQuestion) > 0 {
					re.IsCollection = 1
				}
				if len(v.Edges.Children) > 0 {

					for _, v1 := range v.Edges.Children {
						r2 := response.QuestionIdsCollection{}
						r2.Id = v1.ID
						r2.Type = int(v1.Type)
						if len(v1.Edges.CollectionQuestion) > 0 {
							r2.IsCollection = 1
						}
						re.MaterialsQuestion = append(re.MaterialsQuestion, r2)
					}
				}

				suiJiQuestionIds[int(v.Type)] = append(suiJiQuestionIds[int(v.Type)], re)
			}

		}

		if !app2.IsNil(query.Edges.ExamPartitions) {
			for _, v := range query.Edges.ExamPartitions {
				if v.QuestionCount <= 0 {
					continue
				}
				examPaperPartition := response.ExamPaperPartition{
					QuestionIDs: []response.QuestionIdsCollection{},
				}
				examPaperPartition.ID = v.ID
				examPaperPartition.Name = v.Name
				examPaperPartition.Type = v.Type

				examPaperPartition.Duration = v.Duration
				examPaperPartition.ExamPaperID = v.ExamPaperID
				examPaperPartition.QuestionCount = v.QuestionCount
				examPaperPartition.CreatedAt = v.CreatedAt

				if int(query.ExamType) == SuiJiJuan { //随机卷
					count := 0
					if len(suiJiQuestionIds[v.Type]) > 0 {
						if len(suiJiQuestionIds[v.Type]) < int(v.QuestionCount) {
							count = len(suiJiQuestionIds[v.Type])
						} else {
							count = int(v.QuestionCount)
						}
						//pc端查看答题卡
						/*						if req.LookReport == "look_report"{
												examPaperPartition.QuestionIDs = suiJiQuestionIds
											}*/

						quids := tk.MicsSlice(suiJiQuestionIds[v.Type], count)
						if len(quids) > 0 {
							examPaperPartition.QuestionIDs = quids
						} else {
							examPaperPartition.QuestionIDs = []response.QuestionIdsCollection{}
						}
					}
					//随机卷取默认表分数
					if len(v.Edges.ExamPartitionScores) > 0 {
						score, count, _ := tk.GetExamPartiTypeScore(v.Type, v.Edges.ExamPartitionScores[0])
						examPaperPartition.PartitionScore = score
						examPaperPartition.PartitionCount = count
						examPaperPartition.PartitionTotalScore = score * count
					}

					examPaperPartition.Desc = tk.GetExamPartiTypeDesc(v.Type, examPaperPartition.PartitionScore, examPaperPartition.PartitionCount)

				} else {

					if !app2.IsNil(v.Edges.ExamPartitionLinks) {

						for _, lv := range v.Edges.ExamPartitionLinks {
							re := response.QuestionIdsCollection{
								MaterialsQuestion: []response.QuestionIdsCollection{},
							}
							if !app2.IsNil(lv.Edges.Question) {
								//固定卷分数文案

								examPaperPartition.PartitionScore = examPaperPartition.PartitionScore + int(lv.QuestionScore)
								examPaperPartition.PartitionTotalScore = examPaperPartition.PartitionScore

								re.Id = lv.Edges.Question.ID
								re.Type = int(lv.Edges.Question.Type)
								if len(lv.Edges.Question.Edges.CollectionQuestion) > 0 {
									re.IsCollection = 1
								}

								if len(lv.Edges.Question.Edges.Children) > 0 {
									for _, v1 := range lv.Edges.Question.Edges.Children {
										r2 := response.QuestionIdsCollection{}
										r2.Id = v1.ID
										r2.Type = int(v1.Type)
										if len(v1.Edges.CollectionQuestion) > 0 {
											r2.IsCollection = 1
										}
										re.MaterialsQuestion = append(re.MaterialsQuestion, r2)
									}
								}
								examPaperPartition.QuestionIDs = append(examPaperPartition.QuestionIDs, re)
							}
						}

						examPaperPartition.Desc = tk.GetGuDingExamPartiTypeDesc(v.Type, examPaperPartition.PartitionScore, int(v.QuestionCount))

					}
				}

				res = append(res, examPaperPartition)
			}
		}

		//做题记录

		if req.QuestionUserAction > 0 {
			ma, _ := s.MakeUserQuestionRecord.Query().Where(makeuserquestionrecord.ExamID(req.Id), makeuserquestionrecord.UserID(uid.(int)), makeuserquestionrecord.QuestionBankID(query.QuestionBankID)).First(ctx)
			if ma == nil {
				_, err = s.MakeUserQuestionRecord.Create().SetUserID(uid.(int)).SetExamID(req.Id).SetExamQuestionType(int(query.ExamQuestionType)).SetQuestionBankID(query.QuestionBankID).Save(ctx)
			} else {
				_, err = s.MakeUserQuestionRecord.UpdateOneID(ma.ID).SetUpdatedAt(time.Now()).Save(ctx)
			}

			if err != nil {
				return err
			}

		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	if len(res) > 0 {
		ca.Set(smsKey, res)
		return res, nil
	}
	ca.Set(smsKey, nil)
	return res, nil
}

//获取章节下题目列表
func GetTkCsQuestionList(ctx *gin.Context) (interface{}, error) {
	var req request.TkCsQuestionsList
	err := ctx.Bind(&req)
	if err != nil {
		return nil, err
	}
	uid, _ := ctx.Get("uid")
	ca := cache.CommCache
	cacheKey := cache.SecQuestionList + strconv.Itoa(*req.SectionId)

	if data, ok := ca.Get(cacheKey); ok {
		return data, nil
	}
	res := []*response.TkSecList{}
	err = store.WithTx(ctx, func(ctx context.Context) error {
		s := store.WithContext(ctx)
		tk := app.TkQuestionBank{}

		if req.AnswerAgain == 1 {
			tk.DelUserExam(ctx, *req.SectionId, uid.(int), 2)
		}

		//做题记录
		if req.QuestionUserAction > 0 {
			ma, _ := s.MakeUserQuestionRecord.Query().Where(makeuserquestionrecord.SecID(*req.SectionId), makeuserquestionrecord.UserID(uid.(int)), makeuserquestionrecord.QuestionBankID(req.QuestionBankId)).First(ctx)
			if ma == nil {
				_, err = s.MakeUserQuestionRecord.Create().SetUserID(uid.(int)).SetSecID(*req.SectionId).SetExamQuestionType(app.SectionLianXi).SetQuestionBankID(req.QuestionBankId).Save(ctx)
			} else {
				_, err = s.MakeUserQuestionRecord.UpdateOneID(ma.ID).SetUpdatedAt(time.Now()).Save(ctx)
			}

			if err != nil {
				return err
			}

		}

		query := s.TkQuestionSection.Query().SoftDelete().
			Where(tkquestionsection.SectionID(*req.SectionId))
		////列表
		list, err := query.
			SoftDelete().WithSectionQuestion(func(query *ent.TkQuestionQuery) {
			query.SoftDelete().Where(tkquestion.PidIsNil()).WithCollectionQuestion(func(query *ent.CollectionQuery) {
				query.SoftDelete().Where(collection.UserID(uid.(int)))
			}).WithChildren(func(query *ent.TkQuestionQuery) {
				query.SoftDelete()
			})
		}).Order(func(s *sql2.Selector) {
			t := sql2.Table(tkquestion.Table)
			s.Join(t).On(s.C(tkquestionsection.FieldQuestionID), t.C(tkquestion.FieldID))
			s.OrderBy(t.C(sql2.Asc(tkquestion.FieldType)))
		}).All(ctx)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil
			}
			return err
		}

		for _, v := range list {
			resSpe := &response.TkSecList{
				Materials: []response.TkSecList{},
			}
			if !app2.IsNil(v.Edges.SectionQuestion) {
				resSpe.Id = v.Edges.SectionQuestion.ID
				resSpe.Type = int(v.Edges.SectionQuestion.Type)
				if len(v.Edges.SectionQuestion.Edges.CollectionQuestion) > 0 {
					resSpe.IsCollection = 1
				}
			}

			materialDetail := response.TkSecList{
				Materials: []response.TkSecList{},
			}
			if !app2.IsNil(v.Edges.SectionQuestion.Edges.Children) {
				for _, cv := range v.Edges.SectionQuestion.Edges.Children {
					materialDetail.Id = cv.ID
					materialDetail.Type = int(cv.Type)
					resSpe.Materials = append(resSpe.Materials, materialDetail)
				}
			}
			res = append(res, resSpe)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	if len(res) == 0 {
		ca.Set(cacheKey, nil)

		return nil, nil
	}
	ca.Set(cacheKey, res)

	return res, nil
}

func AdRrr(ctx *gin.Context) (interface{}, error) {
	qulist, exists := ctx.GetPostFormArray("random_question_ids")
	fmt.Printf("%v,%v ,%v", ctx.Request.Form, qulist, exists)

	return nil, nil
}

//用户提交答题情况记录
func AddUserQuestionRecode(ctx *gin.Context) (interface{}, error) {
	var req request.UserRecodeList
	err := ctx.Bind(&req)
	randomIds := ctx.PostForm("random_question_ids")
	qulist := ctx.PostForm("question_list")

	if randomIds != "" {
		err = json.Unmarshal([]byte(randomIds), &req.RandomQuestionIds)
	}

	if qulist != "" {
		err = json.Unmarshal([]byte(qulist), &req.QuestionList)
	}

	if err != nil {
		return nil, err
	}

	uid, _ := ctx.Get("uid")
	if req.ExamId <= 0 && req.SecId <= 0 {
		return nil, errorno.NewParamErr(errorno.Errno{
			Code:    0,
			Message: "参数id不能为空",
		})
	}
	if len(req.QuestionList) > 0 {
		err = store.WithTx(ctx, func(ctx context.Context) error {
			s := store.WithContext(ctx)

			//做题记录统计用
			newQuestionIds := []int{} //统计
			for _, v := range req.QuestionList {
				newQuestionIds = append(newQuestionIds, v.QuestionId)
			}
			noNeedStatisticQuestionIds := []int{}
			if len(newQuestionIds) > 0 {
				userRecordedQuestions, err := s.TkUserQuestionRecord.Query().
					Where(tkuserquestionrecord.UserID(uid.(int))).
					Where(tkuserquestionrecord.QuestionBankID(req.QuestionBankId)).
					Where(tkuserquestionrecord.QuestionIDIn(newQuestionIds...)).
					Select("id", "question_id").
					All(ctx)
				if err != nil {
					log.Error("get question bank record error", err.Error())
				} else {
					for _, uv := range userRecordedQuestions {
						noNeedStatisticQuestionIds = append(noNeedStatisticQuestionIds, uv.QuestionID)
					}
				}
			}

			qy := s.TkUserQuestionRecord.Delete().Where(tkuserquestionrecord.UserID(uid.(int))).Where(tkuserquestionrecord.QuestionBankID(req.QuestionBankId))
			if req.ExamQuestionType == 2 {
				qy.Where(tkuserquestionrecord.SectionID(req.SecId))
			} else {
				qy.Where(tkuserquestionrecord.ExamPaperID(req.ExamId))
			}
			_, err = qy.Exec(ctx)

			if err != nil {
				return err
			}

			numb := make([]*ent.TkUserQuestionRecordCreate, len(req.QuestionList))

			wrongList := []request.UserRecodeAdd{}

			wrongListIds := make([]int, len(req.QuestionList))
			for i, v := range req.QuestionList {

				req.QuestionList[i].IsRecordCount = 1 //统计
				if len(noNeedStatisticQuestionIds) > 0 {
					if app2.IsContainInt(noNeedStatisticQuestionIds, v.QuestionId) {
						req.QuestionList[i].IsRecordCount = 2
					}

					/*	for _, nv := range noNeedStatisticQuestionIds {
						if v.QuestionId == nv {
							req.QuestionList[i].IsRecordCount = 2
							break
						}
					}*/
				}

				query := s.TkUserQuestionRecord.
					Create().
					SetAnswer(v.Answer).
					SetIsRight(uint8(v.IsRight))
				if req.ExamQuestionType == app.SectionLianXi {
					query.SetSectionID(req.SecId)
				} else {
					query.SetExamPaperID(req.ExamId)
				}
				query.SetExamQuestionType(uint8(req.ExamQuestionType)).
					SetQuestionBankID(req.QuestionBankId).
					SetQuestionType(uint8(v.QuestionType)).
					SetQuestionID(v.QuestionId).
					SetUserID(uid.(int))

				numb[i] = query

				if v.IsRight == 1 {
					wrongList = append(wrongList, v)
					wrongListIds = append(wrongListIds, v.QuestionId)
				}
			}
			_, err = s.TkUserQuestionRecord.CreateBulk(numb...).Save(ctx)

			if err != nil {
				return err
			}

			//跟新试卷成绩表
			id := 0
			if req.ExamQuestionType == app.SectionLianXi {
				id = req.SecId
			} else {
				id = req.ExamId
			}

			isExam, err := app.NewUserRecode().IsUserExamScoreRecode(ctx, id, req.ExamQuestionType, uid.(int))

			if err != nil {
				return err
			}

			//统计分数
			scoreTotal := 0
			if req.ExamQuestionType != app.SectionLianXi {
				examParIds := []int{} //试卷子ids
				scoreList := map[int]int{}
				exam := s.TkExamPaper.Query().SoftDelete().Where(tkexampaper.ID(req.ExamId)).WithExamPartitions().FirstX(ctx)

				if exam != nil && len(exam.Edges.ExamPartitions) > 0 {
					for _, v := range exam.Edges.ExamPartitions {
						if v != nil {
							examParIds = append(examParIds, v.ID)
						}
					}
					scoreInfo := s.TkExamPartitionQuestionLink.Query().SoftDelete().Where(tkexampartitionquestionlink.ExamPaperPartitionIDIn(examParIds...), tkexampartitionquestionlink.QuestionIDIn(newQuestionIds...)).AllX(ctx)

					for _, v := range scoreInfo {
						if v != nil {
							scoreList[v.QuestionID] = int(v.QuestionScore)
						}
					}

					for _, v := range req.QuestionList {
						if v.IsRight == 2 { //正确
							if _, ok := scoreList[v.QuestionId]; ok {
								scoreTotal += scoreList[v.QuestionId]
							}
						}
					}
				}
			}
			/*			noCount := req.TotalCount - (req.RightCount + req.WrongCount)
			 */
			if isExam {
				//有数据就跟新
				query := s.TkUserExamScoreRecord.Update()
				query.Where(tkuserexamscorerecord.UserID(uid.(int)))

				if req.ExamQuestionType == app.SectionLianXi {
					query.Where(tkuserexamscorerecord.SectionID(req.SecId)).SetObjectiveQuestionScore(uint8(req.Score))
				} else {
					query.Where(tkuserexamscorerecord.ExamPaperID(req.ExamId)).SetObjectiveQuestionScore(uint8(scoreTotal))
				}
				_, err := query.
					SetRightCount(req.RightCount).
					SetWrongCount(req.WrongCount).
					SetTotalCount(req.TotalCount).
					SetDuration(req.Duration).
					SetNoAnswerCount(req.NoCount).
					Save(ctx)

				if err != nil {
					return err
				}

			} else {
				query := s.TkUserExamScoreRecord.Create()

				query.SetRightCount(req.RightCount).
					SetWrongCount(req.WrongCount).
					SetTotalCount(req.TotalCount).
					SetDuration(req.Duration).
					SetNoAnswerCount(req.NoCount).
					SetUserID(uid.(int))
				if req.ExamQuestionType == app.SectionLianXi {
					query.SetSectionID(req.SecId).SetObjectiveQuestionScore(uint8(req.Score))
				} else {
					query.SetExamPaperID(req.ExamId).SetObjectiveQuestionScore(uint8(scoreTotal))
				}
				query.SaveX(ctx)
			}

			if len(wrongList) > 0 {
				//跟新错题记录表
				s.TkUserWrongQuestionRecode.Delete().Where(tkuserwrongquestionrecode.UserID(uid.(int)), tkuserwrongquestionrecode.QuestionIDIn(wrongListIds...), tkuserwrongquestionrecode.QuestionBankID(req.QuestionBankId)).ExecX(ctx)

				WrongNumb := make([]*ent.TkUserWrongQuestionRecodeCreate, len(wrongList))

				for k, v := range wrongList {
					WrongNumb[k] = s.TkUserWrongQuestionRecode.Create().SetAnswer(v.Answer).SetWrongExamType(req.ExamQuestionType).SetWrongQuestionType(v.QuestionType).
						SetUserID(uid.(int)).
						SetQuestionID(v.QuestionId).
						SetQuestionBankID(req.QuestionBankId)
				}
				s.TkUserWrongQuestionRecode.CreateBulk(WrongNumb...).SaveX(ctx)
			}

			//随机卷处理
			if req.ExamQuestionType != app.SectionLianXi {
				if req.ExamType == SuiJiJuan {
					s.TkUserRandomExamRecode.Delete().Where(tkuserrandomexamrecode.UserID(uid.(int)), tkuserrandomexamrecode.ExamID(req.ExamId)).ExecX(ctx)
					s.TkUserRandomExamRecode.Create().SetUserID(uid.(int)).SetExamID(req.ExamId).AddRandomExamQuestionIDs(req.RandomQuestionIds...).SaveX(ctx)
				}
			}
			return nil
		})
	}

	if err != nil {
		return nil, errorno.NewInternalErr(err)
	}

	//做题统计异步逻辑
	syncQuestionList, err := json.Marshal(req.QuestionList) //统计
	if err != nil {
		log.Error("set question bank record error", err.Error())
	} else {
		err = asynctask.Enqueue("user_question_record", map[string]interface{}{
			"user_id":          uid.(int),
			"question_bank_id": req.QuestionBankId,
			"question_list":    string(syncQuestionList),
		})
		if err != nil {
			log.Error("set question bank record param error", err.Error())
		}
	} //统计

	return nil, nil
}

//查看答题卡
func GetUserAnswerSheet(ctx *gin.Context) (interface{}, error) {
	var req request.UserAnswerSheet
	err := ctx.Bind(&req)
	if err != nil {
		return nil, err
	}
	uid, _ := ctx.Get("uid")
	var id int
	s := store.WithContext(ctx)
	query := s.TkUserExamScoreRecord.Query().Where(tkuserexamscorerecord.UserID(uid.(int)))
	if req.ExamQuestionType == app.SectionLianXi {
		id = req.SecId
		query.Where(tkuserexamscorerecord.SectionID(req.SecId))
	} else {
		id = req.ExamId
		query.Where(tkuserexamscorerecord.ExamPaperID(req.ExamId))
	}

	userExamInfo := query.FirstX(ctx)

	if userExamInfo == nil {
		return nil, nil
	}

	tk := app.TkQuestionBank{}

	qlist := []*ent.TkQuestion{}

	if req.ExamType == SuiJiJuan {
		q := s.TkUserRandomExamRecode.Query().Where(tkuserrandomexamrecode.ExamID(req.ExamId), tkuserrandomexamrecode.UserID(uid.(int))).WithRandomExamQuestion(func(query1 *ent.TkQuestionQuery) {
			query1.Where(tkquestion.PidIsNil()).SoftDelete().WithChildren()
		}).FirstX(ctx)
		if q != nil {
			qlist = q.Edges.RandomExamQuestion
		}
		if qlist == nil {
			return nil, nil
		}

	} else {
		qlist, _ = tk.GetQuestionTypeSeEList(ctx, id, req.ExamQuestionType)

		if qlist == nil {
			return nil, nil
		}
	}

	//查询题目
	query1 := s.TkUserQuestionRecord.Query().SoftDelete().Where(tkuserquestionrecord.UserID(uid.(int)))
	//试卷查询分数
	examParIds := []int{} //试卷子ids
	scoreList := map[int]int{}
	if req.ExamQuestionType == app.SectionLianXi {
		query1.Where(tkuserquestionrecord.SectionID(req.SecId))
	} else {
		query1.Where(tkuserquestionrecord.ExamPaperID(req.ExamId))
		//试卷查询分数
		exam := s.TkExamPaper.Query().SoftDelete().Where(tkexampaper.ID(req.ExamId)).WithExamPartitions().FirstX(ctx)

		if exam != nil && len(exam.Edges.ExamPartitions) > 0 {
			for _, v := range exam.Edges.ExamPartitions {
				if v != nil {
					examParIds = append(examParIds, v.ID)
				}
			}
			scoreInfo := s.TkExamPartitionQuestionLink.Query().SoftDelete().Where(tkexampartitionquestionlink.ExamPaperPartitionIDIn(examParIds...)).AllX(ctx)

			for _, v := range scoreInfo {
				if v != nil {
					scoreList[v.QuestionID] = int(v.QuestionScore)
				}
			}
		}
	}

	quList, _ := query1.All(ctx)

	res := response.UserAnswerSheet{
		QuestionList: []response.UserQuestionTypeList{},
	}

	res.Score = int(userExamInfo.ObjectiveQuestionScore)
	res.WrongCount = userExamInfo.WrongCount
	res.RightCount = userExamInfo.RightCount
	res.NoCount = userExamInfo.NoAnswerCount
	res.QuestionTotalCount = userExamInfo.TotalCount
	res.Duration = userExamInfo.Duration
	res.Time = userExamInfo.CreatedAt.Format("2006-01-02 15:04")
	if userExamInfo.TotalCount > 0 {
		op := userExamInfo.RightCount + userExamInfo.WrongCount
		if op > 0 {
			res.Accuracy, _ = strconv.ParseFloat(fmt.Sprintf("%.1f", userExamInfo.RightCount*100/op), 64)
		}
	}

	var qulistSep = make(map[int]*ent.TkUserQuestionRecord)

	if len(quList) > 0 {
		for _, v := range quList {
			qulistSep[v.QuestionID] = v
		}
	}

	if len(qlist) > 0 {
		for _, v := range qlist {
			r := response.UserQuestionTypeList{}
			r.Id = v.ID
			if _, ok := scoreList[v.ID]; ok {
				r.QuestionScore = scoreList[v.ID]
			}
			if ok := qulistSep[v.ID]; ok != nil && v.Type != 5 {
				r.IsRight = int(qulistSep[v.ID].IsRight)
				r.IsDo = 1
				r.UserAnswer = qulistSep[v.ID].Answer
			}

			r.QuestionType = int(v.Type)
			//材料题
			if v.Type == 5 {
				if len(v.Edges.Children) > 0 {
					for _, v1 := range v.Edges.Children {
						re := response.UserQuestionTypeList{}
						re.Id = v1.ID
						re.Pid = v1.Pid
						if _, ok := scoreList[v1.ID]; ok {
							r.QuestionScore = scoreList[v1.ID]
						}

						if ok := qulistSep[v1.ID]; ok != nil {
							r.IsDo = 1
							re.IsDo = 1
							re.UserAnswer = qulistSep[v1.ID].Answer
						}
						re.QuestionType = int(v1.Type)
						res.QuestionList = append(res.QuestionList, re)
					}

				}

			}
			res.QuestionList = append(res.QuestionList, r)

		}
	}

	return res, nil

}

//错题本首页
func GetUserWrongQuestionList(ctx *gin.Context) (interface{}, error) {
	var req request.OnlyQuestionBankId

	err := ctx.Bind(&req)
	if err != nil {
		return nil, err
	}

	uid, _ := ctx.Get("uid")

	s := store.WithContext(ctx)

	list, err := s.TkUserWrongQuestionRecode.Query().Where(tkuserwrongquestionrecode.UserID(uid.(int)), tkuserwrongquestionrecode.WrongExamTypeGT(1), tkuserwrongquestionrecode.QuestionBankID(req.QuestionBankId)).All(ctx)

	if err != nil {
		return nil, err
	}

	data := []response.UserTypeQuestion{}

	//用户学习在该题库总数
	var typeSpec = make([]response.UserTypeQuestion, 4)
	for i := 0; i < 4; i++ {
		typeSpec[i].ExamQuestionType = i + 2
		typeSpec[i].Count = 0
		typeSpec[i].QuestionIds = []response.Question{}
	}
	//，1：模拟考试，2：考点练习，3：历年真题，4：通关必做300题，5：考前密押卷
	for _, v := range list {
		typeSpec[v.WrongExamType-2].ExamQuestionType = v.WrongExamType
		typeSpec[v.WrongExamType-2].Count = typeSpec[v.WrongExamType-2].Count + 1
		uw := response.Question{}
		uw.Id = v.QuestionID
		uw.Type = v.WrongQuestionType
		typeSpec[v.WrongExamType-2].QuestionIds = append(typeSpec[v.WrongExamType-2].QuestionIds, uw)
	}

	for _, v := range typeSpec {

		data = append(data, v)
	}

	return data, nil

}

//错题列表
func GetWrongList(ctx *gin.Context) (interface{}, error) {
	var req request.UserWrongQuestionList
	err := ctx.Bind(&req)
	if err != nil {
		return nil, err
	}
	uid, _ := ctx.Get("uid")

	s := store.WithContext(ctx)

	dataList, _ := s.TkUserWrongQuestionRecode.Query().Where(tkuserwrongquestionrecode.UserID(uid.(int)), tkuserwrongquestionrecode.WrongExamType(req.ExamQuestionType), tkuserwrongquestionrecode.QuestionBankID(req.QuestionBankId)).All(ctx)
	com := app.Common{}
	data := make(map[int]*ent.TkUserWrongQuestionRecode)

	if len(dataList) > 0 {
		ids := make([]int, len(dataList))

		for _, v := range dataList {
			ids = append(ids, v.QuestionID)

			data[v.QuestionID] = v
		}

		res := []response.TkQuestionUserByIdDetail{}

		list, _ := s.TkQuestion.Query().Where(tkquestion.IDIn(ids...)).Where(tkquestion.PidIsNil()).SoftDelete().WithAnswerOptions().WithKnowledgePoints().
			WithChildren(func(query *ent.TkQuestionQuery) {
				query.SoftDelete().WithKnowledgePoints().WithAnswerOptions()
			}).All(ctx)

		if len(list) > 0 {
			for _, v := range list {
				resSpe := &response.TkQuestionUserByIdDetail{
					Materials: []response.TkQuestionDetail{},
				}
				resSpe.Id = v.ID
				resSpe.Name = v.Name
				resSpe.Desc = v.Desc
				resSpe.UserAnswer = data[v.ID].Answer //用户的上一次答案
				resSpe.Type = int(v.Type)
				resSpe.Difficulty = com.GetDifficultyMap(int(v.Difficulty))
				if len(v.Edges.KnowledgePoints) > 0 {
					resSpe.KnowledgePoints = v.Edges.KnowledgePoints[0].Name
				}
				resSpe.Options = ent.TkQuestionAnswerOptions{}

				materialDetail := response.TkQuestionDetail{}
				if !app2.IsNil(v.Edges.Children) {
					for _, cv := range v.Edges.Children {
						materialDetail.Id = cv.ID
						materialDetail.Name = cv.Name
						materialDetail.Desc = cv.Desc
						materialDetail.Type = int(cv.Type)
						/*					materialDetail.Score = int(v.QuestionScore)
						 */materialDetail.Difficulty = com.GetDifficultyMap(int(cv.Difficulty))
						/*						materialDetail.KnowledgePoints = ent.TkKnowledgePoints{}
						 */materialDetail.Options = ent.TkQuestionAnswerOptions{}
						if len(cv.Edges.KnowledgePoints) > 0 {
							materialDetail.KnowledgePoints = cv.Edges.KnowledgePoints[0].Name
						}
						if !app2.IsNil(cv.Edges.AnswerOptions) {
							materialDetail.Options = cv.Edges.AnswerOptions
						}
						resSpe.Materials = append(resSpe.Materials, materialDetail)
					}
				}
				res = append(res, *resSpe)
			}
			return res, nil
		}
	}
	return nil, nil
}

//通过id查题目
func GetQuestionDetail(ctx *gin.Context) (interface{}, error) {
	var req request.QuestionDetail
	err := ctx.Bind(&req)
	if err != nil {
		return nil, err
	}
	a := app.Common{}
	s := store.WithContext(ctx)
	ca := cache.QuestionCache
	smsKey := cache.GetQuestionKey + strconv.Itoa(req.QuestionId)
	if data, ok := ca.Get(smsKey); ok {
		return data, nil
	}
	info := s.TkQuestion.Query().Where(tkquestion.ID(req.QuestionId)).Where(tkquestion.PidIsNil()).WithKnowledgePoints().WithAnswerOptions().WithChildren(func(query *ent.TkQuestionQuery) {
		query.SoftDelete().WithKnowledgePoints().WithAnswerOptions()
	}).FirstX(ctx)

	if info != nil {
		res := response.QuestionDetail{
			Materials: []response.TkListDetail{},
		}
		res.Type = int(info.Type)
		res.Id = info.ID
		res.Difficulty = a.GetDifficultyMap(int(info.Difficulty))

		Name, _ := htmltojson.HtmlToJson(info.Name)
		res.NameIos = Name
		Desc1, _ := htmltojson.HtmlToJson(info.Desc)
		res.DescIos = Desc1

		res.Name = info.Name
		res.Desc = info.Desc
		res.QuestionBankId = info.QuestionBankID
		res.Options = ent.TkQuestionAnswerOptions{}

		if len(info.Edges.KnowledgePoints) > 0 {
			res.KnowledgePoints = info.Edges.KnowledgePoints[0].Name
		}
		if !app2.IsNil(info.Edges.AnswerOptions) {
			res.Options = info.Edges.AnswerOptions
			res.OptionsIos = []response.AnswerListIos{}
			for _, v := range info.Edges.AnswerOptions {
				ar := response.AnswerListIos{}
				ar.QuestionID = v.QuestionID
				ar.IsRight = int(v.IsRight)
				ar.OptionName = v.OptionName
				ar.Content, _ = htmltojson.HtmlToJson(v.Content)
				if info.Type == 4 || info.Type == 5 {
					ar.OptionNameIos, _ = htmltojson.HtmlToJson(v.OptionName)
				}
				res.OptionsIos = append(res.OptionsIos, ar)
			}
		}

		materialDetail := response.TkListDetail{}
		if !app2.IsNil(info.Edges.Children) {
			for _, cv := range info.Edges.Children {
				materialDetail.Id = cv.ID
				materialDetail.Name = cv.Name
				NameIos, _ := htmltojson.HtmlToJson(cv.Name)
				DescIos, _ := htmltojson.HtmlToJson(cv.Desc)
				materialDetail.NameIos = NameIos
				materialDetail.Desc = cv.Desc
				materialDetail.DescIos = DescIos
				materialDetail.Type = int(cv.Type)
				materialDetail.Difficulty = a.GetDifficultyMap(int(cv.Difficulty))
				materialDetail.Options = ent.TkQuestionAnswerOptions{}
				if len(cv.Edges.KnowledgePoints) > 0 {
					materialDetail.KnowledgePoints = cv.Edges.KnowledgePoints[0].Name
				}
				if !app2.IsNil(cv.Edges.AnswerOptions) {
					materialDetail.Options = cv.Edges.AnswerOptions
					materialDetail.OptionsIos = []response.AnswerListIos{}
					for _, a := range cv.Edges.AnswerOptions {
						r := response.AnswerListIos{}
						r.QuestionID = a.QuestionID
						r.IsRight = int(a.IsRight)
						r.OptionName = a.OptionName
						r.Content, _ = htmltojson.HtmlToJson(a.Content)
						if info.Type == 4 || info.Type == 5 {
							r.OptionNameIos, _ = htmltojson.HtmlToJson(a.OptionName)
						}
						materialDetail.OptionsIos = append(materialDetail.OptionsIos, r)
					}
				}
				res.Materials = append(res.Materials, materialDetail)
			}
		}
		ca.Set(smsKey, res)
		return res, nil
	}
	ca.Set(smsKey, nil)
	return nil, nil
}

//根据题目ID组查询
func GetQuestionListDetail(ctx *gin.Context) (interface{}, error) {
	var req request.QuestionIdArrDetail
	err := ctx.Bind(&req)
	if err != nil {
		return nil, err
	}
	a := app.Common{}
	s := store.WithContext(ctx)
	list := s.TkQuestion.Query().Where(tkquestion.IDIn(req.QuestionIds...)).Where(tkquestion.PidIsNil()).WithKnowledgePoints().WithAnswerOptions().WithChildren(func(query *ent.TkQuestionQuery) {
		query.SoftDelete().WithKnowledgePoints().WithAnswerOptions()
	}).AllX(ctx)

	if list != nil && len(list) > 0 {
		data := []response.QuestionDetail{}
		for _, info := range list {
			res := response.QuestionDetail{
				Materials: []response.TkListDetail{},
			}
			res.Type = int(info.Type)
			res.Id = info.ID
			res.Difficulty = a.GetDifficultyMap(int(info.Difficulty))

			/*	Name, _ := htmltojson.HtmlToJson(info.Name)
				res.NameIos = Name
				Desc1, _ := htmltojson.HtmlToJson(info.Desc)
				res.DescIos = Desc1*/

			res.Name = info.Name
			res.Desc = info.Desc
			res.Options = ent.TkQuestionAnswerOptions{}

			if len(info.Edges.KnowledgePoints) > 0 {
				res.KnowledgePoints = info.Edges.KnowledgePoints[0].Name
			}
			if !app2.IsNil(info.Edges.AnswerOptions) {
				res.Options = info.Edges.AnswerOptions
				/*	res.OptionsIos = []response.AnswerListIos{}
					for _, v := range info.Edges.AnswerOptions {
						ar := response.AnswerListIos{}
						ar.QuestionID = v.QuestionID
						ar.IsRight = int(v.IsRight)
						ar.OptionName = v.OptionName
						ar.Content, _ = htmltojson.HtmlToJson(v.Content)
						if info.Type == 4 || info.Type == 5 {
							ar.OptionNameIos, _ = htmltojson.HtmlToJson(v.OptionName)
						}
						res.OptionsIos = append(res.OptionsIos, ar)
					}*/
			}

			materialDetail := response.TkListDetail{}
			if !app2.IsNil(info.Edges.Children) {
				for _, cv := range info.Edges.Children {
					materialDetail.Id = cv.ID
					materialDetail.Name = cv.Name
					/*	NameIos, _ := htmltojson.HtmlToJson(cv.Name)
						DescIos, _ := htmltojson.HtmlToJson(cv.Desc)
						materialDetail.NameIos = NameIos*/
					materialDetail.Desc = cv.Desc
					/*					materialDetail.DescIos = DescIos
					 */materialDetail.Type = int(cv.Type)
					materialDetail.Difficulty = a.GetDifficultyMap(int(cv.Difficulty))
					materialDetail.Options = ent.TkQuestionAnswerOptions{}
					if len(cv.Edges.KnowledgePoints) > 0 {
						materialDetail.KnowledgePoints = cv.Edges.KnowledgePoints[0].Name
					}
					if !app2.IsNil(cv.Edges.AnswerOptions) {
						materialDetail.Options = cv.Edges.AnswerOptions
						/*materialDetail.OptionsIos = []response.AnswerListIos{}
						for _, a := range cv.Edges.AnswerOptions {
							r := response.AnswerListIos{}
							r.QuestionID = a.QuestionID
							r.IsRight = int(a.IsRight)
							r.OptionName = a.OptionName
							r.Content, _ = htmltojson.HtmlToJson(a.Content)
							if info.Type == 4 || info.Type == 5 {
								r.OptionNameIos, _ = htmltojson.HtmlToJson(a.OptionName)
							}
							materialDetail.OptionsIos = append(materialDetail.OptionsIos, r)
						}*/
					}
					res.Materials = append(res.Materials, materialDetail)
				}
			}
			data = append(data, res)
		}

		return data, err
	}

	return nil, nil
}

//添加错题
func AddWrongQuestion(ctx *gin.Context) (interface{}, error) {
	var req request.AddWrongQuestion
	err := ctx.Bind(&req)
	if err != nil {
		return nil, err
	}

	uid, _ := ctx.Get("uid")

	s := store.WithContext(ctx).TkUserWrongQuestionRecode

	info, err := s.Query().Where(tkuserwrongquestionrecode.UserID(uid.(int)), tkuserwrongquestionrecode.QuestionID(req.QuestionId), tkuserwrongquestionrecode.QuestionBankID(req.QuestionBankId)).First(ctx)

	if info != nil {
		_, err := s.UpdateOneID(info.ID).SetAnswer(req.Answer).SetWrongExamType(req.ExamType).SetWrongQuestionType(req.QuestionType).SetQuestionBankID(req.QuestionBankId).Save(ctx)

		if err != nil {
			return nil, errorno.NewInternalErr(err)
		}
	} else {
		_, err := s.Create().SetAnswer(req.Answer).SetWrongExamType(req.ExamType).SetWrongQuestionType(req.QuestionType).
			SetUserID(uid.(int)).
			SetQuestionID(req.QuestionId).SetQuestionBankID(req.QuestionBankId).Save(ctx)
		if err != nil {
			return nil, errorno.NewInternalErr(err)
		}
	}

	return nil, nil

}

//试题纠错
func AddQuestionFeedBack(ctx *gin.Context) (interface{}, error) {
	var req request.AddFeedBack
	err := ctx.Bind(&req)
	if err != nil {
		return nil, err
	}

	s := store.WithContext(ctx)

	_, err2 := s.TkQuestionErrorFeedback.Create().
		SetQuestionID(req.QuestionId).
		SetUsername(req.UserName).
		SetPhone(req.Phone).
		SetErrorDesc(req.ErrorDesc).
		SetErrorType(uint8(req.ErrorType)).
		Save(ctx)

	if err2 != nil {
		return nil, errorno.NewInternalErr(err2)
	}

	return nil, nil
}

//收藏
func AddCollection(ctx *gin.Context) (interface{}, error) {
	var req request.AddCollection
	err := ctx.Bind(&req)
	if err != nil {
		return nil, err
	}
	uid, _ := ctx.Get("uid")
	err = store.WithTx(ctx, func(ctx context.Context) error {
		s := store.WithContext(ctx)

		isCollection := s.Collection.Query().SoftDelete().Where(collection.UserID(uid.(int)), collection.Type(req.TypeId), collection.ValueID(req.Id), collection.QuestionBankID(req.QuestionBankId)).FirstX(ctx)

		if isCollection == nil {

			query := s.Collection.Create().SetUserID(uid.(int)).SetType(req.TypeId).SetValueID(req.Id).SetExamQuestionType(req.ExamQuestionType).SetQuestionBankID(req.QuestionBankId)

			if req.ExamId > 0 {
				query.SetExamID(req.ExamId)
			} else if req.SecId > 0 {
				query.SetSecID(req.SecId)
			}

			query.SaveX(ctx)
		} else {
			s.Collection.DeleteOneID(isCollection.ID).ExecX(ctx)
		}

		return nil
	})

	if err != nil {
		return nil, errorno.NewInternalErr(err)
	}

	return nil, nil
}

//做题记录围度(试卷,章节)
func GetMakeUserQuestionRecode(ctx *gin.Context) (interface{}, error) {
	var req request.GetMakeUserQuestion
	err := ctx.Bind(&req)
	if err != nil {
		return nil, err
	}
	uid, _ := ctx.Get("uid")
	s := store.WithContext(ctx)

	currentTime := time.Now()

	oldTime := currentTime.AddDate(0, 0, -7)

	list, err := s.MakeUserQuestionRecord.Query().Where(makeuserquestionrecord.UpdatedAtGT(oldTime)).Where(makeuserquestionrecord.UserID(uid.(int))).Where(makeuserquestionrecord.QuestionBankID(req.QuestionBankId)).WithExamPaper(func(query *ent.TkExamPaperQuery) {
		query.SoftDelete().WithUserExamPapers(func(query *ent.TkUserExamScoreRecordQuery) {
			query.SoftDelete().Select("right_count", "wrong_count", "total_count", "section_id", "exam_paper_id").Where(tkuserexamscorerecord.UserID(uid.(int)))
		})
	}).WithSection(func(sectionQuery *ent.TkSectionQuery) {
		sectionQuery.SoftDelete().WithUserSectionExam(func(recordQuery *ent.TkUserExamScoreRecordQuery) {
			recordQuery.SoftDelete().Select("right_count", "wrong_count", "total_count", "section_id", "exam_paper_id").Where(tkuserexamscorerecord.UserID(uid.(int)))
		})
	}).All(ctx)

	if err != nil {
		return nil, err
	}

	res := []response.UserMakeQuestion{}
	if len(list) > 0 {
		for _, v := range list {
			re := response.UserMakeQuestion{}
			re.ExamQuestionType = v.ExamQuestionType
			if v.Edges.ExamPaper != nil {
				re.Name = v.Edges.ExamPaper.Name
				re.ExamType = int(v.Edges.ExamPaper.ExamType)
				re.ExamId = v.Edges.ExamPaper.ID
				re.QuestionCount = v.Edges.ExamPaper.QuestionCount
				re.PracticeExamAt = v.Edges.ExamPaper.Duration
				if len(v.Edges.ExamPaper.Edges.UserExamPapers) > 0 {
					re.UserQuestionCount = v.Edges.ExamPaper.Edges.UserExamPapers[0].TotalCount
					re.IsDo = 1

					op := v.Edges.ExamPaper.Edges.UserExamPapers[0].WrongCount + v.Edges.ExamPaper.Edges.UserExamPapers[0].RightCount
					if op > 0 {
						re.Accuracy, _ = strconv.ParseFloat(fmt.Sprintf("%.1f", v.Edges.ExamPaper.Edges.UserExamPapers[0].RightCount*100/op), 64)
						/*strconv.ParseFloat(strconv.FormatFloat(float64(v.Edges.ExamPaper.Edges.UserExamPapers[0].RightCount*100)/float64(op), 'f', 1, 64), 64)*/
					}
				}

			}
			if v.Edges.Section != nil {
				re.Name = v.Edges.Section.Name
				re.Chapter = v.Edges.Section.ChapterID
				re.SecId = v.Edges.Section.ID
				re.QuestionCount = v.Edges.Section.QuestionCount
				if len(v.Edges.Section.Edges.UserSectionExam) > 0 && v.Edges.Section.Edges.UserSectionExam[0] != nil {
					re.UserQuestionCount = v.Edges.Section.Edges.UserSectionExam[0].TotalCount
					re.IsDo = 1
					op := v.Edges.Section.Edges.UserSectionExam[0].WrongCount + v.Edges.Section.Edges.UserSectionExam[0].RightCount

					if op > 0 {
						re.Accuracy, _ = strconv.ParseFloat(fmt.Sprintf("%.1f", v.Edges.Section.Edges.UserSectionExam[0].RightCount*100/op), 64)
						/*strconv.ParseFloat(strconv.FormatFloat(float64(v.Edges.Section.Edges.UserSectionExam[0].RightCount*100)/float64(op), 'f', 1, 64), 64)*/
					}

				}

			}

			res = append(res, re)
		}
		return res, nil

	}

	return nil, nil
}

//收藏首页
func GetCollectionList(ctx *gin.Context) (interface{}, error) {
	var req request.OnlyQuestionBankId
	err := ctx.Bind(&req)
	if err != nil {
		return nil, err
	}
	uid, _ := ctx.Get("uid")

	s := store.WithContext(ctx)

	list, err := s.Collection.Query().Where(collection.UserID(uid.(int)), collection.ExamQuestionTypeNotNil(), collection.QuestionBankID(req.QuestionBankId)).WithQuestion(func(query *ent.TkQuestionQuery) {
		query.SoftDelete().Select("id", "type")
	}).All(ctx)

	if err != nil {
		return nil, err
	}

	data := []response.UserTypeQuestion{}

	//用户学习在该题库总数
	var typeSpec = make([]response.UserTypeQuestion, 5)
	for i := 0; i < 5; i++ {
		typeSpec[i].ExamQuestionType = i + 1
		typeSpec[i].Count = 0
		typeSpec[i].QuestionIds = []response.Question{}
	}
	//，1：模拟考试，2：考点练习，3：历年真题，4：通关必做300题，5：考前密押卷
	for _, v := range list {
		typeSpec[v.ExamQuestionType-1].ExamQuestionType = v.ExamQuestionType
		typeSpec[v.ExamQuestionType-1].Count = typeSpec[v.ExamQuestionType-1].Count + 1
		r := response.Question{}
		r.Id = v.ValueID
		r.Type = int(v.Edges.Question.Type)
		typeSpec[v.ExamQuestionType-1].QuestionIds = append(typeSpec[v.ExamQuestionType-1].QuestionIds, r)
	}

	for _, v := range typeSpec {

		data = append(data, v)
	}

	return data, nil
}

//错题记录移除
func RemoveWrongQuestion(ctx *gin.Context) (interface{}, error) {
	var req request.RemoveWrongQuestion
	err := ctx.Bind(&req)
	if err != nil {
		return nil, err
	}
	uid, _ := ctx.Get("uid")

	s := store.WithContext(ctx)

	s.TkUserWrongQuestionRecode.Delete().Where(tkuserwrongquestionrecode.UserID(uid.(int)), tkuserwrongquestionrecode.QuestionID(req.QuestionId), tkuserwrongquestionrecode.QuestionBankID(req.QuestionBankId)).ExecX(ctx)

	return nil, nil

}
