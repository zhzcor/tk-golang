package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

//班级教师表
type KcClassTeacher struct {
	ent.Schema
}

// Fields of the Class.
func (KcClassTeacher) Fields() []ent.Field {
	return []ent.Field{
		field.Uint8("show_status").Default(1).Comment("显示状态：1：显示，2：不显示"),
		field.Int("sort_order").Default(0).Comment("排序"),
		field.Int("class_id").Optional().Comment("班级"),
		field.Int("teacher_id").Optional().Comment("老师id"),
	}
}

// Edges of the Class.
func (KcClassTeacher) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("teacher", Teacher.Type).Ref("teacher_classes").Field("teacher_id").Unique(),
		edge.From("class", KcClass.Type).Ref("class_teachers").Field("class_id").Unique(),
	}
}
