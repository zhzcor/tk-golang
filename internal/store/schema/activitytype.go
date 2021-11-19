package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

//活动类型
type ActivityType struct {
	ent.Schema
}

func (ActivityType) Mixin() []ent.Mixin {
	return []ent.Mixin{
		CommonMixin{},
	}
}

// Fields of the User.
func (ActivityType) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Default("").Comment("活动类型名称"),
		field.Uint8("status").Default(1).Comment("状态：1、锁定 2:启用"),
	}
}

// Edges of the User.
func (ActivityType) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("activities", Activity.Type),
	}
}
