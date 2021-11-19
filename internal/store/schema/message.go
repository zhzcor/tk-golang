package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

//消息管理
type Message struct {
	ent.Schema
}

func (Message) Mixin() []ent.Mixin {
	return []ent.Mixin{
		CommonMixin{},
	}
}

func (Message) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Default("").Comment("消息名称"),
		field.String("detail").Default("").Comment("内容详情"),
		field.Uint8("publish_type").Default(0).Comment("发布类型 1：定时发布，2：实时发布"),
		field.Uint8("status").Default(2).Comment("状态 1：启用，2：未启用"),
		field.Uint8("publish_status").Default(2).Comment("发送状态 1：未发送，2：发送中，3：已发送"),
		field.Time("auto_publish_at").Optional().Nillable().Comment("定时发布时间"),
		field.Int("message_type_id").Optional().Comment("消息类型id"),
		field.Int("created_admin_id").Optional().Comment("创建人id"),
		field.Int("attachment_id").Optional().Comment("封面图id"),
		field.Int("course_id").Optional().Comment("消息所属课程id"),
		field.Int("class_id").Optional().Comment("消息所属班级id"),
	}
}

func (Message) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("Message_type", MessageType.Type).Ref("messages").Field("message_type_id").Unique(),
		edge.From("admin", Admin.Type).Ref("message_created_admin_id").Field("created_admin_id").Unique(),
		edge.From("attachment", Attachment.Type).Ref("message_attachment").Field("attachment_id").Unique(),
		edge.From("course", KcCourse.Type).Ref("message_courses").Field("course_id").Unique(),
		edge.From("class", KcClass.Type).Ref("message_classes").Field("class_id").Unique(),
		edge.From("users", User.Type).Ref("messages"),
	}
}
