package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

//用户模拟考试主观题判分表
type TkUserSimulationTeacherMark struct {
	ent.Schema
}

func (TkUserSimulationTeacherMark) Mixin() []ent.Mixin {
	return []ent.Mixin{
		CommonMixin{},
	}
}

func (TkUserSimulationTeacherMark) Fields() []ent.Field {
	return []ent.Field{
		field.Uint8("score").Default(0).Comment("主观题分数"),
		field.Text("desc").Optional().Comment("老师解析"),
		field.Int("question_id").Optional().Comment("题目id"),
		field.Int("user_exam_record_id").Optional().Comment("考试成绩表id"),
	}
}

func (TkUserSimulationTeacherMark) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("question", TkQuestion.Type).Ref("user_exam_questions").Field("question_id").Unique(),
		edge.From("user_exam_record", TkUserExamScoreRecord.Type).Ref("user_exam_details").Field("user_exam_record_id").Unique(),
	}
}
