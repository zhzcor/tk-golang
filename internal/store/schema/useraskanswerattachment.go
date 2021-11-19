package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

//用户 老师问答 附件
type UserAskAnswerAttachment struct {
	ent.Schema
}

// Fields of the KcCourse.
func (UserAskAnswerAttachment) Fields() []ent.Field {
	return []ent.Field{
		field.Int("attachment_id").Optional().Comment("问答附件id"),
		field.Int("ask_id").Optional().Comment("问答id"),
		field.Uint8("type").Default(1).Comment("图片类型 1：学生提问，2：老师回复"),
	}
}

// Edges of the Course.
func (UserAskAnswerAttachment) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("attachment", Attachment.Type).Ref("ask_attachments").Field("attachment_id").Unique(), //专业 (多)
		edge.From("ask", UserAskAnswer.Type).Ref("ask_answers_attachments").Field("ask_id").Unique(),    //专业 (多)
	}
}
