package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

//层次
type Level struct {
	ent.Schema
}

func (Level) Mixin() []ent.Mixin {
	return []ent.Mixin{
		CommonMixin{},
	}
}

func (Level) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Default("").Comment("名称"),
		field.Uint8("status").Default(1).Comment("状态：1、锁定。2:生效"),
		field.String("code").Default("").Comment("编码"),
		field.String("desc").Default("").Comment("描述"),
		field.Int("sort_order").Default(0).Comment("排序"),
	}
}

func (Level) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("level_question_banks", TkQuestionBank.Type), //题库所属层次
	}
}
