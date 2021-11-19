package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

//用户做题成绩表
type TkUserExamScoreRecord struct {
	ent.Schema
}

func (TkUserExamScoreRecord) Mixin() []ent.Mixin {
	return []ent.Mixin{
		CommonMixin{},
	}
}

func (TkUserExamScoreRecord) Fields() []ent.Field {
	return []ent.Field{
		field.Uint8("subjective_question_score").Default(0).Comment("主观题分数"),
		field.Uint8("objective_question_score").Default(0).Comment("客观题分数"),
		field.Uint8("total_score").Default(0).Comment("总分"),
		field.Int("duration").Default(0).Comment("答题时长"),
		field.Int("right_count").Default(0).Comment("答对题数"),
		field.Int("wrong_count").Default(0).Comment("答错题数"),
		field.Int("total_count").Default(0).Comment("总题数"),
		field.Int("no_answer_count").Default(0).Comment("未作答题数"),
		field.Int("rank").Default(0).Comment("排名"),
		field.Uint8("status").Default(0).Comment("判卷状态，1：未完成，2：已完成"),         //模拟考试用
		field.Uint8("order_status").Default(0).Comment("用户预约状态，1：未预约，2：已预约"), //模拟考试
		field.Int("exam_paper_id").Optional().Comment("试卷id"),
		field.Int("section_id").Optional().Comment("节id"),
		field.Int("user_id").Optional().Comment("学生用户id"),
		field.Int("operation_teacher_id").Optional().Comment("老师id"),
	}
}

func (TkUserExamScoreRecord) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("exam_paper", TkExamPaper.Type).Ref("user_exam_papers").Field("exam_paper_id").Unique(),
		edge.From("section", TkSection.Type).Ref("user_section_exam").Field("section_id").Unique(),
		edge.From("teacher", Teacher.Type).Ref("user_exams_teachers").Field("operation_teacher_id").Unique(),
		edge.From("user", User.Type).Ref("user_exams_records").Field("user_id").Unique(),
		edge.To("user_exam_details", TkUserSimulationTeacherMark.Type),
	}
}
