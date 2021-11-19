package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

//题目纠错
type TkQuestionErrorFeedback struct {
	ent.Schema
}

func (TkQuestionErrorFeedback) Mixin() []ent.Mixin {
	return []ent.Mixin{
		CommonMixin{},
	}
}

func (TkQuestionErrorFeedback) Fields() []ent.Field {
	return []ent.Field{
		field.String("username").Default("").Comment("用户名"),
		field.String("phone").Default("").Comment("手机号"),
		field.String("error_desc").Default("").Comment("错误描述"),
		field.Uint8("error_type").Default(5).Comment("错误类型，1：题干错误，2：选项错误，3：答案错误，4：解析错误，5：其他"),
		field.Uint8("status").Default(2).Comment("处理状态，1：已处理，2：未处理，3：没问题"),
		field.String("deal_remark").Default("").Comment("处理信息备注"),
		field.Int("operator_admin_id").Optional().Comment("处理人id"),
		field.Int("question_id").Optional().Comment("题目id"),
	}
}

func (TkQuestionErrorFeedback) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("question", TkQuestion.Type).Ref("question_error_feedbacks").Field("question_id").Unique(),
		edge.From("admin", Admin.Type).Ref("admin_error_feedbacks").Field("operator_admin_id").Unique(),
	}
}
