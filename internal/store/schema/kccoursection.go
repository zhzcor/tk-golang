package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

//课程节
type KcCourseSection struct {
	ent.Schema
}

func (KcCourseSection) Fields() []ent.Field {
	return []ent.Field{
		//field.Int("live_course_id").Default(0).Comment("直播课程id"),
		field.Int("course_chapter_id").Optional().Comment("直播章节id"),
		field.String("title").Default("").Comment("名称"),
		//field.Int("pid").Default(0).Comment("父章节id，0的话为章节，不为0就是小节"),
	}
}

func (KcCourseSection) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("chapter", KcCourseChapter.Type).Ref("chapter_sections").Field("course_chapter_id").Unique(), //章-节
		edge.To("course_small_sections", KcCourseSmallCategory.Type),                                           //节-课时
	}
}
