package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

//题目-节中间表
type TkQuestionSection struct {
	ent.Schema
}

func (TkQuestionSection) Mixin() []ent.Mixin {
	return []ent.Mixin{
		CommonMixin{},
	}
}

func (TkQuestionSection) Fields() []ent.Field {
	return []ent.Field{
		field.Int("section_id").Optional().Comment("节id"),
		field.Int("question_id").Optional().Comment("题目id"),
	}
}

func (TkQuestionSection) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("question_section", TkSection.Type).Ref("tk_section_links").Field("section_id").Unique(),
		edge.From("section_question", TkQuestion.Type).Ref("question_section_links").Field("question_id").Unique(),
	}
}
