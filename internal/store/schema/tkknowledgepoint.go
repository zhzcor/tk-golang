package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

//知识点管理
type TkKnowledgePoint struct {
	ent.Schema
}

func (TkKnowledgePoint) Mixin() []ent.Mixin {
	return []ent.Mixin{
		CommonMixin{},
	}
}

func (TkKnowledgePoint) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Default("").Comment("知识点名称"),
		field.Int("question_bank_id").Default(0).Optional().Comment("题库id"),
		field.Int("question_count").Default(0).Comment("知识点题目数量"),
	}
}

func (TkKnowledgePoint) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("question_bank", TkQuestionBank.Type).Ref("knowledge_points").Field("question_bank_id").Unique(),
		edge.From("questions", TkQuestion.Type).Ref("knowledge_points"), //题目知识点（多）
	}
}
