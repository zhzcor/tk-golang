package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

//题库地区中间表
type TkQuestionBankCity struct {
	ent.Schema
}

func (TkQuestionBankCity) Mixin() []ent.Mixin {
	return []ent.Mixin{
		CommonMixin{},
	}
}

func (TkQuestionBankCity) Fields() []ent.Field {
	return []ent.Field{
		field.Int("city_id").Optional().Comment("地区id"),
		field.Int("question_bank_id").Optional().Comment("题库id"),
	}
}

func (TkQuestionBankCity) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("tk_question_bank", TkQuestionBank.Type).Ref("city_question_banks").Field("question_bank_id").Unique(),
		edge.From("city", City.Type).Ref("question_bank_cities").Field("city_id").Unique(),
	}
}
