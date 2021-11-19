package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

//课时资料
type KcSmallCategoryAttachment struct {
	ent.Schema
}

func (KcSmallCategoryAttachment) Mixin() []ent.Mixin {
	return []ent.Mixin{
		CommonMixin{},
	}
}

func (KcSmallCategoryAttachment) Fields() []ent.Field {
	return []ent.Field{
		field.Int("attachment_id").Optional().Comment("资料id"),
		field.String("attachment_name").Default("").Comment("资料名称"),
		field.Int("small_category_id").Optional().Comment("课时id"),
	}
}

func (KcSmallCategoryAttachment) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("attachment", Attachment.Type).Ref("small_category_attachments").Field("attachment_id").Unique(),
		edge.From("small_category", KcCourseSmallCategory.Type).Ref("course_small_category_attachments").Field("small_category_id").Unique(),
	}
}
