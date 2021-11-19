package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

//管理员登录日志
type AdminLoginLog struct {
	ent.Schema
}

func (AdminLoginLog) Mixin() []ent.Mixin {
	return []ent.Mixin{
		CommonMixin{},
	}
}

// Fields of the User.
func (AdminLoginLog) Fields() []ent.Field {
	return []ent.Field{
		field.String("ip").Default("").Comment("登录ip"),
		field.String("city").Default("").Comment("登录地区"),
		field.String("browser").Default("").Comment("登录浏览器"),
		field.Int("admin_id").Optional().Comment("管理员id"),
	}
}

// Edges of the User.
func (AdminLoginLog) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("admin", Admin.Type).Ref("admin_login_logs").Field("admin_id").Unique(),
	}
}
