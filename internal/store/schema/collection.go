package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

//收藏
type Collection struct {
	ent.Schema
}

func (Collection) Mixin() []ent.Mixin {
	return []ent.Mixin{
		CommonMixin{},
	}
}

// Fields of the User.
func (Collection) Fields() []ent.Field {
	return []ent.Field{
		field.Int("type").Optional().Comment("1,题目"),
		field.Int("value_id").Optional().Comment("type对应表的id"),
		field.Int("user_id").Optional().Comment("userId"),
		field.Int("exam_id").Optional().Comment("试卷id"),
		field.Int("sec_id").Optional().Comment("节id"),
		field.Int("question_bank_id").Optional().Comment("题库id"),
		field.Int("exam_question_type").Optional().Comment("试卷题目分类，1：模拟考试，2：章节练习，3：历年真题，4：通关必做300题，5：考前密押卷"),
	}
}

// Edges of the User.
func (Collection) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("question", TkQuestion.Type).Ref("collection_question").Field("value_id").Unique(),
	}
}

func (Collection) Indexes() []ent.Index {
	return []ent.Index{
		// 非唯一约束索引
		index.Fields("user_id"),
		index.Fields("value_id"),
		index.Fields("sec_id"),
		index.Fields("exam_id"),
		index.Fields("exam_question_type"),
	}
}
