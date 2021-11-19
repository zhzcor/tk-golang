package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

//用户做题记录
type TkUserQuestionRecord struct {
	ent.Schema
}

func (TkUserQuestionRecord) Mixin() []ent.Mixin {
	return []ent.Mixin{
		CommonMixin{},
	}
}

func (TkUserQuestionRecord) Fields() []ent.Field {
	return []ent.Field{
		field.Int("correct_count").Default(0).Comment("正确次数"),
		field.Int("answer_count").Default(0).Comment("答题次数"),
		field.String("answer").Default("").Comment("答案 例：多选题 A,B"),
		field.Uint8("is_right").Default(1).Comment("是否正确 1：否，2：是"),
		field.Uint8("exam_question_type").Default(0).Comment("试卷题目分类，1：模拟考试，2：考点练习，3：历年真题，4：通关必做300题，5：考前密押卷"),
		field.Uint8("question_type").Default(0).Comment("题目类型，1：单选题，2：多选题，3：判断题，4：简答题，5：材料题"),
		field.Int("user_id").Optional().Comment("用户id"),
		field.Int("question_bank_id").Optional().Comment("题库id"),
		field.Int("question_id").Optional().Comment("题目id"),
		field.Int("exam_paper_id").Optional().Comment("试卷id"),
		field.Int("section_id").Optional().Comment("节id"),
	}
}

func (TkUserQuestionRecord) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("question_bank", TkQuestionBank.Type).Ref("user_bank_records").Field("question_bank_id").Unique(),
		edge.From("user", User.Type).Ref("user_question_records").Field("user_id").Unique(),
		edge.From("question", TkQuestion.Type).Ref("user_records").Field("question_id").Unique(),
		edge.From("exam_paper", TkExamPaper.Type).Ref("exam_paper_records").Field("exam_paper_id").Unique(),
		edge.From("section", TkSection.Type).Ref("section_records").Field("section_id").Unique(),
	}
}
