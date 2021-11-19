package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

//课程上传任务
type KcVideoUploadTask struct {
	ent.Schema
}

func (KcVideoUploadTask) Mixin() []ent.Mixin {
	return []ent.Mixin{
		CommonMixin{},
	}
}

// Fields of the KcCourse.
func (KcVideoUploadTask) Fields() []ent.Field {
	return []ent.Field{
		field.Int("video_id").Optional().Comment("视频id"),
		field.Int("attachment_id").Optional().Comment("视频oss id"),
		field.Uint8("type").Default(1).Comment("类型 1：点播，2：伪直播 3：回放"),
		field.Int("course_id").Default(0).Comment("课程id"),
		field.Int("status").Default(1).Comment("任务状态:1:未开始，2：上传中，3：上传成功，4：上传失败，5：转码成功，6：转码失败，7：正在生成回放"),
		field.String("remark").Optional().Comment("任务描述"),
		field.String("title").Default("").Comment("标题"),
		field.String("video_name").Default("").Comment("video名称"),
		field.Int("total_size").Default(0).Comment("视频大小，单位是字节"),
		field.Int("length").Default(0).Comment("视频时长，单位为秒（转码成功回调才有）"),
	}
}

// Edges of the Course.
func (KcVideoUploadTask) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("attachment", Attachment.Type).
			Ref("video_task_attachment").Field("attachment_id").Unique(), //oss 附件
	}
}
