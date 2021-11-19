package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

//专业详情
type MajorDetail struct {
	ent.Schema
}

func (MajorDetail) Mixin() []ent.Mixin {
	return []ent.Mixin{
		CommonMixin{},
	}
}

func (MajorDetail) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Default("").Comment("名称"),
		field.Text("desc").Optional().Comment("专业简介"),
		field.Uint8("subject_count").Default(0).Comment("科目数量"),
		field.String("star").Default("0").Comment("通关系数，0-5星"),
		field.Uint8("pass_rate").Default(0).Comment("通过率"),
		field.Uint32("student_count").Default(0).Comment("学习人数"),
		field.Uint8("study_duration").Default(0).Comment("学习时长"),
		field.Int("major_id").Optional().Comment("专业id"),
		field.Int("cover_img_id").Optional().Comment("封面图id"),
		field.Int("subject_img_id").Optional().Comment("科目图id"),
	}
}

func (MajorDetail) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("major_detail_tags", MajorDetailTag.Type),
		edge.From("cover_attachment", Attachment.Type).Ref("major_detail_cover_img").Field("cover_img_id").Unique(),
		edge.From("subject_attachment", Attachment.Type).Ref("major_detail_subject_img").Field("subject_img_id").Unique(),
		edge.From("teacher_attachments", Attachment.Type).Ref("major_teacher_attachment"),
		edge.From("service_attachments", Attachment.Type).Ref("major_service_attachment"),
		edge.From("major", Major.Type).Ref("major_detail").Field("major_id").Unique(),
	}
}
