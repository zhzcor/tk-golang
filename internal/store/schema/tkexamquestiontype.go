package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

//试卷分类表
type TkExamQuestionType struct {
	ent.Schema
}

func (TkExamQuestionType) Mixin() []ent.Mixin {
	return []ent.Mixin{
		CommonMixin{},
	}
}

func (TkExamQuestionType) Fields() []ent.Field {
	return []ent.Field{
		field.Uint8("exam_question_type").Default(0).Comment("试卷题目分类，1：模拟考试，2：考点练习，3：历年真题，4：通关必做300题，5：考前密押卷"),
		field.Int("question_count").Default(0).Comment("题目总数"),
		field.Int("answered_count").Default(0).Comment("已作答数"),
		field.Int("answered_user_count").Default(0).Comment("总答题人数"),
		field.Int("question_bank_id").Optional().Comment("题库id"),
	}
}

func (TkExamQuestionType) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("question_bank", TkQuestionBank.Type).Ref("exam_question_types").Field("question_bank_id").Unique(),
	}
}
