package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// TkUserWrongQuestionRecode holds the schema definition for the TkUserWrongQuestionRecode entity.
//错题记录
type TkUserWrongQuestionRecode struct {
	ent.Schema
}

func (TkUserWrongQuestionRecode) Mixin() []ent.Mixin {
	return []ent.Mixin{
		CommonMixin{},
	}
}

// Fields of the TkUserWrongQuestionRecode.
func (TkUserWrongQuestionRecode) Fields() []ent.Field {
	return []ent.Field{
		field.Int("user_id").Optional().Comment("用户id"),
		field.Int("question_id").Optional().Comment("题目id"),
		field.Int("question_bank_id").Optional().Comment("题库id"),
		field.String("answer").Optional().Comment("答案"),
		field.Int("wrong_exam_type").Default(0).Comment("试卷题目分类，1：模拟考试，2：章节练习，3：历年真题，4：通关必做300题，5：考前密押卷"),
		field.Int("wrong_question_type").Default(0).Comment("题目类型，1：单选题，2：多选题，3：判断题，4：简答题，5：材料题"),
	}
}

// Edges of the TkUserWrongQuestionRecode.
func (TkUserWrongQuestionRecode) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("question_wrong", TkQuestion.Type).Ref("user_recode_wrong").Field("question_id").Unique(),
	}
}

func (TkUserWrongQuestionRecode) Indexes() []ent.Index {
	return []ent.Index{
		// 非唯一约束索引
		index.Fields("user_id"),
		index.Fields("question_id"),
	}
}
