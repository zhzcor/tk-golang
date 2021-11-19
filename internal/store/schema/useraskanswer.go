package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

//用户 老师问答
type UserAskAnswer struct {
	ent.Schema
}

func (UserAskAnswer) Mixin() []ent.Mixin {
	return []ent.Mixin{
		CommonMixin{},
	}
}

// Fields of the KcCourse.
func (UserAskAnswer) Fields() []ent.Field {
	return []ent.Field{
		field.String("ask_desc").Default("").Comment("问题描述"),
		field.Uint8("answer_status").Default(2).Comment("1:已解答，2:未解答"),
		field.Int("user_id").Optional().Comment("用户id"),
		field.Int("teacher_id").Optional().Comment("老师id"),
		field.Uint8("show_status").Default(1).Comment("问答显示状态：1:显示，2:隐藏"),
		field.Uint8("reply_show_status").Default(1).Comment("回复显示状态：1:显示，2:隐藏"),
		field.String("answer_desc").Default("").Comment("问题回复"),
		field.Time("answer_at").Optional().Comment("回复时间"),
	}
}

// Edges of the Course.
func (UserAskAnswer) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("teacher", Teacher.Type).Ref("ask_teachers").Field("teacher_id").Unique(), //专业 (多)
		edge.From("user", User.Type).Ref("ask_users").Field("user_id").Unique(),
		edge.To("ask_answers_attachments", UserAskAnswerAttachment.Type),
	}
}
