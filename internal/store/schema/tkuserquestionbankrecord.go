package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

//用户刷题题库记录
type TkUserQuestionBankRecord struct {
	ent.Schema
}

func (TkUserQuestionBankRecord) Mixin() []ent.Mixin {
	return []ent.Mixin{
		CommonMixin{},
	}
}

func (TkUserQuestionBankRecord) Fields() []ent.Field {
	return []ent.Field{
		field.Int("record_count").Default(0).Comment("刷题总数"),
		field.Int("correct_count").Default(0).Comment("正确题数"),
		field.Int("wrong_count").Default(0).Comment("错误题数"),
		field.Float("correct_rate").Default(0).Comment("正确率"),
		field.Float("finish_rate").Default(0).Comment("完成率"),
		field.Int("user_id").Optional().Comment("用户id"),
		field.Int("question_bank_id").Optional().Comment("题库id"),
	}
}

func (TkUserQuestionBankRecord) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("question_bank", TkQuestionBank.Type).Ref("user_question_bank").Field("question_bank_id").Unique(),
		edge.From("user", User.Type).Ref("user_question_bank_records").Field("user_id").Unique(),
	}
}

func (TkUserQuestionBankRecord) Indexes() []ent.Index {
	return []ent.Index{
		// 非唯一约束索引
		//index.Fields("major_id"),
		index.Fields("question_bank_id"),
		index.Fields("user_id"),
		//index.Fields("live_small_next_id"),
	}
}
