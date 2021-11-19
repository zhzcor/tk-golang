package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

//权限表
type Permission struct {
	ent.Schema
}

func (Permission) Mixin() []ent.Mixin {
	return []ent.Mixin{
		CommonMixin{},
	}
}

func (Permission) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Default("").Comment("权限名称"),
		field.String("desc").Default("").Comment("描述"),
		field.Uint32("pid").Default(0).Comment("父级id"),
	}
}

func (Permission) Edges() []ent.Edge {
	return []ent.Edge{}
}
