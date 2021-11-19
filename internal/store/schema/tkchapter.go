package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

//题库章
type TkChapter struct {
	ent.Schema
}

func (TkChapter) Mixin() []ent.Mixin {
	return []ent.Mixin{
		CommonMixin{},
	}
}

func (TkChapter) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Default("").Comment("章名称"),
		field.Int("question_bank_id").Optional().Comment("题库id"),
		field.Int("question_count").Default(0).Comment("章题目数量"),
	}
}

func (TkChapter) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("question_bank", TkQuestionBank.Type).Ref("question_chapters").Field("question_bank_id").Unique(),
		edge.To("sections", TkSection.Type),
	}
}
