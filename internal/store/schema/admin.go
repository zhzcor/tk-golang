package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

//管理员
type Admin struct {
	ent.Schema
}

func (Admin) Mixin() []ent.Mixin {
	return []ent.Mixin{
		CommonMixin{},
		PasswordMixin{},
	}
}

// Fields of the User.
func (Admin) Fields() []ent.Field {
	return []ent.Field{
		field.Int("boss_admin_id").Default(0).Comment("工作台管理员id"),
		field.String("real_name").Default("").Comment("管理员姓名"),
		field.String("email").Default("").Comment("绑定的邮箱地址"),
		field.String("phone").Default("").Comment("手机号"),
		//field.String("oauth_code").Default("").Comment("鉴权时用来登录的凭证"),
		field.String("third_openid").Default("").Comment("第三方id"),
		field.Uint8("platform").Default(0).Comment("首次注册的平台，0.后台，1.钉钉，2.企业微信"),
		field.Uint8("is_active").Default(0).Comment("管理员是否在职  0：在职，1：离职"),
		field.Uint8("status").Default(1).Comment("状态  1：启用，2：未启用"),
		field.Int("admin_avatar_id").Optional().Comment("头像id"),
		field.String("remark").Default("").Comment("备注"),
		//field.Uint8("first_third").Optional().Comment("首次注册的平台，0.后台，1.钉钉，2.企业微信"),
	}
}

// Edges of the User.
func (Admin) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("admin_login_logs", AdminLoginLog.Type),
		edge.To("operation_admin_logs", AdminOperationLog.Type),
		edge.To("roles", Role.Type),
		edge.To("activities", Activity.Type),
		edge.To("message_created_admin_id", Message.Type),                                                     //消息发布人
		edge.To("admin_question_banks", TkQuestionBank.Type),                                                  //题库创建人
		edge.To("admin_Questions", TkQuestion.Type),                                                           //题目创建人
		edge.To("admin_exam_papers", TkExamPaper.Type),                                                        //试卷创建人
		edge.To("admin_error_feedbacks", TkQuestionErrorFeedback.Type),                                        //错题反馈处理人
		edge.To("course", KcCourse.Type),                                                                      //课程创建人
		edge.To("class_admin", KcClass.Type),                                                                  //班级创建人
		edge.From("admin_attachments", Attachment.Type).Ref("admin_img_id").Field("admin_avatar_id").Unique(), //管理员头像
	}
}
