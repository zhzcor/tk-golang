package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// KcCourseVideo holds the schema definition for the CourseVideo entity.
//课程视频音频文件等
type KcCourseVideo struct {
	ent.Schema
}

// Fields of the KcCourseVideo.
func (KcCourseVideo) Fields() []ent.Field {
	return []ent.Field{
		field.String("video_title").Default("").Comment("名称"),
		field.Int8("file_type").Default(1).Comment("1:视频，2:音频，后面其他的自己定义拓展"),
		field.String("file_path").Default("").Comment("文件路径"),
		field.String("subtitle_path").Default("").Comment("字幕文件路径"),
		field.String("video_time").Default("").Comment("视频时长"),
	}
}

// Edges of the KcCourseVideo.
func (KcCourseVideo) Edges() []ent.Edge {
	return []ent.Edge{
		//edge.To("small_course", KcCourseVideo.Type).Unique(),//课程小节
		//
		//edge.To("live_small_course", KcCourseSmallCategory.Type).Unique(),
		//
		//edge.To("false_video", KcCourseSmallCategory.Type).Unique(),//伪直播视频id

	}
}

func (KcCourseVideo) Indexes() []ent.Index {
	return []ent.Index{
		// 非唯一约束索引
		index.Fields("video_title"),
	}
}
