package admin

import (
	"context"
	"database/sql"
	"github.com/gin-gonic/gin"
	"gserver/httpapi/admin/request"
	"gserver/httpapi/admin/response"
	"gserver/internal/app"
	"gserver/internal/store"
	"gserver/internal/store/ent"
	"gserver/internal/store/ent/user"
	"gserver/internal/store/ent/useraskanswer"
	"gserver/internal/store/ent/usercourseappraise"
	app2 "gserver/pkg/requestparam"
)

//获取学生问答列表
func GetAskAnswerPageList(ctx *gin.Context) (interface{}, error) {
	var req request.UserAskAnswerPageList
	res := response.AskAnswerPageListResponse{
		List: []response.AskAnswerDetail{},
	}
	err := ctx.Bind(&req)
	if err != nil {
		return nil, err
	}
	err = store.WithTx(ctx, func(ctx context.Context) error {
		s := store.WithContext(ctx)

		query := s.UserAskAnswer.Query().SoftDelete()
		if !app2.IsNil(req.TeacherId) {
			query = query.Where(useraskanswer.TeacherID(*req.TeacherId))
		}
		if !app2.IsNil(req.AnswerStatus) {
			query = query.Where(useraskanswer.AnswerStatus(uint8(*req.AnswerStatus)))
		}

		count, err := query.Clone().Count(ctx)
		if err != nil {
			return err
		}
		res.Page.Total = count

		resp, err := query.WithAskAnswersAttachments(
			func(query *ent.UserAskAnswerAttachmentQuery) {
				query.WithAttachment(
					func(query *ent.AttachmentQuery) {
						query.Select("id", "filename")
					})
			}).
			WithUser(func(query *ent.UserQuery) {
				query.Select("id", "username", "avatar")
			}).
			WithTeacher(func(query *ent.TeacherQuery) {
				query.Select("id", "name")
			}).ForPage(req.Page, req.PageSize).Order(ent.Desc("id")).All(ctx)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil
			}
			return err
		}

		for _, v := range resp {
			resDetail := response.AskAnswerDetail{}
			resDetail.ID = v.ID
			resDetail.CreatedAt = v.CreatedAt
			resDetail.ShowStatus = v.ShowStatus
			resDetail.AskDesc = v.AskDesc
			resDetail.UserID = v.UserID
			resDetail.TeacherID = v.TeacherID
			resDetail.AnswerDesc = v.AnswerDesc
			resDetail.AnswerStatus = v.AnswerStatus
			resDetail.AnswerAt = v.AnswerAt
			resDetail.ReplyShowStatus = v.ReplyShowStatus
			resDetail.Username = ""
			resDetail.UserAvatarUrl = ""
			resDetail.TeacherName = ""
			resDetail.AskAttachments = []response.AttachmentDetail{}
			resDetail.AnswerAttachments = []response.AttachmentDetail{}
			if !app2.IsNil(v.Edges.User) {
				if v.Edges.User.Avatar != "" {
					resDetail.UserAvatarUrl = app2.GetOssHost() + v.Edges.User.Avatar
				}
				resDetail.Username = v.Edges.User.Username
			}

			if !app2.IsNil(v.Edges.Teacher) {
				resDetail.TeacherName = v.Edges.Teacher.Name
			}

			if !app2.IsNil(v.Edges.AskAnswersAttachments) && len(v.Edges.AskAnswersAttachments) > 0 {
				for _, a := range v.Edges.AskAnswersAttachments {
					if !app2.IsNil(a.Edges.Attachment) {
						attachmentDetail := response.AttachmentDetail{}
						attachmentDetail.AttachmentId = a.Edges.Attachment.ID
						attachmentDetail.AttachmentUrl = app2.GetOssHost() + a.Edges.Attachment.Filename
						if a.Type == 1 {
							resDetail.AskAttachments = append(resDetail.AskAttachments, attachmentDetail)
						} else {
							resDetail.AnswerAttachments = append(resDetail.AnswerAttachments, attachmentDetail)
						}
					}
				}
			}
			res.List = append(res.List, resDetail)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}

//删除问答
func DelAskAnswer(ctx *gin.Context) (interface{}, error) {
	var req request.IdOnly
	err := ctx.Bind(&req)
	if err != nil {
		return nil, err
	}
	uc := app.UserManager{}
	err = store.WithTx(ctx, func(ctx context.Context) error {
		err = uc.DelUserAskAnswer(ctx, req.Id)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return nil, nil
}

//删除问答回复
func DelAskAnswerReply(ctx *gin.Context) (interface{}, error) {
	var req request.IdOnly
	err := ctx.Bind(&req)
	if err != nil {
		return nil, err
	}
	uc := app.UserManager{}
	err = store.WithTx(ctx, func(ctx context.Context) error {
		err = uc.DelUserAskAnswerReply(ctx, req.Id)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return nil, nil
}

//设置问答状态
func SetAskAnswerShowStatus(ctx *gin.Context) (interface{}, error) {
	var req request.BasicConfigStatus
	err := ctx.Bind(&req)
	if err != nil {
		return nil, err
	}
	uc := app.UserManager{}
	err = store.WithTx(ctx, func(ctx context.Context) error {
		err = uc.SetUserAskAnswerShowStatus(ctx, req.Id, req.Status)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return nil, nil
}

//设置问答回复状态
func SetAskAnswerReplyShowStatus(ctx *gin.Context) (interface{}, error) {
	var req request.BasicConfigStatus
	err := ctx.Bind(&req)
	if err != nil {
		return nil, err
	}
	uc := app.UserManager{}
	err = store.WithTx(ctx, func(ctx context.Context) error {
		err = uc.SetUserAskAnswerReplyShowStatus(ctx, req.Id, req.Status)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return nil, nil
}

//问答回复
func ReplyAskAnswer(ctx *gin.Context) (interface{}, error) {
	var req request.ReplyUserAskAnswer
	err := ctx.Bind(&req)
	if err != nil {
		return nil, err
	}
	uc := app.UserManager{}
	err = store.WithTx(ctx, func(ctx context.Context) error {
		err = uc.ReplyUserAskAnswer(ctx, req)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return nil, nil
}

//获取课程评论列表
func GetUserCourseAppraisePageList(ctx *gin.Context) (interface{}, error) {
	var req request.UserCourseAppraisePageList
	res := response.UserCourseAppraisePageListResponse{
		List: []response.UserCourseAppraiseDetail{},
	}
	err := ctx.Bind(&req)
	if err != nil {
		return nil, err
	}
	err = store.WithTx(ctx, func(ctx context.Context) error {
		s := store.WithContext(ctx)

		query := s.UserCourseAppraise.Query().SoftDelete()
		if !app2.IsNil(req.Username) {
			query = query.Where(usercourseappraise.HasUserWith(user.UsernameContains(*req.Username)))
		}
		if !app2.IsNil(req.StartAt) && !app2.IsNil(req.EndAt) {
			query = query.Where(usercourseappraise.CreatedAtGT(*req.StartAt),
				usercourseappraise.CreatedAtLT(*req.EndAt))
		}

		count, err := query.Clone().Count(ctx)
		if err != nil {
			return err
		}
		res.Page.Total = count

		resp, err := query.WithUser(func(query *ent.UserQuery) {
			query.Select("id", "username", "phone", "avatar")
		}).WithSmallCate(func(query *ent.KcCourseSmallCategoryQuery) {
			query.Select("id", "small_name", "course_id").WithCourse(
				func(query *ent.KcCourseQuery) {
					query.Select("id", "course_name")
				})
		}).ForPage(req.Page, req.Page).Order(ent.Desc("id")).All(ctx)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil
			}
			return err
		}

		for _, v := range resp {
			resDetail := response.UserCourseAppraiseDetail{}
			resDetail.ID = v.ID
			resDetail.CreatedAt = v.CreatedAt
			resDetail.ShowStatus = v.ShowStatus
			resDetail.Type = v.Type
			resDetail.UserID = v.UserID
			resDetail.SmallCateID = v.SmallCateID
			resDetail.CreatedAt = v.CreatedAt
			resDetail.CompositeScore = v.CompositeScore
			resDetail.TeachAtmosphereScore = v.TeachAtmosphereScore
			resDetail.TeachAttitudeScore = v.TeachAttitudeScore
			resDetail.TeachContentScore = v.TeachContentScore
			resDetail.TeacherImpression = v.TeacherImpression
			resDetail.Desc = v.Desc
			resDetail.TeacherReply = v.TeacherReply
			resDetail.UserAvatarUrl = ""
			resDetail.Username = ""
			resDetail.SmallName = ""
			resDetail.CourseName = ""
			if !app2.IsNil(v.Edges.User) {
				if v.Edges.User.Avatar != "" {
					resDetail.UserAvatarUrl = app2.GetOssHost() + v.Edges.User.Avatar
				}
				resDetail.Username = v.Edges.User.Username
				resDetail.Phone = v.Edges.User.Phone
			}

			if !app2.IsNil(v.Edges.SmallCate) {
				resDetail.SmallName = v.Edges.SmallCate.SmallName
				if !app2.IsNil(v.Edges.SmallCate.Edges.Course) {
					resDetail.CourseName = v.Edges.SmallCate.Edges.Course.CourseName
				}
			}
			res.List = append(res.List, resDetail)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}

//删除问答回复
func DelUserCourseAppraise(ctx *gin.Context) (interface{}, error) {
	var req request.IdOnly
	err := ctx.Bind(&req)
	if err != nil {
		return nil, err
	}
	uc := app.UserManager{}
	err = store.WithTx(ctx, func(ctx context.Context) error {
		err = uc.DelUserCourseAppraise(ctx, req.Id)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return nil, nil
}

//设置问答状态
func SetUserCourseAppraiseStatus(ctx *gin.Context) (interface{}, error) {
	var req request.BasicConfigStatus
	err := ctx.Bind(&req)
	if err != nil {
		return nil, err
	}
	uc := app.UserManager{}
	err = store.WithTx(ctx, func(ctx context.Context) error {
		err = uc.SetUserCourseAppraiseStatus(ctx, req.Id, req.Status)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return nil, nil
}

//课程评论回复
func ReplyUserCourseAppraise(ctx *gin.Context) (interface{}, error) {
	var req request.ReplyUserCourseAppraise
	err := ctx.Bind(&req)
	if err != nil {
		return nil, err
	}
	uc := app.UserManager{}
	err = store.WithTx(ctx, func(ctx context.Context) error {
		err = uc.ReplyUserCourseAppraise(ctx, req)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return nil, nil
}
