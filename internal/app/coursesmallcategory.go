package app

import (
	"context"
	"database/sql"
	"gserver/httpapi/admin/request"
	"gserver/internal/errorno"
	"gserver/internal/store"
	"gserver/internal/store/ent"
	"gserver/internal/store/ent/attachment"
	"gserver/internal/store/ent/kccourse"
	"gserver/internal/store/ent/kccoursesmallcategory"
	"gserver/internal/store/ent/kcsmallcategoryattachment"
	"gserver/internal/store/ent/kcsmallcategoryexampaper"
	"gserver/internal/store/ent/kcsmallcategoryquestion"
	"gserver/internal/store/ent/kcvideouploadtask"
	"gserver/pkg/asynctask"
	"gserver/pkg/baijiayun"
	app "gserver/pkg/requestparam"
	"strconv"
	"time"
)

//编辑课时
func (c Course) UpdateKcCourseSmallCategory(ctx context.Context, req request.SetCourseSmallCategory) (int, error) {
	s := store.WithContext(ctx)
	smallCategory, err := s.KcCourseSmallCategory.Query().SoftDelete().
		Where(kccoursesmallcategory.ID(*req.Id)).First(ctx)
	if err != nil && !ent.IsNotFound(err) {
		return 0, errorno.NewStoreErr(err)
	}
	if ent.IsNotFound(err) {
		return 0, errorno.NewErr(errorno.DataNotExistsError)
	}

	creation := s.KcCourseSmallCategory.UpdateOneID(*req.Id).
		SetFinishType(uint8(req.FinishType)).
		SetSmallName(*req.SmallName).
		SetType(uint8(*req.Type))

	if !app.IsNil(req.ViewingTime) {
		creation = creation.SetViewingTime(*req.ViewingTime)
	}
	if *req.Type == 4 {
		creation = creation.SetCoursewareAttachID(*req.CoursewareAttachId)
		if !app.IsNil(req.CoursewareName) {
			creation = creation.SetCoursewareName(*req.CoursewareName)
		}
	} else {
		bjy := baijiayun.BjyCloud{}
		switch *req.CourseType {
		case 1, 4: //普通课程,录播公开课 上传点播视频
			if smallCategory.OrderVideoAttachID != *req.OrderVideoAttachId {
				viewType := 0
				if *req.Type == 2 {
					viewType = 1
				}
				uploadUrl, err := bjy.GetOrderUploadUrl(*req.SmallName, viewType)
				if err != nil {
					return 0, err
				}
				creation = creation.
					SetOrderVideoAttachID(*req.OrderVideoAttachId).
					SetOrderVideoID(uploadUrl.VideoId)

				attachmentUrl, err := c.GetAttachmentDetail(ctx, *req.OrderVideoAttachId)
				if err != nil {
					return 0, err
				}
				//创建点播上传任务
				taskId, err := bjy.CreateUploadTask(ctx, uploadUrl.VideoId, *req.OrderVideoAttachId, *req.CourseId, 1, *req.SmallName, req.OrderVideoName)
				if err != nil {
					return 0, err
				}
				err = asynctask.Enqueue("order_video_upload", map[string]interface{}{
					"id":         taskId,
					"video_url":  attachmentUrl,
					"upload_url": uploadUrl.UploadUrl,
					"filename":   *req.SmallName,
				})
				if err != nil {
					return 0, errorno.NewParamErr(err)
				}
				//bjy.OrderUpload(ctx, taskId, attachmentUrl, uploadUrl.UploadUrl, *req.SmallName)
			}
			break
		case 2: //直播 创建直播间
			if app.IsNil(req.LiveSmallStart) || app.IsNil(req.LiveSmallTime) {
				return 0, errorno.NewErr(errorno.Errno{
					Code:    20001,
					Message: "直播课开始时间不能为空",
				})
			}

			mockVideoId := smallCategory.FalseVideoID
			//伪直播
			if !app.IsNil(req.FalseVideoId) {
				mockVideoId = *req.FalseVideoId
				creation = creation.SetFalseVideoID(*req.FalseVideoId)
			}

			e, _ := time.ParseDuration(strconv.Itoa(*req.LiveSmallTime) + "s")
			startAt := int(req.LiveSmallStart.Unix())
			endAt := int(req.LiveSmallStart.Add(e).Unix())
			err := bjy.UpdateLiveRoom(*req.SmallName, smallCategory.LiveRoomID, startAt, endAt, mockVideoId)
			if err != nil {
				return 0, err
			}
			creation = creation.
				SetLiveSmallStart(*req.LiveSmallStart).
				SetLiveSmallTime(*req.LiveSmallTime)
				//SetLiveSmallStatus(1)
			if !app.IsNil(req.LiveSmallRemark) {
				creation = creation.SetLiveSmallRemark(*req.LiveSmallRemark)
			}
		}
	}

	res, err := creation.Save(ctx)
	if err != nil {
		return 0, errorno.NewStoreErr(err)
	}
	return res.ID, nil
}

//添加课时
func (c Course) AddKcCourseSmallCategory(ctx context.Context, req request.SetCourseSmallCategory) (int, error) {
	s := store.WithContext(ctx)
	fined, err := s.KcCourse.Query().SoftDelete().
		Where(kccourse.ID(*req.CourseId)).
		Where(kccourse.CourseType(uint8(*req.CourseType))).
		Exist(ctx)

	if err != nil {
		return 0, errorno.NewStoreErr(err)
	}
	if !fined {
		return 0, errorno.NewErr(errorno.Errno{
			Code:    20001,
			Message: "对应类型课程不存在",
		})
	}

	creation := s.KcCourseSmallCategory.Create()
	if !app.IsNil(req.ChapterId) && app.IsNil(req.SectionId) {
		if errorno := c.KcCourseChapterIdExist(ctx, *req.ChapterId); errorno != nil {
			return 0, errorno
		}
		creation = creation.SetChapterID(*req.ChapterId)
	}
	if !app.IsNil(req.SectionId) {
		if errorno := c.KcCourseSectionIdExist(ctx, *req.SectionId); errorno != nil {
			return 0, errorno
		}
		creation = creation.SetSectionID(*req.SectionId)
	}

	creation = creation.SetFinishType(uint8(req.FinishType)).
		SetCourseID(*req.CourseId).
		SetSmallName(*req.SmallName).
		SetType(uint8(*req.Type))

	if !app.IsNil(req.ViewingTime) {
		creation = creation.SetViewingTime(*req.ViewingTime)
	}

	if *req.Type == 4 {
		creation = creation.SetCoursewareAttachID(*req.CoursewareAttachId)
		if !app.IsNil(req.CoursewareName) {
			creation = creation.SetCoursewareName(*req.CoursewareName)
		}
	} else {
		bjy := baijiayun.BjyCloud{}
		switch *req.CourseType {
		case 1, 4: //普通课程,录播公开课 上传点播视频
			if app.IsNil(req.OrderVideoAttachId) {
				return 0, errorno.NewErr(errorno.Errno{
					Code:    20001,
					Message: "请添加点播附件",
				})
			}
			viewType := 0
			if *req.Type == 2 {
				viewType = 1
			}
			uploadUrl, err := bjy.GetOrderUploadUrl(*req.SmallName, viewType)
			if err != nil {
				return 0, err
			}
			creation = creation.
				SetOrderVideoAttachID(*req.OrderVideoAttachId).
				SetOrderVideoID(uploadUrl.VideoId)

			attachmentUrl, err := c.GetAttachmentDetail(ctx, *req.OrderVideoAttachId)
			if err != nil {
				return 0, err
			}
			//创建点播上传任务
			taskId, err := bjy.CreateUploadTask(ctx, uploadUrl.VideoId, *req.OrderVideoAttachId, *req.CourseId, 1, *req.SmallName, req.OrderVideoName)
			if err != nil {
				return 0, err
			}

			err = asynctask.Enqueue("order_video_upload", map[string]interface{}{
				"id":         taskId,
				"video_url":  attachmentUrl,
				"upload_url": uploadUrl.UploadUrl,
				"filename":   *req.SmallName,
			})
			if err != nil {
				return 0, errorno.NewParamErr(err)
			}
			//bjy.OrderUpload(ctx, taskId, attachmentUrl, uploadUrl.UploadUrl, *req.SmallName)
			break
		case 2: //直播 创建直播间
			if app.IsNil(req.LiveSmallStart) || app.IsNil(req.LiveSmallTime) {
				return 0, errorno.NewErr(errorno.Errno{
					Code:    20001,
					Message: "直播课开始时间不能为空",
				})
			}

			isMock := 0
			mockVideoId := 0
			//伪直播
			if !app.IsNil(req.FalseVideoId) {
				isMock = 1
				mockVideoId = *req.FalseVideoId
				creation = creation.SetFalseVideoID(*req.FalseVideoId)
			}

			e, _ := time.ParseDuration(strconv.Itoa(*req.LiveSmallTime) + "s")
			startAt := int(req.LiveSmallStart.Unix())
			endAt := int(req.LiveSmallStart.Add(e).Unix())
			room, err := bjy.CreateLiveRoom(*req.SmallName, startAt, endAt, 2, isMock, mockVideoId)
			if err != nil {
				return 0, err
			}
			roomId, err := strconv.Atoi(room.RoomId)
			if err != nil {
				return 0, err
			}
			creation = creation.SetLiveRoomID(roomId).
				SetLiveSmallStart(*req.LiveSmallStart).
				SetLiveSmallTime(*req.LiveSmallTime).
				SetLiveSmallStatus(1)
			if !app.IsNil(req.LiveSmallRemark) {
				creation = creation.SetLiveSmallRemark(*req.LiveSmallRemark)
			}
		}
	}

	res, err := creation.Save(ctx)
	if err != nil {
		return 0, errorno.NewStoreErr(err)
	}
	return res.ID, nil
}

//删除课时
func (c Course) DelKcSmallCategory(ctx context.Context, id int) error {
	s := store.WithContext(ctx)
	if errorno := c.KcSmallCategoryIdExist(ctx, id); errorno != nil {
		return errorno
	}
	err := s.KcCourseSmallCategory.UpdateOneID(id).SoftDelete().Exec(ctx)
	if err != nil {
		return errorno.NewStoreErr(err)
	}
	return nil
}

//获取oss附件地址
func (c Course) GetAttachmentDetail(ctx context.Context, id int) (string, error) {
	res, err := store.WithContext(ctx).Attachment.Query().SoftDelete().
		Where(attachment.ID(id)).First(ctx)

	if err != nil && !ent.IsNotFound(err) {
		return "", errorno.NewStoreErr(err)
	}
	if ent.IsNotFound(err) {
		return "", errorno.NewErr(errorno.DataNotExistsError)
	}
	url := app.GetOssHost() + res.Filename
	return url, nil
}

//添加课时试卷
func (c Course) SetSmallCategoryExamPaper(ctx context.Context, req request.SetSmallCateExamPagerRequest) (int, error) {
	s := store.WithContext(ctx)
	if errorno := c.KcSmallCategoryIdExist(ctx, *req.SmallCategoryId); errorno != nil {
		return 0, errorno
	}
	existsPaperIds := []int{}
	examPapers, err := s.KcSmallCategoryExamPaper.Query().
		Where(kcsmallcategoryexampaper.Type(uint8(*req.Type))).
		Where(kcsmallcategoryexampaper.SmallCategoryID(*req.SmallCategoryId)).
		Select("exam_paper_id").All(ctx)
	if err != nil && err != sql.ErrNoRows {
		return 0, err
	}

	if err != sql.ErrNoRows {
		for _, v := range examPapers {
			existsPaperIds = append(existsPaperIds, v.ExamPaperID)
		}
	}
	insertIds := app.SliceDiff(*req.ExamPaperIds, existsPaperIds) //去重

	examBulk := make([]*ent.KcSmallCategoryExamPaperCreate, len(insertIds))
	for i, v := range insertIds {
		examBulk[i] = s.KcSmallCategoryExamPaper.Create().
			SetType(uint8(*req.Type)).SetExamPaperID(v)
	}
	courseExamPapers, err := s.KcSmallCategoryExamPaper.CreateBulk(examBulk...).Save(ctx)
	if err != nil {
		return 0, err
	}

	md := s.KcCourseSmallCategory.
		UpdateOneID(*req.SmallCategoryId).
		AddCourseSmallCategoryExampapers(courseExamPapers...)
	if *req.Type == 1 {
		md = md.AddExamCount(len(insertIds))
	} else {
		md = md.AddHomeworkCount(len(insertIds))
	}
	res, err := md.Save(ctx)
	if err != nil {
		return 0, errorno.NewStoreErr(err)
	}
	return res.ID, nil
}

//添加课时习题
func (c Course) SetSmallCategoryQuestion(ctx context.Context, req request.SetSmallCateQuestionRequest) (int, error) {
	s := store.WithContext(ctx)
	if errorno := c.KcSmallCategoryIdExist(ctx, *req.SmallCategoryId); errorno != nil {
		return 0, errorno
	}
	existsQuestionIds := []int{}
	questions, err := s.KcSmallCategoryQuestion.Query().
		Where(kcsmallcategoryquestion.SmallCategoryID(*req.SmallCategoryId)).
		Select("question_id").All(ctx)
	if err != nil && err != sql.ErrNoRows {
		return 0, err
	}

	if err != sql.ErrNoRows {
		for _, v := range questions {
			existsQuestionIds = append(existsQuestionIds, v.QuestionID)
		}
	}
	insertIds := app.SliceDiff(*req.QuestionIds, existsQuestionIds) //去重

	questionBulk := make([]*ent.KcSmallCategoryQuestionCreate, len(insertIds))
	for i, v := range insertIds {
		questionBulk[i] = s.KcSmallCategoryQuestion.Create().
			SetSmallCategoryID(*req.SmallCategoryId).
			SetQuestionID(v)
	}
	_, err = s.KcSmallCategoryQuestion.CreateBulk(questionBulk...).Save(ctx)
	if err != nil {
		return 0, err
	}
	res, err := s.KcCourseSmallCategory.
		UpdateOneID(*req.SmallCategoryId).
		AddQuestionCount(len(insertIds)).
		Save(ctx)
	if err != nil {
		return 0, errorno.NewStoreErr(err)
	}
	return res.ID, nil
}

//添加课时附件
func (c Course) SetSmallCategoryAttachment(ctx context.Context, req request.SetSmallCateAttachmentRequest) (int, error) {
	s := store.WithContext(ctx)
	if errorno := c.KcSmallCategoryIdExist(ctx, *req.SmallCategoryId); errorno != nil {
		return 0, errorno
	}

	AttachmentBulk := make([]*ent.KcSmallCategoryAttachmentCreate, len(req.Attachments))
	for i, v := range req.Attachments {
		AttachmentBulk[i] = s.KcSmallCategoryAttachment.Create().
			SetSmallCategoryID(*req.SmallCategoryId).
			SetAttachmentID(*v.AttachmentId).
			SetAttachmentName(*v.AttachmentName)
	}
	_, err := s.KcSmallCategoryAttachment.CreateBulk(AttachmentBulk...).Save(ctx)
	if err != nil {
		return 0, err
	}

	res, err := s.KcCourseSmallCategory.
		UpdateOneID(*req.SmallCategoryId).
		AddAttachmentCount(len(req.Attachments)).
		Save(ctx)
	if err != nil {
		return 0, errorno.NewStoreErr(err)
	}
	return res.ID, nil
}

//移除课时试卷
func (c Course) RemoKcSmallCategoryExamPaper(ctx context.Context, req request.RemoveSmallCateExamPagerRequest) (int, error) {
	s := store.WithContext(ctx)
	if errorno := c.KcSmallCategoryIdExist(ctx, *req.SmallCategoryId); errorno != nil {
		return 0, errorno
	}
	_, err := s.KcSmallCategoryExamPaper.Delete().
		Where(kcsmallcategoryexampaper.SmallCategoryID(*req.SmallCategoryId)).
		Where(kcsmallcategoryexampaper.ExamPaperID(*req.ExamPaperId)).
		Where(kcsmallcategoryexampaper.Type(uint8(*req.Type))).
		Exec(ctx)
	if err != nil {
		return 0, err
	}
	md := s.KcCourseSmallCategory.
		UpdateOneID(*req.SmallCategoryId)
	if *req.Type == 1 {
		md = md.AddExamCount(-1)
	} else {
		md = md.AddHomeworkCount(-1)
	}
	res, err := md.Save(ctx)
	if err != nil {
		return 0, errorno.NewStoreErr(err)
	}
	return res.ID, nil
}

//移除课时练习
func (c Course) RemoveSmallCategoryQuestion(ctx context.Context, req request.RemoveSmallCateQuestionRequest) (int, error) {
	s := store.WithContext(ctx)
	if errorno := c.KcSmallCategoryIdExist(ctx, *req.SmallCategoryId); errorno != nil {
		return 0, errorno
	}
	_, err := s.KcSmallCategoryQuestion.Delete().
		Where(kcsmallcategoryquestion.SmallCategoryID(*req.SmallCategoryId)).
		Where(kcsmallcategoryquestion.QuestionID(*req.QuestionId)).Exec(ctx)
	if err != nil {
		return 0, err
	}
	res, err := s.KcCourseSmallCategory.
		UpdateOneID(*req.SmallCategoryId).
		AddQuestionCount(-1).
		Save(ctx)
	if err != nil {
		return 0, errorno.NewStoreErr(err)
	}
	return res.ID, nil
}

//移除课时附件
func (c Course) RemoveSmallCategoryAttachment(ctx context.Context, req request.RemoveSmallCateAttachmentRequest) (int, error) {
	s := store.WithContext(ctx)
	if errorno := c.KcSmallCategoryIdExist(ctx, *req.SmallCategoryId); errorno != nil {
		return 0, errorno
	}
	_, err := s.KcSmallCategoryAttachment.Delete().
		Where(kcsmallcategoryattachment.SmallCategoryID(*req.SmallCategoryId)).
		Where(kcsmallcategoryattachment.AttachmentID(*req.AttachmentId)).Exec(ctx)
	if err != nil {
		return 0, err
	}
	res, err := s.KcCourseSmallCategory.
		UpdateOneID(*req.SmallCategoryId).
		AddAttachmentCount(-1).
		Save(ctx)
	if err != nil {
		return 0, errorno.NewStoreErr(err)
	}
	return res.ID, nil
}

//回放重新录制
func (c Course) UpdateLiveBack(ctx context.Context, id int) error {
	s := store.WithContext(ctx)
	if errorno := c.KcSmallCategoryIdExist(ctx, id); errorno != nil {
		return errorno
	}

	smallCate, err := s.KcCourseSmallCategory.Query().SoftDelete().
		Where(kccoursesmallcategory.LiveSmallStatusIn(3, 4, 5, 6)).
		Where(kccoursesmallcategory.ID(id)).First(ctx)
	if err != nil && !ent.IsNotFound(err) {
		return errorno.NewStoreErr(err)
	}
	if ent.IsNotFound(err) {
		return errorno.NewErr(errorno.DataNotExistsError)
	}
	bjy := baijiayun.BjyCloud{}
	liveBackInfo, err := bjy.GetLiveBackInfo(smallCate.LiveRoomID)
	if err != nil {
		return err
	}
	backStatus := 4
	switch liveBackInfo.Status {
	case 30:
		backStatus = 6
		break
	case 100:
		backStatus = 5
		break
	}
	_, err = s.KcCourseSmallCategory.UpdateOneID(id).
		SetBackVideoID(liveBackInfo.VideoId).
		SetLiveSmallStatus(int8(backStatus)).
		Save(ctx)
	if err != nil {
		return errorno.NewStoreErr(err)
	}
	return nil
}

//点播替换回放
func (c Course) ReplaceLiveBack(ctx context.Context, id, attachmentId int, videoName string) error {
	s := store.WithContext(ctx)
	if errorno := c.KcSmallCategoryIdExist(ctx, id); errorno != nil {
		return errorno
	}

	smallCate, err := s.KcCourseSmallCategory.Query().SoftDelete().
		Where(kccoursesmallcategory.LiveSmallStatusIn(3, 4, 5, 6)).
		Where(kccoursesmallcategory.ID(id)).First(ctx)
	if err != nil && !ent.IsNotFound(err) {
		return errorno.NewStoreErr(err)
	}
	if ent.IsNotFound(err) {
		return errorno.NewErr(errorno.DataNotExistsError)
	}

	bjy := baijiayun.BjyCloud{}
	attachmentUrl, err := c.GetAttachmentDetail(ctx, attachmentId)
	if err != nil {
		return err
	}

	uploadUrl, err := bjy.GetOrderUploadUrl(smallCate.SmallName, 0)
	if err != nil {
		return err
	}

	//创建点播上传任务
	taskId, err := bjy.CreateUploadTask(ctx, uploadUrl.VideoId, attachmentId, smallCate.CourseID, 3, smallCate.SmallName, videoName)
	if err != nil {
		return err
	}
	//bjy.OrderUpload(ctx, taskId, attachmentUrl, uploadUrl.UploadUrl, smallCate.SmallName)

	err = asynctask.Enqueue("order_video_upload", map[string]interface{}{
		"id":         taskId,
		"video_url":  attachmentUrl,
		"upload_url": uploadUrl.UploadUrl,
		"filename":   smallCate.SmallName,
	})
	if err != nil {
		return errorno.NewParamErr(err)
	}

	err = bjy.ReplaceLiveBack(smallCate.LiveRoomID, uploadUrl.VideoId)
	if err != nil {
		return err
	}
	_, err = s.KcCourseSmallCategory.UpdateOneID(id).
		SetBackVideoID(uploadUrl.VideoId).
		SetLiveSmallStatus(7).
		Save(ctx)
	if err != nil {
		return errorno.NewStoreErr(err)
	}
	return nil
}

//回放回调处理
func (c Course) ProcessCallbackBack(ctx context.Context, req request.VideoCallbackRequest) error {
	s := store.WithContext(ctx)
	smallCate, err := s.KcCourseSmallCategory.Query().
		Where(kccoursesmallcategory.LiveRoomID(*req.RoomId)).
		First(ctx)
	if err != nil {
		return err
	}
	status := 3
	switch req.Status {
	case 20:
		status = 7
		break
	case 30:
		status = 6
		break
	case 100:
		status = 5
		break
	}
	_, err = s.KcCourseSmallCategory.UpdateOneID(smallCate.ID).
		SetLiveSmallStatus(int8(status)).
		SetBackVideoID(req.VideoId).
		Save(ctx)
	if err != nil {
		return err
	}
	return nil
}

//点播回调处理
func (c Course) ProcessCallbackOrder(ctx context.Context, req request.VideoCallbackRequest) error {
	s := store.WithContext(ctx)
	task, err := s.KcVideoUploadTask.Query().
		Where(kcvideouploadtask.VideoID(req.VideoId)).
		First(ctx)
	if err != nil {
		return err
	}
	status := 3
	switch req.Status {
	case 20:
		status = 3
		break
	case 30:
		status = 6
		break
	case 100:
		status = 5
		break
	}
	_, err = s.KcVideoUploadTask.UpdateOneID(task.ID).
		SetStatus(status).
		SetTotalSize(req.TotalSize).
		SetLength(req.Length).
		Save(ctx)
	if err != nil {
		return err
	}
	return nil
}
