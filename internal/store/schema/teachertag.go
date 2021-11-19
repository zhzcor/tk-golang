package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

//教师标签
type TeacherTag struct {
	ent.Schema
}

func (TeacherTag) Mixin() []ent.Mixin {
	return []ent.Mixin{
		CommonMixin{},
	}
}

func (TeacherTag) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Default("").Comment("标签名称"),
		field.Int("teacher_id").Optional().Comment("讲师id"),
	}
}

func (TeacherTag) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("teacher", Teacher.Type).Ref("teacher_tags").Field("teacher_id").Unique(),
	}
}
