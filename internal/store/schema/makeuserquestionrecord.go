package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

//收藏
type MakeUserQuestionRecord struct {
	ent.Schema
}

func (MakeUserQuestionRecord) Mixin() []ent.Mixin {
	return []ent.Mixin{
		CommonMixin{},
	}
}

// Fields of the User.
func (MakeUserQuestionRecord) Fields() []ent.Field {
	return []ent.Field{
		field.Int("user_id").Optional().Comment("userId"),
		field.Int("exam_id").Optional().Comment("试卷id"),
		field.Int("sec_id").Optional().Comment("节id"),
		field.Int("question_bank_id").Optional().Comment("题库"),
		field.Int("exam_question_type").Optional().Comment("试卷题目分类，1：模拟考试，2：章节练习，3：历年真题，4：通关必做300题，5：考前密押卷"),
	}
}

// Edges of the User.
func (MakeUserQuestionRecord) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("section", TkSection.Type).Ref("make_user_question_sec").Field("sec_id").Unique(),
		edge.From("exam_paper", TkExamPaper.Type).Ref("make_user_question_exam").Field("exam_id").Unique(),
	}
}

func (MakeUserQuestionRecord) Indexes() []ent.Index {
	return []ent.Index{
		// 非唯一约束索引
		index.Fields("user_id"),
		index.Fields("sec_id"),
		index.Fields("question_bank_id"),
		index.Fields("exam_id"),
		index.Fields("exam_question_type"),
	}
}
