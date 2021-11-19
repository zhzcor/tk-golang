package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

//活动
type Activity struct {
	ent.Schema
}

func (Activity) Mixin() []ent.Mixin {
	return []ent.Mixin{
		CommonMixin{},
	}
}

// Fields of the User.
func (Activity) Fields() []ent.Field {
	return []ent.Field{
		field.String("title").Default("").Comment("标题"),
		field.String("sub_title").Default("").Comment("副标题"),
		field.Int("cover_img_id").Default(0).Comment("封面图id"),
		field.Text("notice").Optional().Comment("活动须知"),
		field.Text("detail").Optional().Comment("活动详情"),
		field.String("place").Default("").Comment("活动地点"),
		field.Uint8("is_free").Default(0).Comment("是否免费，0：否，1：是"),
		field.Uint8("is_publish").Default(0).Comment("是否发布，0：否，1：是"),
		field.Int("amount").Default(0).Comment("金额"),
		field.Time("start_at").Optional().Comment("活动开始时间"),
		field.Time("end_at").Optional().Comment("活动结束时间"),
		field.Time("apply_start_at").Optional().Comment("报名开始时间"),
		field.Time("apply_end_at").Optional().Comment("报名结束时间"),
		field.Uint8("is_hot").Default(0).Comment("是否热门，0：否，1：是"),
		field.Uint8("is_auto_publish").Default(0).Comment("是否自动发布，0：否，1：是"),
		field.Int("apply_count").Default(0).Comment("报名人数"),
		field.Int("join_count").Default(0).Comment("参与人数"),
		field.Uint8("is_limit_join_count").Default(0).Comment("是否限制活动人数 0:否 1:是"),
		field.Time("birthday").Optional(),
		field.Text("sign_remark").Optional(),
		field.Int("activity_type_id").Optional().Comment("活动类型id"),
		field.Int("created_admin_id").Optional().Comment("创建人id"),
	}
}

// Edges of the User.
func (Activity) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("activity_type", ActivityType.Type).Ref("activities").Field("activity_type_id").Unique(),
		edge.To("apply_activities", ActivityApplyInfo.Type),
		edge.From("admin", Admin.Type).Ref("activities").Field("created_admin_id").Unique(),
	}
}
