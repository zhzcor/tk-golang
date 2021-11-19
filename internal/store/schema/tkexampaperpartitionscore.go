package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

//试卷 子部分分值管理
type TkExamPaperPartitionScore struct {
	ent.Schema
}

func (TkExamPaperPartitionScore) Mixin() []ent.Mixin {
	return []ent.Mixin{
		CommonMixin{},
	}
}

func (TkExamPaperPartitionScore) Fields() []ent.Field {
	return []ent.Field{
		field.Uint8("singe_select").Default(0).Comment("单选题分值"),
		field.Uint8("judge_question").Default(0).Comment("判断题分值"),
		field.Uint8("multi_select").Default(0).Comment("多选题分值"),
		field.Uint8("shorter_answer").Default(0).Comment("简答题分值"),
		field.Uint8("material_question").Default(0).Comment("材料题分值"),
		field.Uint8("singe_select_count").Default(0).Comment("单选题数量"),
		field.Uint8("judge_question_count").Default(0).Comment("判断题数量"),
		field.Uint8("multi_select_count").Default(0).Comment("多选题数量"),
		field.Uint8("shorter_answer_count").Default(0).Comment("简答题数量"),
		field.Uint8("material_question_count").Default(0).Comment("材料题数量"),
		field.Int("exam_paper_partition_id").Optional().Comment("试卷子部分id"),
	}
}

func (TkExamPaperPartitionScore) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("exam_paper_partition", TkExamPaperPartition.Type).Ref("exam_partition_scores").Field("exam_paper_partition_id").Unique(),
	}
}
