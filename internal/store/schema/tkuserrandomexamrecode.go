package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// TkUserRandomExamRecode holds the schema definition for the TkUserRandomExamRecode entity.
//用户随机卷记录
type TkUserRandomExamRecode struct {
	ent.Schema
}

func (TkUserRandomExamRecode) Mixin() []ent.Mixin {
	return []ent.Mixin{
		CommonMixin{},
	}
}

// Fields of the TkUserRandomExamRecode.
func (TkUserRandomExamRecode) Fields() []ent.Field {
	return []ent.Field{
		field.Int("user_id").Optional().Comment("用户id"),
		field.Int("exam_id").Optional().Comment("试卷id"),
	}
}

// Edges of the TkUserRandomExamRecode.
func (TkUserRandomExamRecode) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("random_exam_question", TkQuestion.Type),
		edge.From("exam_info", TkExamPaper.Type).Ref("user_random_exam").Field("exam_id").Unique(),
	}
}

func (TkUserRandomExamRecode) Indexes() []ent.Index {
	return []ent.Index{
		// 非唯一约束索引
		index.Fields("user_id"),
		index.Fields("exam_id"),
	}
}
