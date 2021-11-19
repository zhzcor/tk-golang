package app

import (
	"context"
	"time"
	"tkserver/httpapi/admin/request"
	"tkserver/internal/errorno"
	"tkserver/internal/store"
	"tkserver/internal/store/ent"
	"tkserver/internal/store/ent/useraskanswer"
	"tkserver/internal/store/ent/useraskanswerattachment"
	"tkserver/internal/store/ent/usercourseappraise"
	app "tkserver/pkg/requestparam"
)

type UserManager struct{}

//删除问答
func (u UserManager) DelUserAskAnswer(ctx context.Context, id int) error {
	s := store.WithContext(ctx)
	if errorno := u.UserAskAnswerIdExist(ctx, id); errorno != nil {
		return errorno
	}
	_, err := s.UserAskAnswer.UpdateOneID(id).SoftDelete().Save(ctx)
	if err != nil {
		return errorno.NewStoreErr(err)
	}
	return nil
}

//删除问答回复
func (u UserManager) DelUserAskAnswerReply(ctx context.Context, id int) error {
	s := store.WithContext(ctx)
	if errorno := u.UserAskAnswerIdExist(ctx, id); errorno != nil {
		return errorno
	}
	_, err := s.UserAskAnswer.UpdateOneID(id).
		SetAnswerDesc("").
		ClearAnswerAt().
		SetAnswerStatus(2).
		Save(ctx)
	if err != nil {
		return errorno.NewStoreErr(err)
	}
	_, err = s.UserAskAnswerAttachment.Delete().
		Where(useraskanswerattachment.Type(2)).
		Where(useraskanswerattachment.AskID(id)).
		Exec(ctx)
	if err != nil {
		return errorno.NewStoreErr(err)
	}
	return nil
}

//设置问答显示状态
func (u UserManager) SetUserAskAnswerShowStatus(ctx context.Context, id, status int) error {
	s := store.WithContext(ctx)
	if errorno := u.UserAskAnswerIdExist(ctx, id); errorno != nil {
		return errorno
	}
	_, err := s.UserAskAnswer.UpdateOneID(id).SetShowStatus(uint8(status)).Save(ctx)
	if err != nil {
		return errorno.NewStoreErr(err)
	}
	return nil
}

//设置问答回复显示状态
func (u UserManager) SetUserAskAnswerReplyShowStatus(ctx context.Context, id, status int) error {
	s := store.WithContext(ctx)
	if errorno := u.UserAskAnswerIdExist(ctx, id); errorno != nil {
		return errorno
	}
	_, err := s.UserAskAnswer.UpdateOneID(id).SetReplyShowStatus(uint8(status)).Save(ctx)
	if err != nil {
		return errorno.NewStoreErr(err)
	}
	return nil
}

//问答回复
func (u UserManager) ReplyUserAskAnswer(ctx context.Context, req request.ReplyUserAskAnswer) error {
	s := store.WithContext(ctx)
	if errorno := u.UserAskAnswerIdExist(ctx, *req.Id); errorno != nil {
		return errorno
	}
	res, err := s.UserAskAnswer.UpdateOneID(*req.Id).
		SetAnswerStatus(1).
		SetAnswerDesc(*req.AnswerDesc).
		SetAnswerAt(time.Now()).
		Save(ctx)
	if err != nil {
		return errorno.NewStoreErr(err)
	}
	if !app.IsNil(req.AttachmentIds) && len(req.AttachmentIds) > 0 {
		replyBulk := make([]*ent.UserAskAnswerAttachmentCreate, len(req.AttachmentIds))
		for i, v := range req.AttachmentIds {
			replyBulk[i] = s.UserAskAnswerAttachment.Create().
				SetAttachmentID(v).SetAskID(res.ID).SetType(2)
		}
		_, err := s.UserAskAnswerAttachment.CreateBulk(replyBulk...).Save(ctx)
		if err != nil {
			return err
		}
	}
	return nil
}

//id判断问答是否存着
func (u UserManager) UserAskAnswerIdExist(ctx context.Context, id int) error {
	s := store.WithContext(ctx)
	fined, err := s.UserAskAnswer.
		Query().SoftDelete().
		Where(useraskanswer.ID(id)).
		Exist(ctx)
	if err != nil {
		return errorno.NewStoreErr(err)
	}
	if !fined {
		return errorno.NewErr(errorno.Errno{
			Code:    20001,
			Message: "用户问答不存在",
		})
	}
	return nil
}

//删除课程评论
func (u UserManager) DelUserCourseAppraise(ctx context.Context, id int) error {
	s := store.WithContext(ctx)
	if errorno := u.UserCourseAppraiseIdExist(ctx, id); errorno != nil {
		return errorno
	}
	_, err := s.UserCourseAppraise.UpdateOneID(id).SoftDelete().Save(ctx)
	if err != nil {
		return errorno.NewStoreErr(err)
	}
	return nil
}

//设置课程评论显示状态
func (u UserManager) SetUserCourseAppraiseStatus(ctx context.Context, id, status int) error {
	s := store.WithContext(ctx)
	if errorno := u.UserCourseAppraiseIdExist(ctx, id); errorno != nil {
		return errorno
	}
	_, err := s.UserCourseAppraise.UpdateOneID(id).SetShowStatus(uint8(status)).Save(ctx)
	if err != nil {
		return errorno.NewStoreErr(err)
	}
	return nil
}

//id判断课程评论是否存着
func (u UserManager) UserCourseAppraiseIdExist(ctx context.Context, id int) error {
	s := store.WithContext(ctx)
	fined, err := s.UserCourseAppraise.
		Query().SoftDelete().
		Where(usercourseappraise.ID(id)).
		Exist(ctx)
	if err != nil {
		return errorno.NewStoreErr(err)
	}
	if !fined {
		return errorno.NewErr(errorno.Errno{
			Code:    20001,
			Message: "课程评论不存在",
		})
	}
	return nil
}

//课程评论回复
func (u UserManager) ReplyUserCourseAppraise(ctx context.Context, req request.ReplyUserCourseAppraise) error {
	s := store.WithContext(ctx)
	if errorno := u.UserCourseAppraiseIdExist(ctx, *req.Id); errorno != nil {
		return errorno
	}
	_, err := s.UserCourseAppraise.UpdateOneID(*req.Id).
		SetTeacherReply(*req.TeacherReply).
		Save(ctx)
	if err != nil {
		return errorno.NewStoreErr(err)
	}
	return nil
}
