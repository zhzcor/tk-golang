package app

import (
	"context"
	"fmt"
	"strconv"
	"time"
	"tkserver/httpapi/admin/request"
	"tkserver/httpapi/appapi/response"
	"tkserver/internal/errorno"
	"tkserver/internal/store"
	"tkserver/internal/store/ent"
	"tkserver/internal/store/ent/kccourse"
	"tkserver/internal/store/ent/kccoursechapter"
	"tkserver/internal/store/ent/kccoursesection"
	"tkserver/internal/store/ent/kccoursesmallcategory"
	"tkserver/internal/store/ent/kccourseteacher"
	"tkserver/internal/store/ent/kcuserclass"
	"tkserver/internal/store/ent/kcusercourse"
	"tkserver/internal/store/ent/kcvideouploadtask"
	"tkserver/internal/store/ent/teacher"
	"tkserver/internal/store/ent/user"
	"tkserver/pkg/asynctask"
	"tkserver/pkg/baijiayun"
	app "tkserver/pkg/requestparam"
)

type Course struct {
}

type CourseFile struct {
	Id         int    `json:"id"`
	CourseId   int    `json:"course_id"`
	CourseName string `json:"course_name"`
	ImgCover   string `json:"img_cover"`
	Title      string `json:"title"`
	FileName   string `json:"file_name"`
	FileSize   int    `json:"file_size"` //字节b
	MimeType   string `json:"mime_type"` //文件类型
}

type CourseLiveType struct {
	SmallType  int `json:"small_type"`  //1 直播 2 回放 3 点播
	LiveStatus int `json:"live_status"` // 1.未开始，2:直播中，3：直播结束，4：回放
}

func NewCourse() *Course {
	return &Course{}
}

//添加课程
func (c Course) AddCourse(ctx context.Context, req request.SetCourse, operatorAdminId int) (int, error) {
	s := store.WithContext(ctx)
	fined, err := s.KcCourse.
		Query().SoftDelete().
		Where(kccourse.CourseName(*req.CourseName)).
		Exist(ctx)
	if err != nil {
		return 0, errorno.DataExistsError
	}
	if fined {
		return 0, errorno.NewErr(errorno.DataExistsError)
	}
	courseCreation := s.KcCourse.Create().
		SetCourseName(*req.CourseName).
		SetCourseType(uint8(*req.CourseType)).
		SetCreatedAdminID(operatorAdminId)
	if !app.IsNil(req.ItemId) && *req.ItemId > 0 {
		courseCreation.SetItemID(*req.ItemId)
	}
	if !app.IsNil(req.CityId) && *req.CityId > 0 {
		courseCreation.SetCityID(*req.CityId)
	}
	if !app.IsNil(req.CoursePrice) {
		courseCreation.SetCoursePrice(*req.CoursePrice)
	}
	if !app.IsNil(req.CoverAttachmentId) {
		courseCreation.SetCourseCoverImgID(*req.CoverAttachmentId)
	}
	if !app.IsNil(req.CourseDesc) {
		courseCreation.SetCourseDesc(*req.CourseDesc)
	}
	if !app.IsNil(req.OpenLiveStart) && !app.IsNil(req.OpenLiveDuration) {
		if *req.CourseType == 3 { //直播公开课
			smallCate, err := c.setOpenLive(ctx, req)
			if err != nil {
				return 0, err
			}
			courseCreation.AddCourseSmallCategorys(smallCate)
		}
	}

	if len(req.MajorIds) > 0 {
		courseCreation.AddMajorIDs(req.MajorIds...)
	}
	course, err := courseCreation.Save(ctx)
	if err != nil {
		return 0, errorno.NewStoreErr(err)
	}
	return course.ID, nil
}

//编辑课程
func (c Course) UpdateCourse(ctx context.Context, req request.SetCourse, operatorAdminId int) (int, error) {
	s := store.WithContext(ctx)
	fined, err := s.KcCourse.
		Query().SoftDelete().
		Where(kccourse.CourseName(*req.CourseName)).
		Where(kccourse.IDNEQ(*req.Id)).
		Exist(ctx)
	if err != nil {
		return 0, errorno.DataExistsError
	}
	if fined {
		return 0, errorno.NewErr(errorno.DataExistsError)
	}
	courseCreation := s.KcCourse.UpdateOneID(*req.Id).
		ClearMajor().
		SetCourseName(*req.CourseName).
		SetCourseType(uint8(*req.CourseType)).
		SetCreatedAdminID(operatorAdminId)
	if !app.IsNil(req.ItemId) && *req.ItemId > 0 {
		courseCreation.SetItemID(*req.ItemId)
	}
	if !app.IsNil(req.CityId) && *req.CityId > 0 {
		courseCreation.SetCityID(*req.CityId)
	}
	if !app.IsNil(req.CoursePrice) {
		courseCreation.SetCoursePrice(*req.CoursePrice)
	}
	if !app.IsNil(req.CoverAttachmentId) {
		courseCreation.SetCourseCoverImgID(*req.CoverAttachmentId)
	}
	if !app.IsNil(req.CourseDesc) {
		courseCreation.SetCourseDesc(*req.CourseDesc)
	}
	if !app.IsNil(req.OpenLiveStart) && !app.IsNil(req.OpenLiveDuration) {
		if *req.CourseType == 3 { //直播公开课
			smallCate, err := c.setOpenLive(ctx, req)
			if err != nil {
				return 0, err
			}
			courseCreation.ClearCourseSmallCategorys().AddCourseSmallCategorys(smallCate)
		}
	}
	if len(req.MajorIds) > 0 {
		courseCreation.AddMajorIDs(req.MajorIds...)
	}
	course, err := courseCreation.Save(ctx)
	if err != nil {
		return 0, err
	}
	return course.ID, err
}

func (c Course) setOpenLive(ctx context.Context, req request.SetCourse) (*ent.KcCourseSmallCategory, error) {
	s := store.WithContext(ctx)
	bjy := baijiayun.BjyCloud{}
	startAt := req.OpenLiveStart.Unix()
	duration := strconv.Itoa(*req.OpenLiveDuration) + "s"
	newDuration, err := time.ParseDuration(duration)
	if err != nil {
		return nil, err
	}
	endAt := req.OpenLiveStart.Add(newDuration).Unix()
	fmt.Println(newDuration, startAt, endAt)
	LiveRoom, err := bjy.CreateLiveRoom(*req.CourseName, int(startAt), int(endAt), 2, 0, 0)
	if err != nil {
		return nil, err
	}
	roomId := LiveRoom.RoomId
	intRoomId, err := strconv.Atoi(roomId)
	if err != nil {
		return nil, err
	}
	smallCate, err := s.KcCourseSmallCategory.Create().
		SetSmallName(*req.CourseName).SetTeachType(2).
		SetLiveSmallStart(*req.OpenLiveStart).
		SetLiveSmallTime(*req.OpenLiveDuration).
		SetType(3).
		SetLiveSmallStatus(1).
		SetLiveRoomID(intRoomId).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return smallCate, nil
}

//删除课程
func (c Course) DelCourse(ctx context.Context, id int) error {
	s := store.WithContext(ctx)
	if errorno := c.CourseIdExist(ctx, id); errorno != nil {
		return errorno
	}
	_, err := s.KcCourse.UpdateOneID(id).SoftDelete().Save(ctx)
	if err != nil {
		return errorno.NewStoreErr(err)
	}
	return nil
}

//设置课程状态
func (c Course) SetCourseStatus(ctx context.Context, id, status int) error {
	s := store.WithContext(ctx)
	if errorno := c.CourseIdExist(ctx, id); errorno != nil {
		return errorno
	}
	_, err := s.KcCourse.UpdateOneID(id).SetPushStatus(uint8(status)).Save(ctx)
	if err != nil {
		return errorno.NewStoreErr(err)
	}
	return nil
}

//id判断课程是否存着
func (c Course) CourseIdExist(ctx context.Context, id int) error {
	s := store.WithContext(ctx)
	fined, err := s.KcCourse.
		Query().
		Where(kccourse.ID(id)).
		Where(kccourse.DeletedAtIsNil()).
		Exist(ctx)
	if err != nil {
		return errorno.NewStoreErr(err)
	}
	if !fined {
		return errorno.NewErr(errorno.Errno{
			Code:    20001,
			Message: "课程id不存在",
		})
	}
	return nil
}

//id判断课程是否存着
func (c Course) TeacherIdExist(ctx context.Context, id int) error {
	s := store.WithContext(ctx)
	fined, err := s.Teacher.
		Query().
		Where(teacher.ID(id)).
		Where(teacher.DeletedAtIsNil()).
		Exist(ctx)
	if err != nil {
		return errorno.NewStoreErr(err)
	}
	if !fined {
		return errorno.NewErr(errorno.Errno{
			Code:    20001,
			Message: "老师id不存在",
		})
	}
	return nil
}

//添加课程老师
func (c Course) AddCourseTeacher(ctx context.Context, req request.SetCourseTeacher) error {
	s := store.WithContext(ctx)
	if errorno := c.CourseIdExist(ctx, *req.CourseId); errorno != nil {
		return errorno
	}
	if errorno := c.TeacherIdExist(ctx, *req.TeacherId); errorno != nil {
		return errorno
	}
	creation := s.KcCourseTeacher.Create().
		SetCourseID(*req.CourseId).SetTeacherID(*req.TeacherId)
	if !app.IsNil(req.ShowStatus) {
		creation.SetShowStatus(uint8(*req.ShowStatus))
	}
	if !app.IsNil(req.SortOrder) {
		creation.SetSortOrder(*req.SortOrder)
	}
	_, err := creation.Save(ctx)
	if err != nil {
		return errorno.NewStoreErr(err)
	}
	return nil
}

//设置课程教师状态
func (c Course) SetCourseTeacherStatus(ctx context.Context, id, status int) error {
	s := store.WithContext(ctx)
	_, err := s.KcCourseTeacher.UpdateOneID(id).SetShowStatus(uint8(status)).Save(ctx)
	if err != nil {
		return errorno.NewStoreErr(err)
	}
	return nil
}

//删除课程老师
func (c Course) RemoveCourseTeacher(ctx context.Context, req request.SetCourseTeacher) error {
	s := store.WithContext(ctx)
	fined, err := s.KcCourseTeacher.Query().
		Where(kccourseteacher.CourseID(*req.CourseId)).
		Where(kccourseteacher.TeacherID(*req.TeacherId)).
		Exist(ctx)
	if err != nil {
		return errorno.NewStoreErr(err)
	}
	if !fined {
		return errorno.NewErr(errorno.Errno{
			Code:    20001,
			Message: "课程老师不存在",
		})
	}
	_, err = s.KcCourseTeacher.Delete().
		Where(kccourseteacher.CourseID(*req.CourseId)).
		Where(kccourseteacher.TeacherID(*req.TeacherId)).
		Exec(ctx)
	if err != nil {
		return errorno.NewStoreErr(err)
	}
	return nil
}

//添加课程学员
func (c Course) AddCourseUser(ctx context.Context, req request.SetCourseUser) error {
	s := store.WithContext(ctx)
	if errorno := c.CourseIdExist(ctx, *req.CourseId); errorno != nil {
		return errorno
	}
	if errorno := c.UserIdExist(ctx, *req.UserId); errorno != nil {
		return errorno
	}
	if errorno := c.UserCourseIdExist(ctx, *req.CourseId, *req.UserId); errorno != nil {
		return errorno
	}
	_, err := s.KcUserCourse.Create().
		SetCourseID(*req.CourseId).
		SetUserID(*req.UserId).
		Save(ctx)
	if err != nil {
		return errorno.NewStoreErr(err)
	}
	_, err = s.KcCourse.UpdateOneID(*req.CourseId).AddPeopleNum(1).Save(ctx)
	if err != nil {
		return errorno.NewStoreErr(err)
	}
	return nil
}

//删除课程学员
func (c Course) RemoveCourseUser(ctx context.Context, req request.SetCourseUser) error {
	s := store.WithContext(ctx)
	fined, err := s.KcUserCourse.Query().
		Where(kcusercourse.CourseID(*req.CourseId)).
		Where(kcusercourse.UserID(*req.UserId)).
		Exist(ctx)
	if err != nil {
		return errorno.NewStoreErr(err)
	}
	if !fined {
		return errorno.NewErr(errorno.Errno{
			Code:    20001,
			Message: "课程学员不存在",
		})
	}
	_, err = s.KcUserCourse.Delete().
		Where(kcusercourse.CourseID(*req.CourseId)).
		Where(kcusercourse.UserID(*req.UserId)).
		Exec(ctx)
	if err != nil {
		return errorno.NewStoreErr(err)
	}
	_, err = s.KcCourse.UpdateOneID(*req.CourseId).AddPeopleNum(-1).Save(ctx)
	if err != nil {
		return errorno.NewStoreErr(err)
	}
	return nil
}

//id判断用户是否存着
func (c Course) UserIdExist(ctx context.Context, id int) error {
	s := store.WithContext(ctx)
	fined, err := s.User.
		Query().
		Where(user.ID(id)).
		Where(user.DeletedAtIsNil()).
		Exist(ctx)
	if err != nil {
		return errorno.NewStoreErr(err)
	}
	if !fined {
		return errorno.NewErr(errorno.Errno{
			Code:    20001,
			Message: "用户id不存在",
		})
	}
	return nil
}

//判断课程用户是否存着
func (c Course) UserCourseIdExist(ctx context.Context, courseId, userId int) error {
	s := store.WithContext(ctx)
	fined, err := s.KcUserCourse.
		Query().
		Where(kcusercourse.CourseID(courseId)).
		Where(kcusercourse.UserID(userId)).
		Exist(ctx)
	if err != nil {
		return errorno.NewStoreErr(err)
	}
	if fined {
		return errorno.NewErr(errorno.Errno{
			Code:    20001,
			Message: "用户已存在",
		})
	}
	return nil
}

//设置学员课程有效期
func (c Course) SetCourseUserValidity(ctx context.Context, req request.SetCourseUserValidity) (int, error) {
	s := store.WithContext(ctx)
	fined, err := s.KcUserCourse.
		Query().
		Where(kcusercourse.UserID(*req.UserId)).
		Where(kcusercourse.CourseID(*req.CourseId)).
		Exist(ctx)
	if err != nil {
		return 0, errorno.NewStoreErr(err)
	}
	if !fined {
		return 0, errorno.NewErr(errorno.DataNotExistsError)
	}
	_, err = s.KcUserCourse.Update().
		Where(kcusercourse.UserID(*req.UserId)).
		Where(kcusercourse.CourseID(*req.CourseId)).
		SetPeriodType(2).
		SetClosingDate(*req.ClosingDate).
		Save(ctx)
	if err != nil {
		return 0, errorno.NewStoreErr(err)
	}
	return 0, nil
}

//添加课程题库
func (c Course) AddCourseQuestionBank(ctx context.Context, req request.SetCourseQuestionBank) error {
	s := store.WithContext(ctx)
	if errorno := c.CourseIdExist(ctx, *req.CourseId); errorno != nil {
		return errorno
	}
	tk := TkQuestionBank{}
	if errorno := tk.TkQuestionBankIdExist(ctx, *req.QuestionBankId); errorno != nil {
		return errorno
	}
	_, err := s.KcCourse.UpdateOneID(*req.CourseId).
		SetQuestionBankID(*req.QuestionBankId).
		Save(ctx)
	if err != nil {
		return errorno.NewStoreErr(err)
	}
	return nil
}

//删除课程题库
func (c Course) RemoveCourseQuestionBank(ctx context.Context, req request.SetCourseQuestionBank) error {
	s := store.WithContext(ctx)
	if errorno := c.CourseIdExist(ctx, *req.CourseId); errorno != nil {
		return errorno
	}
	_, err := s.KcCourse.UpdateOneID(*req.CourseId).
		ClearQuestionBankID().
		Save(ctx)
	if err != nil {
		return errorno.NewStoreErr(err)
	}
	return nil
}

//编辑章
func (c Course) UpdateKcCourseChapter(ctx context.Context, id int, title string, courseId int) (int, error) {
	s := store.WithContext(ctx)
	fined, err := s.KcCourseChapter.
		Query().
		Where(kccoursechapter.Title(title)).
		Where(kccoursechapter.IDNEQ(id)).
		Exist(ctx)
	if err != nil {
		return 0, errorno.NewStoreErr(err)
	}
	if fined {
		return 0, errorno.NewErr(errorno.DataExistsError)
	}
	if errorno := c.CourseIdExist(ctx, courseId); errorno != nil {
		return 0, errorno
	}
	chapter, err := s.KcCourseChapter.UpdateOneID(id).
		SetTitle(title).
		SetCourseID(courseId).
		Save(ctx)
	if err != nil {
		return 0, errorno.NewStoreErr(err)
	}
	return chapter.ID, nil
}

//添加章
func (c Course) AddKcCourseChapter(ctx context.Context, title string, courseId int) (int, error) {
	s := store.WithContext(ctx)
	fined, err := s.KcCourseChapter.
		Query().
		Where(kccoursechapter.Title(title)).
		Where(kccoursechapter.CourseID(courseId)).
		Exist(ctx)
	if err != nil {
		return 0, errorno.NewStoreErr(err)
	}
	if fined {
		return 0, errorno.NewErr(errorno.DataExistsError)
	}
	if errorno := c.CourseIdExist(ctx, courseId); errorno != nil {
		return 0, errorno
	}
	chapter, err := s.KcCourseChapter.Create().
		SetTitle(title).
		SetCourseID(courseId).
		Save(ctx)
	if err != nil {
		return 0, errorno.NewStoreErr(err)
	}
	return chapter.ID, nil
}

//删除章
func (c Course) DelKcCourseChapter(ctx context.Context, id int) error {
	s := store.WithContext(ctx)
	if errorno := c.KcCourseChapterIdExist(ctx, id); errorno != nil {
		return errorno
	}
	err := s.TkChapter.DeleteOneID(id).Exec(ctx)
	if err != nil {
		return errorno.NewStoreErr(err)
	}
	return nil
}

//id判断题库章是否存在
func (c Course) KcCourseChapterIdExist(ctx context.Context, id int) error {
	s := store.WithContext(ctx)
	fined, err := s.KcCourseChapter.
		Query().
		Where(kccoursechapter.ID(id)).
		Exist(ctx)
	if err != nil {
		return errorno.NewStoreErr(err)
	}
	if !fined {
		return errorno.NewErr(errorno.Errno{
			Code:    20001,
			Message: "课程章不存在",
		})
	}
	return nil
}

//编辑节
func (c Course) UpdateKcCourseSection(ctx context.Context, id int, title string, chapterId int) (int, error) {
	s := store.WithContext(ctx)
	fined, err := s.KcCourseSection.
		Query().
		Where(kccoursesection.Title(title)).
		Where(kccoursesection.IDNEQ(id)).
		Exist(ctx)
	if err != nil {
		return 0, errorno.NewStoreErr(err)
	}
	if fined {
		return 0, errorno.NewErr(errorno.DataExistsError)
	}
	if errorno := c.KcCourseChapterIdExist(ctx, chapterId); errorno != nil {
		return 0, errorno
	}
	section, err := s.KcCourseSection.UpdateOneID(id).
		SetTitle(title).
		SetCourseChapterID(chapterId).
		Save(ctx)
	if err != nil {
		return 0, errorno.NewStoreErr(err)
	}
	return section.ID, nil
}

//添加节
func (c Course) AddKcCourseSection(ctx context.Context, title string, chapterId int) (int, error) {
	s := store.WithContext(ctx)
	fined, err := s.KcCourseSection.
		Query().
		Where(kccoursesection.Title(title)).
		Where(kccoursesection.CourseChapterID(chapterId)).
		Exist(ctx)
	if err != nil {
		return 0, errorno.NewStoreErr(err)
	}
	if fined {
		return 0, errorno.NewErr(errorno.DataExistsError)
	}
	if errorno := c.KcCourseChapterIdExist(ctx, chapterId); errorno != nil {
		return 0, errorno
	}
	section, err := s.KcCourseSection.Create().
		SetTitle(title).
		SetCourseChapterID(chapterId).
		Save(ctx)
	if err != nil {
		return 0, errorno.NewStoreErr(err)
	}
	return section.ID, nil
}

//删除节
func (c Course) DelKcCourseSection(ctx context.Context, id int) error {
	s := store.WithContext(ctx)
	if errorno := c.KcCourseSectionIdExist(ctx, id); errorno != nil {
		return errorno
	}
	err := s.TkSection.DeleteOneID(id).Exec(ctx)
	if err != nil {
		return errorno.NewStoreErr(err)
	}
	return nil
}

//id判断课程节是否存在
func (c Course) KcCourseSectionIdExist(ctx context.Context, id int) error {
	s := store.WithContext(ctx)
	fined, err := s.KcCourseSection.
		Query().
		Where(kccoursesection.ID(id)).
		Exist(ctx)
	if err != nil {
		return errorno.NewStoreErr(err)
	}
	if !fined {
		return errorno.NewErr(errorno.Errno{
			Code:    20001,
			Message: "课程节不存在",
		})
	}
	return nil
}

//id判断课时是否存在
func (c Course) KcSmallCategoryIdExist(ctx context.Context, id int) error {
	s := store.WithContext(ctx)
	fined, err := s.KcCourseSmallCategory.
		Query().
		Where(kccoursesmallcategory.ID(id)).
		Exist(ctx)
	if err != nil {
		return errorno.NewStoreErr(err)
	}
	if !fined {
		return errorno.NewErr(errorno.Errno{
			Code:    20001,
			Message: "课时不存在",
		})
	}
	return nil
}

//添加伪直播视频
func (c Course) AddFalseVideo(ctx context.Context, req request.SetFalseVideoRequest) (int, error) {
	s := store.WithContext(ctx)
	if errorno := c.CourseIdExist(ctx, *req.CourseId); errorno != nil {
		return 0, errorno
	}
	bjy := baijiayun.BjyCloud{}

	uploadUrl, err := bjy.GetOrderUploadUrl(*req.Title, 0)
	if err != nil {
		return 0, errorno.NewStoreErr(err)
	}

	attachmentUrl, err := c.GetAttachmentDetail(ctx, *req.AttachmentId)
	if err != nil {
		return 0, errorno.NewStoreErr(err)
	}
	//创建点播上传任务
	taskId, err := bjy.CreateUploadTask(ctx, uploadUrl.VideoId, *req.AttachmentId, *req.CourseId, 2, *req.Title, *req.VideoName)
	if err != nil {
		return 0, err
	}

	if !app.IsNil(req.Length) {
		_, err := s.KcVideoUploadTask.UpdateOneID(taskId).SetLength(*req.Length).Save(ctx)
		if err != nil {
			return 0, err
		}
	}
	err = asynctask.Enqueue("order_video_upload", map[string]interface{}{
		"id":         taskId,
		"video_url":  attachmentUrl,
		"upload_url": uploadUrl.UploadUrl,
		"filename":   *req.Title,
	})
	if err != nil {
		return 0, errorno.NewParamErr(err)
	}
	//bjy.OrderUpload(ctx, taskId, attachmentUrl, uploadUrl.UploadUrl, *req.Title)

	return taskId, nil
}

//编辑伪直播视频
func (c Course) UpdateFalseVideo(ctx context.Context, req request.SetFalseVideoRequest) error {
	s := store.WithContext(ctx)
	if errorno := c.CourseIdExist(ctx, *req.CourseId); errorno != nil {
		return errorno
	}
	bjy := baijiayun.BjyCloud{}

	task, err := s.KcVideoUploadTask.Query().SoftDelete().
		Where(kcvideouploadtask.ID(*req.Id)).
		First(ctx)
	if err != nil && !ent.IsNotFound(err) {
		return errorno.NewStoreErr(err)
	}
	if ent.IsNotFound(err) {
		return errorno.NewErr(errorno.DataNotExistsError)
	}

	err = bjy.UpdateVideoInfo(*req.Title, task.VideoID)
	if err != nil {
		return err
	}

	_, err = s.KcVideoUploadTask.UpdateOneID(*req.Id).
		SetTitle(*req.Title).
		Save(ctx)
	if err != nil {
		return err
	}
	return nil
}

//删除伪直播
func (c Course) DelFalseVideo(ctx context.Context, id int) error {
	s := store.WithContext(ctx)
	err := s.KcVideoUploadTask.UpdateOneID(id).SoftDelete().Exec(ctx)
	if err != nil {
		return errorno.NewStoreErr(err)
	}
	return nil
}

//获取课程目录
func (c Course) GetCourseChapterList(ctx context.Context, id int) (*[]response.ChapterDetail2, error) {
	res2 := []response.ChapterDetail2{}
	s := store.WithContext(ctx)

	result, err := s.KcCourse.Query().SoftDelete().
		Where(kccourse.ID(id)).WithCourseSmallCategorys(func(query *ent.KcCourseSmallCategoryQuery) {
		query.SoftDelete().Where(kccoursesmallcategory.ChapterIDIsNil(), kccoursesmallcategory.SectionIDIsNil()).
			Select("id", "small_name", "live_small_start", "live_small_time", "push_status", "type", "live_small_status", "live_room_id",
				"course_id", "section_id", "chapter_id", "attachment_count", "question_count", "exam_count", "homework_count", "courseware_attach_id").WithCsAttachment(func(query *ent.AttachmentQuery) {
			query.SoftDelete()
		})
	}).WithCourseChapters(func(query *ent.KcCourseChapterQuery) {
		query.WithCourseSmallChapters(func(query *ent.KcCourseSmallCategoryQuery) {
			query.SoftDelete().Select("id", "small_name", "live_small_start", "live_small_time", "push_status", "type", "live_small_status", "live_room_id",
				"course_id", "section_id", "chapter_id", "attachment_count", "question_count", "exam_count", "homework_count", "courseware_attach_id").WithCsAttachment(func(query *ent.AttachmentQuery) {
				query.SoftDelete()
			})
		}).WithChapterSections(func(query *ent.KcCourseSectionQuery) {
			query.WithCourseSmallSections(func(query *ent.KcCourseSmallCategoryQuery) {
				query.SoftDelete().Select("id", "small_name", "live_small_start", "live_small_time", "push_status", "type", "live_small_status", "live_room_id",
					"course_id", "section_id", "chapter_id", "attachment_count", "question_count", "exam_count", "homework_count", "courseware_attach_id").WithCsAttachment(func(query *ent.AttachmentQuery) {
					query.SoftDelete()
				})
			})
		})
	}).First(ctx)

	if err != nil {
		if ent.IsNotFound(err) {
			return nil, nil
		}
		return nil, err
	}

	for _, v := range result.Edges.CourseSmallCategorys {
		re := response.ChapterDetail2{
			List: []response.ChapterDetail2{},
		}
		re.Type = 3
		re.SmallId = v.ID
		status := c.CheckLiveType(int(result.CourseType), int(v.LiveSmallStatus), v.LiveRoomID)
		re.LiveStatus = status.LiveStatus
		re.SmallType = status.SmallType
		if int(v.Type) == 4 { //如果是课时为资料下载
			re.Type = 4
			if v.Edges.CsAttachment != nil {
				re.FileUrl = app.GetOssHost() + v.Edges.CsAttachment.Filename
			}
			re.SmallId = v.CourseID
		}

		//直播状态
		if v.LiveSmallTime > 0 {
			end := v.LiveSmallStart.Unix() + int64(v.LiveSmallTime)
			ends := time.Unix(end, 0)
			re.LiveAt = v.LiveSmallStart.Format("1月2日 ") + v.LiveSmallStart.Format("15:04") + "-" + ends.Format("15:04")
		}
		re.Name = v.SmallName

		res2 = append(res2, re)
	}

	for _, v := range result.Edges.CourseChapters { //课程下的章
		re := response.ChapterDetail2{
			List: []response.ChapterDetail2{},
		}
		re.Type = 1
		re.Name = v.Title
		if v.Edges.CourseSmallChapters != nil {
			for _, v1 := range v.Edges.CourseSmallChapters {
				r1 := response.ChapterDetail2{
					List: []response.ChapterDetail2{},
				}
				r1.Name = v1.SmallName
				r1.Type = 3
				r1.SmallId = v1.ID
				status := c.CheckLiveType(int(result.CourseType), int(v1.LiveSmallStatus), v1.LiveRoomID)
				r1.LiveStatus = status.LiveStatus
				r1.SmallType = status.SmallType
				if int(v1.Type) == 4 { //如果是课时为资料下载

					if v1.Edges.CsAttachment != nil {
						r1.FileUrl = app.GetOssHost() + v1.Edges.CsAttachment.Filename
					}
					r1.Type = 4
					r1.SmallId = v1.CourseID
				}

				//直播状态
				if v1.LiveSmallTime > 0 {
					end := v1.LiveSmallStart.Unix() + int64(v1.LiveSmallTime)
					ends := time.Unix(end, 0)
					r1.LiveAt = v1.LiveSmallStart.Format("1月2日 ") + v1.LiveSmallStart.Format("15:04") + "-" + ends.Format("15:04")
				}
				re.List = append(re.List, r1)
			}
		}

		if v.Edges.ChapterSections != nil { //章下的节
			for _, v2 := range v.Edges.ChapterSections {
				r2 := response.ChapterDetail2{
					List: []response.ChapterDetail2{},
				}
				r2.Name = v2.Title
				r2.Type = 2

				for _, v3 := range v2.Edges.CourseSmallSections {
					r3 := response.ChapterDetail2{
						List: []response.ChapterDetail2{},
					}
					r2.Count = r2.Count + 1

					r3.Type = 3
					r3.Name = v3.SmallName
					r3.SmallId = v3.ID
					status := c.CheckLiveType(int(result.CourseType), int(v3.LiveSmallStatus), v3.LiveRoomID)
					r3.LiveStatus = status.LiveStatus
					r3.SmallType = status.SmallType
					if int(v3.Type) == 4 { //如果是课时为资料下载
						if v3.Edges.CsAttachment != nil {
							r3.FileUrl = app.GetOssHost() + v3.Edges.CsAttachment.Filename
						}
						r3.Type = 4
						r3.SmallId = v3.CourseID
					}

					//直播状态
					if v3.LiveSmallTime > 0 {
						end := v3.LiveSmallStart.Unix() + int64(v3.LiveSmallTime)
						ends := time.Unix(end, 0)
						re.LiveAt = v3.LiveSmallStart.Format("1月2日 ") + v3.LiveSmallStart.Format("15:04") + "-" + ends.Format("15:04")
					}

					r2.List = append(r2.List, r3)
				}

				re.List = append(re.List, r2)
			}

		}

		res2 = append(res2, re)

	}

	return &res2, nil
}

//获取课时的状态
func (c Course) CheckLiveType(courseType int, smallLive int, roomId int) *CourseLiveType {
	var res CourseLiveType
	switch courseType {
	case 2, 3: //直播小节课程
		//判断直播状态
		var statusArr = []int{
			5, //已上传录像
			7, //已生成回放
		}
		if smallLive == 1 {
			// 1.未开始，2:直播中，3：直播结束，4：回放
			res.LiveStatus = 1
		}

		if smallLive == 3 {
			res.LiveStatus = 3
		}
		res.SmallType = 1

		if app.IsContainInt(statusArr, smallLive) { //是否是回放
			if roomId > 0 {
				// 1.未开始，2:直播中，3：直播结束，4：回放
				res.LiveStatus = 4
				res.SmallType = 2
			}
		}

		//直播中获取学生参加码
		if smallLive == 2 {
			// 1.未开始，2:直播中，3：直播结束，4：回放
			res.LiveStatus = 2
		}
	case 1, 4: //普通课程 和录播
		res.SmallType = 3
	}

	return &res
}

//获取课程资料下载
func (c Course) GetCourseFileList(ctx context.Context, id int) (*[]CourseFile, error) {
	//课时类型：1：视频，2：音频，3：直播，4：资料下载
	s := store.WithContext(ctx)
	list := s.KcCourseSmallCategory.Query().SoftDelete().Where(kccoursesmallcategory.CourseID(id), kccoursesmallcategory.Type(4)).WithCourse(func(query *ent.KcCourseQuery) {
		query.SoftDelete().WithAttachment(func(query *ent.AttachmentQuery) {
			query.SoftDelete()
		})
	}).WithCsAttachment(func(query *ent.AttachmentQuery) {
		query.SoftDelete()
	}).AllX(ctx)
	var fileList []CourseFile
	for _, v := range list {
		if v.Edges.CsAttachment != nil {
			file := CourseFile{}
			file.Id = v.ID
			file.CourseId = v.CourseID
			if v.Edges.Course != nil {
				file.CourseName = v.Edges.Course.CourseName
				if v.Edges.Course.Edges.Attachment != nil && v.Edges.Course.Edges.Attachment.Filename != "" {
					file.ImgCover = app.GetOssHost() + v.Edges.Course.Edges.Attachment.Filename
				}
			}
			file.Title = v.CoursewareName
			if v.Edges.CsAttachment.Filename != "" {
				file.FileName = app.GetOssHost() + v.Edges.CsAttachment.Filename
				file.MimeType = v.Edges.CsAttachment.MimeType
				file.FileSize = int(v.Edges.CsAttachment.FileSize)
			}

			fileList = append(fileList, file)
		}
	}

	return &fileList, nil
}

//获取用户报名的课程题库id
func (c Course) GetUserCourseBankIds(ctx context.Context, uid int) ([]int, error) {
	s := store.WithContext(ctx)
	//班级课程
	infoList := s.KcUserClass.Query().Where(kcuserclass.UserID(uid)).
		WithClass(func(query *ent.KcClassQuery) {
			query.WithKcClassCourses(func(query2 *ent.KcCourseQuery) {
				query2.SoftDelete().WithQuestionBank(func(query *ent.TkQuestionBankQuery) {
					query.SoftDelete().Select("id")
				})
			})
		}).AllX(ctx)

	couseList := s.KcUserCourse.Query().SoftDelete().Where(kcusercourse.UserID(uid)).WithCourse(func(query *ent.KcCourseQuery) {
		query.SoftDelete().WithQuestionBank(func(query *ent.TkQuestionBankQuery) {
			query.SoftDelete().Select("id")
		})
	}).AllX(ctx)

	var quesIds []int

	for _, v := range infoList {
		if len(v.Edges.Class.Edges.KcClassCourses) > 0 {
			for _, v1 := range v.Edges.Class.Edges.KcClassCourses {
				if v1.Edges.QuestionBank != nil {
					quesIds = append(quesIds, v1.Edges.QuestionBank.ID)

				}
			}
		}
	}

	for _, v2 := range couseList {
		if v2.Edges.Course != nil {
			if v2.Edges.Course.Edges.QuestionBank != nil {
				quesIds = append(quesIds, v2.Edges.Course.Edges.QuestionBank.ID)
			}
		}
	}

	return quesIds, nil
}
