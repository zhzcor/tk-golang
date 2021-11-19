package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

//群名片
type GroupCard struct {
	ent.Schema
}

func (GroupCard) Mixin() []ent.Mixin {
	return []ent.Mixin{
		CommonMixin{},
	}
}

func (GroupCard) Fields() []ent.Field {
	return []ent.Field{
		field.String("title").Default("").Comment("名片标题"),
		field.String("sub_title").Default("").Comment("名片副标题"),
		field.Uint8("status").Default(1).Comment("状态：1、锁定。2:生效"),
		field.String("desc").Default("").Comment("描述"),
		field.Int("sort_order").Default(0).Comment("排序"),
		field.Int("attachment_id").Optional().Comment("图片id"),
	}
}

func (GroupCard) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("attachment", Attachment.Type).Ref("group_card").Field("attachment_id").Unique(),
	}
}
