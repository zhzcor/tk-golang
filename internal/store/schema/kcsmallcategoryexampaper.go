package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

//课时试卷作业管理
type KcSmallCategoryExamPaper struct {
	ent.Schema
}

func (KcSmallCategoryExamPaper) Mixin() []ent.Mixin {
	return []ent.Mixin{
		CommonMixin{},
	}
}

func (KcSmallCategoryExamPaper) Fields() []ent.Field {
	return []ent.Field{
		field.Uint8("type").Default(0).Comment("分类，1：试卷，2：作业"),
		field.Int("exam_paper_id").Optional().Comment("试卷id"),
		field.Int("small_category_id").Optional().Comment("课时id"),
	}
}

func (KcSmallCategoryExamPaper) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("exam_paper", TkExamPaper.Type).Ref("course_exam_papers").Field("exam_paper_id").Unique(),
		edge.From("course_small_category", KcCourseSmallCategory.Type).Ref("course_small_category_exampapers").Field("small_category_id").Unique(),
	}
}
