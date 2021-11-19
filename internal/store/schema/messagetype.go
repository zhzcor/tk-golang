package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

//消息类型
type MessageType struct {
	ent.Schema
}

func (MessageType) Mixin() []ent.Mixin {
	return []ent.Mixin{
		CommonMixin{},
	}
}

func (MessageType) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Default("").Comment("名称"),
		field.Uint8("status").Default(1).Comment("状态：1、锁定。2:生效"),
	}
}

func (MessageType) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("messages", Message.Type),
	}
}
