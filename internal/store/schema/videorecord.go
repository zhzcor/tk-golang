package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"time"
)

//视频观看记录
type VideoRecord struct {
	ent.Schema
}

func (VideoRecord) Mixin() []ent.Mixin {
	return []ent.Mixin{
		CommonMixin{},
	}
}

func (VideoRecord) Fields() []ent.Field {
	return []ent.Field{
		field.String("video_name").Default("").Comment("视频名称"),
		field.Uint8("view_time").Optional().Comment("观看时长 s"),
		field.Int("user_id").Optional().Comment("学生用户id"),
		field.Int("small_id").Optional().Comment("小节id"),
		field.Int("video_time").Optional().Comment("视频时长"),
		field.Time("view_at").Default(time.Now).Nillable().Optional().Comment("最近一次观看时间"),
	}
}

func (VideoRecord) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("user_video_record").Field("user_id").Unique(), //关联用户
		edge.From("small_course", KcCourseSmallCategory.Type).Ref("video_record_small").Field("small_id").Unique(), //关联小节
	}
}

func (VideoRecord) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("small_id"),
		index.Fields("user_id"),
	}
}
