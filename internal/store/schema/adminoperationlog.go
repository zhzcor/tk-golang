package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

//成员管理操作日志
type AdminOperationLog struct {
	ent.Schema
}

func (AdminOperationLog) Mixin() []ent.Mixin {
	return []ent.Mixin{
		CommonMixin{},
	}
}

func (AdminOperationLog) Fields() []ent.Field {
	return []ent.Field{
		field.String("ip").Default("").Comment("ip"),
		field.String("record").Default("").Comment("记录"),
		field.String("remark").Default("").Comment("备注"),
		field.Int("admin_id").Optional().Comment("管理员id"),
	}
}

func (AdminOperationLog) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("admin", Admin.Type).Ref("operation_admin_logs").Field("admin_id").Unique(),
	}
}
