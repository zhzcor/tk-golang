package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		CommonMixin{},
		PasswordMixin{},
	}
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Int("boss_user_id").Default(0),
		field.String("email").Default(""),
		field.String("phone").Default(""),
		field.String("nickname").Default(""),
		field.String("username").Default(""),
		field.Uint8("status").Default(1).Comment("学员状态：预留"),
		field.Uint8("sex").Default(3).Comment("性别，1：男，2：女，3：未知"),
		field.Uint8("reg_from").Default(1).Comment("用户注册的来源，1：系统同步，2:系统注册，3：人工添加，4：用户注册"),
		field.Uint8("card_type").Default(1).Comment("证件类型，1：无，2：身份证，3：居住证，4：护照，5：港澳台身份证"),
		field.String("id_card").Default(""),
		field.Int("from_city_id").Optional(),
		field.Int("from_item_category_id").Optional(),
		field.Time("birthday").Nillable().Optional(),
		field.Text("sign_remark").Optional(),
		field.String("avatar").Default("").Comment("用户头像"),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("login_log", UserLoginLog.Type),
		edge.To("messages", Message.Type),
		edge.To("user_courses", KcUserCourse.Type), //用户课程
		edge.To("user_classes", KcUserClass.Type),  //用户班级
		edge.To("user_exams_records", TkUserExamScoreRecord.Type),
		edge.To("user_question_bank_records", TkUserQuestionBankRecord.Type),
		edge.To("user_question_records", TkUserQuestionRecord.Type), //用户做题记录
		edge.To("ask_users", UserAskAnswer.Type),                    //用户老师问答
		edge.To("course_appraise_users", UserCourseAppraise.Type),   //用户 课程 评价
		edge.To("user_video_record", VideoRecord.Type),              //用户观看视频记录
		edge.From("city", City.Type).Ref("user_city").Field("from_city_id").Unique(),
		edge.From("cate", ItemCategory.Type).Ref("user_item_cate").Field("from_item_category_id").Unique(),
	}
}
