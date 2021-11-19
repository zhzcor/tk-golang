package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

//分享海报
type SharePoster struct {
	ent.Schema
}

func (SharePoster) Mixin() []ent.Mixin {
	return []ent.Mixin{
		CommonMixin{},
	}
}

func (SharePoster) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Default("").Comment("标题"),
		field.Uint8("status").Default(1).Comment("状态：1、锁定。2:生效"),
		field.String("url").Default("").Comment("海报链接"),
		field.String("remark").Default("").Comment("备注"),
		field.Int("share_poster_img_id").Optional().Comment("封面图id"),
		field.Int("share_count").Default(0).Comment("分享次数"),
		field.Int("sort_order").Default(0).Comment("排序"),
	}
}

func (SharePoster) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("attachment", Attachment.Type).Ref("share_poster_attachments").Field("share_poster_img_id").Unique(),
	}
}
