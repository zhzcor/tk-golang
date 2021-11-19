package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// UserLoginLog holds the schema definition for the UserLoginLog entity.
type UserLoginLog struct {
	ent.Schema
}

// Fields of the UserLoginLog.
func (UserLoginLog) Fields() []ent.Field {
	return []ent.Field{
		field.String("cid").Default(""),
		field.String("platform").Default(""),
		field.String("device").Default(""),
		field.String("version").Default(""),
		field.String("ip"),
		field.String("province"),
		field.String("city"),
		field.Time("latest_login_at").Optional().Comment("上次登录日期"),
	}
}

// Edges of the UserLoginLog.
func (UserLoginLog) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("login_log").Required().Unique(),
	}
}

func (l UserLoginLog) Mixin() []ent.Mixin {
	return []ent.Mixin{
		CommonMixin{},
	}
}
