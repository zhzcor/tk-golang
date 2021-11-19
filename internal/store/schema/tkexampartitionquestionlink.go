package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

//试卷 子部分 题目关联表
type TkExamPartitionQuestionLink struct {
	ent.Schema
}

func (TkExamPartitionQuestionLink) Mixin() []ent.Mixin {
	return []ent.Mixin{
		CommonMixin{},
	}
}

func (TkExamPartitionQuestionLink) Fields() []ent.Field {
	return []ent.Field{
		field.Uint8("question_score").Default(0).Comment("题目分值"),
		field.Int("exam_paper_partition_id").Optional().Comment("试卷子部分id"),
		field.Int("question_id").Optional().Comment("题目id"),
	}
}

func (TkExamPartitionQuestionLink) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("exam_paper_partition", TkExamPaperPartition.Type).Ref("exam_partition_links").Field("exam_paper_partition_id").Unique(),
		edge.From("question", TkQuestion.Type).Ref("exam_partition_questions").Field("question_id").Unique(),
	}
}
