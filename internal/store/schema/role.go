package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

//角色表
type Role struct {
	ent.Schema
}

func (Role) Mixin() []ent.Mixin {
	return []ent.Mixin{
		CommonMixin{},
	}
}

func (Role) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Default("").Comment("角色名称"),
		field.String("desc").Default("").Comment("描述"),
		field.Uint8("status").Default(1).Comment("状态 1：开启 2：关闭"),
	}
}

func (Role) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("admins", Admin.Type).Ref("roles"),
		edge.To("role_permission", RolePermission.Type).Unique(),
	}
}
