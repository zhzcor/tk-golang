package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

//用户课程关联表
type KcUserCourse struct {
	ent.Schema
}

func (KcUserCourse) Mixin() []ent.Mixin {
	return []ent.Mixin{
		CommonMixin{},
	}
}

// Fields of the Class.
func (KcUserCourse) Fields() []ent.Field {
	return []ent.Field{
		field.Uint8("period_type").Default(3).Comment("有效期类型：1:长期有效，2:固定周期"),
		field.Time("closing_date").Optional().Comment("截止日期"),
		field.Float("study_rate").Default(0).Comment("学习进度"),
		field.String("remark").Default("").Default("").Comment("备注"),
		field.Float("price").Default(0).Comment("购买价格"),
		field.Int("user_id").Optional().Comment("用户id"),
		field.Int("course_id").Optional().Comment("课程id"),
	}
}

// Edges of the Class.
func (KcUserCourse) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("user_courses").Field("user_id").Unique(),
		edge.From("course", KcCourse.Type).Ref("kc_user_courses").Field("course_id").Unique(),
	}
}
