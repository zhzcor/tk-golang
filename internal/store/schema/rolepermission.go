package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

//角色权限
type RolePermission struct {
	ent.Schema
}

func (RolePermission) Mixin() []ent.Mixin {
	return []ent.Mixin{}
}

func (RolePermission) Fields() []ent.Field {
	return []ent.Field{
		field.Int("role_id").Optional().Comment("角色id"),
		field.Text("permission_id").Optional().Comment("权限ids"),
	}
}

func (RolePermission) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("role", Role.Type).Ref("role_permission").Field("role_id").Unique(),
	}
}
