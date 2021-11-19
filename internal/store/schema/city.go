package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

//地区
type City struct {
	ent.Schema
}

func (City) Mixin() []ent.Mixin {
	return []ent.Mixin{
		CommonMixin{},
	}
}

func (City) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Default("").Comment("名称"),
		field.Uint8("status").Default(1).Comment("状态：1、锁定。2:生效"),
		field.String("code").Default("").Comment("编码"),
		field.String("desc").Default("").Comment("描述"),
		field.Int("sort_order").Default(0).Comment("排序"),
	}
}

func (City) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("kc_class", KcClass.Type), //班级
		edge.To("course", KcCourse.Type),  //课程
		edge.To("user_city", User.Type),   //用户地区

	}
}
