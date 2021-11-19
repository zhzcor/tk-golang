package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

//广告管理
type Advertise struct {
	ent.Schema
}

func (Advertise) Mixin() []ent.Mixin {
	return []ent.Mixin{
		CommonMixin{},
	}
}

func (Advertise) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Default("").Comment("名称"),
		field.Uint8("status").Default(1).Comment("状态：1、上线。2:下线"),
		field.Uint8("position").Default(1).Comment("广告位置：1、首页轮播图，2:圈子页面，3：我的页面，4:pc首页轮播图"),
		field.Time("start_at").Optional().Nillable().Comment("开始时间"),
		field.Time("end_at").Optional().Nillable().Comment("到期时间"),
		field.Int("click_count").Default(0).Comment("点击次数"),
		field.String("ad_url").Default("").Comment("广告链接"),
		field.String("remark").Default("").Comment("备注"),
		field.Int("attachment_id").Optional().Comment("广告图id"),
		field.Int("sort_order").Default(0).Comment("排序"),
	}
}

func (Advertise) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("attachment", Attachment.Type).Ref("advertise").Field("attachment_id").Unique(),
	}
}
