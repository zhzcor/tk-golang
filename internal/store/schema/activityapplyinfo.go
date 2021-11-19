package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

//活动报名学员信息
type ActivityApplyInfo struct {
	ent.Schema
}

func (ActivityApplyInfo) Mixin() []ent.Mixin {
	return []ent.Mixin{
		CommonMixin{},
	}
}

func (ActivityApplyInfo) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Default("").Comment("学员姓名"),
		field.String("phone").Default("").Comment("手机号"),
		field.String("remark").Default("").Comment("备注"),
		field.Uint8("is_our_student").Default(0).Comment("是否明德学员，0：否，1：是"),
		field.Int("activity_id").Optional().Comment("活动id"),
	}
}

func (ActivityApplyInfo) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("activity", Activity.Type).Ref("apply_activities").Field("activity_id").Unique(),
	}
}
