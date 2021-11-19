package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

//课程教师表
type KcCourseTeacher struct {
	ent.Schema
}

// Fields of the Class.
func (KcCourseTeacher) Fields() []ent.Field {
	return []ent.Field{
		field.Uint8("show_status").Default(1).Comment("显示状态：1：显示，2：不显示"),
		field.Int("sort_order").Default(0).Comment("排序"),
		field.Int("course_id").Optional().Comment("课程id"),
		field.Int("teacher_id").Optional().Comment("老师id"),
	}
}

// Edges of the Class.
func (KcCourseTeacher) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("teacher", Teacher.Type).Ref("teacher_courses").Field("teacher_id").Unique(),
		edge.From("course", KcCourse.Type).Ref("course_teachers").Field("course_id").Unique(),
	}
}
