package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

//专业详情标签
type MajorDetailTag struct {
	ent.Schema
}

func (MajorDetailTag) Mixin() []ent.Mixin {
	return []ent.Mixin{
		CommonMixin{},
	}
}

func (MajorDetailTag) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Default("").Comment("标签名称"),
		field.Int("major_detail_id").Optional().Comment("专业详情id"),
	}
}

func (MajorDetailTag) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("major_detail", MajorDetail.Type).Ref("major_detail_tags").Field("major_detail_id").Unique(),
	}
}
