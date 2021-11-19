package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

//题库专业中间表
type TkQuestionBankMajor struct {
	ent.Schema
}

func (TkQuestionBankMajor) Mixin() []ent.Mixin {
	return []ent.Mixin{
		CommonMixin{},
	}
}

func (TkQuestionBankMajor) Fields() []ent.Field {
	return []ent.Field{
		field.Int("major_id").Optional().Comment("专业id"),
		field.Int("question_bank_id").Optional().Comment("题库id"),
	}
}

func (TkQuestionBankMajor) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("tk_question_bank", TkQuestionBank.Type).Ref("major_question_banks").Field("question_bank_id").Unique(),
		edge.From("major", Major.Type).Ref("question_bank_majors").Field("major_id").Unique(),
	}
}
