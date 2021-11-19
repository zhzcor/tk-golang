package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

//专业
type Major struct {
	ent.Schema
}

func (Major) Mixin() []ent.Mixin {
	return []ent.Mixin{
		CommonMixin{},
	}
}

func (Major) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Default("").Comment("名称"),
		field.Uint8("status").Default(1).Comment("状态：1、锁定。2:生效"),
		field.String("code").Default("").Comment("编码"),
		field.String("desc").Default("").Comment("描述"),
		field.Int("sort_order").Default(0).Comment("排序"),
	}
}

func (Major) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("teachers", Teacher.Type).Ref("majors"), //老师-专业（多）
		edge.To("major_detail", MajorDetail.Type).Unique(),
		edge.To("kc_classes", KcClass.Type), //班级-专业（多）
		edge.To("courses", KcCourse.Type),   //课程-专业（多）

	}
}
