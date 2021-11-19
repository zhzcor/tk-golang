package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

//考试-模拟考试表 (弃用)
type TkExamPaperSimulation struct {
	ent.Schema
}

func (TkExamPaperSimulation) Mixin() []ent.Mixin {
	return []ent.Mixin{
		CommonMixin{},
	}
}

func (TkExamPaperSimulation) Fields() []ent.Field {
	return []ent.Field{
		field.Time("start_at").Optional().Comment("考试开始试卷"),
		field.Time("end_at").Optional().Comment("考试结束时间"),
		//field.Uint8("exam_status").Default(0).Comment("考试状态：1：未开考，2：考试中，3：已结束"),
		field.Uint8("enable_status").Default(0).Comment("启用状态 1：未启用，2：启用"),
		field.Int("exam_paper_id").Optional().Comment("试卷id"),
	}
}

func (TkExamPaperSimulation) Edges() []ent.Edge {
	return []ent.Edge{
		//edge.From("exam_paper", TkExamPaper.Type).Ref("exam_simulations").Field("exam_paper_id").Unique(),
	}
}
