package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

//课时练习
type KcSmallCategoryQuestion struct {
	ent.Schema
}

func (KcSmallCategoryQuestion) Mixin() []ent.Mixin {
	return []ent.Mixin{
		CommonMixin{},
	}
}

func (KcSmallCategoryQuestion) Fields() []ent.Field {
	return []ent.Field{
		field.Int("question_id").Optional().Comment("题目id"),
		field.Int("small_category_id").Optional().Comment("课时id"),
	}
}

func (KcSmallCategoryQuestion) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("question", TkQuestion.Type).Ref("small_category_questions").Field("question_id").Unique(),
		edge.From("course_small_category", KcCourseSmallCategory.Type).
			Ref("course_small_category_questions").Field("small_category_id").Unique(),
	}
}
