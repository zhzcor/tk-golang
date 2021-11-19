package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

//试卷 子部分管理
type TkExamPaperPartition struct {
	ent.Schema
}

func (TkExamPaperPartition) Mixin() []ent.Mixin {
	return []ent.Mixin{
		CommonMixin{},
	}
}

func (TkExamPaperPartition) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Default("").Comment("试卷部分名称"),
		field.String("desc").Default("").Comment("答题说明"),
		field.Int("duration").Default(0).Comment("作答时长，0为不限制"),
		field.Int("type").Default(0).Comment("题目类型，1：单选题，2：多选题，3：判断题，4：简答题，5：材料题"),
		field.Uint8("question_count").Default(0).Comment("题目数"),
		field.Int("exam_paper_id").Optional().Comment("试卷id"),
	}
}

func (TkExamPaperPartition) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("exam_paper", TkExamPaper.Type).Ref("exam_partitions").Field("exam_paper_id").Unique(),
		edge.To("exam_partition_links", TkExamPartitionQuestionLink.Type),
		edge.To("exam_partition_scores", TkExamPaperPartitionScore.Type),
	}
}
