package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

//题目答案选项
type TkQuestionAnswerOption struct {
	ent.Schema
}

func (TkQuestionAnswerOption) Mixin() []ent.Mixin {
	return []ent.Mixin{
		CommonMixin{},
	}
}

func (TkQuestionAnswerOption) Fields() []ent.Field {
	return []ent.Field{
		field.String("option_name").Default("").Comment("选项名称：A,B,C,D等"),
		field.Text("content").Default("").Comment("选项内容"),
		field.Uint8("is_right").Default(1).Comment("是否正确 1：否，2：是"),
		field.Int("question_id").Optional().Comment("题目id"),
	}
}

func (TkQuestionAnswerOption) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("question", TkQuestion.Type).Ref("answer_options").Field("question_id").Unique(),
	}
}
