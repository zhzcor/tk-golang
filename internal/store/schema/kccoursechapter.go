package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// CourseChapter holds the schema definition for the CourseChapter entity.
//课程节
type KcCourseChapter struct {
	ent.Schema
}

// Fields of the CourseChapter.
func (KcCourseChapter) Fields() []ent.Field {
	return []ent.Field{
		field.Int("course_id").Optional().Comment("课程id"),
		field.String("title").Default("").Comment("章节名称"),
	}
}

// Edges of the CourseChapter.
func (KcCourseChapter) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("chapter_course", KcCourse.Type).Ref("course_chapters").Field("course_id").Unique(), //课程-章
		edge.To("chapter_sections", KcCourseSection.Type),                                             //章-节
		edge.To("course_small_chapters", KcCourseSmallCategory.Type),                                  //章-课时
	}
}

func (KcCourseChapter) Indexes() []ent.Index {
	return []ent.Index{
		// 非唯一约束索引
		index.Fields("course_id"),
	}
}
