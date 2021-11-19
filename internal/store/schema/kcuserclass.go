package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

//用户班级关联表
type KcUserClass struct {
	ent.Schema
}

func (KcUserClass) Mixin() []ent.Mixin {
	return []ent.Mixin{
		CommonMixin{},
	}
}

// Fields of the Class.
func (KcUserClass) Fields() []ent.Field {
	return []ent.Field{
		field.Uint8("period_type").Default(3).Comment("1:长期有效，2： 固定周期"),
		field.Time("closing_date").Optional().Comment("截止日期"),
		field.Float("study_rate").Default(0).Comment("学习进度"),
		field.String("remark").Default("").Comment("备注"),
		field.Float("price").Default(0).Comment("购买价格"),
		field.Int("user_id").Optional().Comment("用户id"),
		field.Int("class_id").Optional().Comment("班级id"),
	}
}

// Edges of the Class.
func (KcUserClass) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("user_classes").Field("user_id").Unique(),
		edge.From("class", KcClass.Type).Ref("kc_user_classes").Field("class_id").Unique(), //班级课程
	}
}
