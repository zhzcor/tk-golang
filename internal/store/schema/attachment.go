package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

//附件
type Attachment struct {
	ent.Schema
}

func (Attachment) Mixin() []ent.Mixin {
	return []ent.Mixin{
		CommonMixin{},
	}
}

func (Attachment) Fields() []ent.Field {
	return []ent.Field{
		field.String("filename").Default("").Comment("文件名称"),
		field.Uint32("file_size").Default(0).Comment("文件大小,字节"),
		field.String("mime_type").Default("").Comment("文件类型"),
		field.Int("user_id").Optional().Comment("上传时的用户ID"),
		field.Int("admin_id").Optional().Comment("如果是后台上传，上传时的管理员ID"),
	}
}

func (Attachment) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("major_detail_cover_img", MajorDetail.Type),                   //专业封面图
		edge.To("major_detail_subject_img", MajorDetail.Type),                 //专业科目图
		edge.To("major_teacher_attachment", MajorDetail.Type),                 //专业讲师图（多）
		edge.To("major_service_attachment", MajorDetail.Type),                 //专业服务保障图（多）
		edge.To("advertise", Advertise.Type),                                  //广告图
		edge.To("message_attachment", Message.Type),                           //消息
		edge.To("share_poster_attachments", SharePoster.Type),                 //分享海报
		edge.To("teacher_attachments", Teacher.Type),                          //讲师头像
		edge.To("admin_img_id", Admin.Type).Unique(),                          //管理员头像
		edge.To("small_category_attachments", KcSmallCategoryAttachment.Type), //课时资料
		edge.To("course_attachments", KcCourse.Type),                          //课程封面图
		edge.To("class_cover_attachments", KcClass.Type),                      //班级封面图
		edge.To("courseware_attachment", KcCourseSmallCategory.Type).Unique(), //课时资料下载
		//edge.To("false_live_attachment", KcCourseSmallCategory.Type).Unique(), //课时伪直播视频oss 附件
		edge.To("order_attachment", KcCourseSmallCategory.Type).Unique(), //课时点播视频oss 附件
		edge.To("video_task_attachment", KcVideoUploadTask.Type),         //video task 视频附件
		edge.To("ask_attachments", UserAskAnswerAttachment.Type),         //用户问答附件

	}
}
