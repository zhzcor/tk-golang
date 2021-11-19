package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// CourseSmallCategory holds the schema definition for the CourseSmallCategory entity.
//课程小节表
type KcCourseSmallCategory struct {
	ent.Schema
}

func (KcCourseSmallCategory) Mixin() []ent.Mixin {
	return []ent.Mixin{
		CommonMixin{},
	}
}

// Fields of the CourseSmallCategory.
func (KcCourseSmallCategory) Fields() []ent.Field {
	return []ent.Field{
		//field.Int("small_next_id").Optional().Comment("下一节直播小节课id"),
		field.String("small_name").Default("").Comment("小节课程名称"),
		field.Int("viewing_time").Default(0).Comment("0分钟表示完成到最后"),
		field.Uint8("finish_type").Default(1).
			Comment("完成条件：1：视频观看到最后，2：视频观看时长，3：音频收听到最后，4：音频收听时长，5：进入过直播间，6：下载过资料"),
		field.Uint8("teach_type").Default(1).Comment("1:大班，2:小班"),
		field.Uint8("type").Default(1).Comment("课时类型：1：视频，2：音频，3：直播，4：资料下载"),

		field.Time("live_small_start").Optional().Comment("直播开始时间"),
		field.Int8("live_small_status").Default(0).Comment(
			"直播状态1.未开始，2:直播中，3：直播结束，4：回放生成中，5：已生成回放，6回放生成失败，7：已上传录像"),
		field.Int("live_room_id").Optional().Comment("直播房间id"),
		field.Int("back_video_id").Optional().Comment("回放video_id"),
		field.Int("false_video_id").Optional().Comment("伪直播video_id"),
		field.Int("order_video_id").Optional().Comment("点播video_id"),
		field.Int("order_video_attach_id").Optional().Comment("点播视频oss 附件"),
		field.Int("live_small_time").Default(0).Comment("直播时长"),
		field.Uint8("push_status").Default(2).Comment("1，发布 2：未发布 3 关闭"),
		field.String("live_small_remark").Default("").Comment("直播说明"),
		field.Int("courseware_attach_id").Optional().Comment("资料下载附件ID"),
		field.String("courseware_name").Default("").Comment("资料名称"),
		field.Int("attachment_count").Default(0).Comment("课时资料数量"),
		field.Int("question_count").Default(0).Comment("课时练习数量"),
		field.Int("exam_count").Default(0).Comment("课时考试试卷数量"),
		field.Int("homework_count").Default(0).Comment("课时作业数量"),
		field.Int("study_count").Default(0).Comment("学习人数"),
		field.Int("finish_count").Default(0).Comment("完成人数"),
		field.Int("average_view_duration").Default(0).Comment("平均观看时长"),

		field.Int("course_id").Optional().Comment("课程id"),
		field.Int("chapter_id").Optional().Comment("章id"),
		field.Int("section_id").Optional().Comment("节id"),
	}
}

// Edges of the CourseSmallCategory.
func (KcCourseSmallCategory) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("course", KcCourse.Type).
			Ref("course_small_categorys").Field("course_id").Unique(), //课程
		edge.From("chapter", KcCourseChapter.Type).
			Ref("course_small_chapters").Field("chapter_id").Unique(), //关联章id
		edge.From("section", KcCourseSection.Type).
			Ref("course_small_sections").Field("section_id").Unique(), //关联小节id

		edge.To("course_small_category_attachments", KcSmallCategoryAttachment.Type), //课时资料（多）
		edge.To("course_small_category_exampapers", KcSmallCategoryExamPaper.Type),   //课时考试任务（多）
		edge.To("course_small_category_questions", KcSmallCategoryQuestion.Type),     //课时练习（多）
		edge.To("course_appraise_smalls", UserCourseAppraise.Type),                   //用户 课程 评价
		edge.To("video_record_small", VideoRecord.Type),                   //用户观看视频记录

		edge.From("cs_attachment", Attachment.Type).
			Ref("courseware_attachment").Field("courseware_attach_id").Unique(), //资料下载
		//edge.From("fl_attachment", Attachment.Type).
		//	Ref("false_live_attachment").Field("false_video_attach_id").Unique(), //伪直播视频oss 附件
		edge.From("od_attachment", Attachment.Type).
			Ref("order_attachment").Field("order_video_attach_id").Unique(), //点播视频oss 附件

	}
}

func (KcCourseSmallCategory) Indexes() []ent.Index {
	return []ent.Index{
		// 非唯一约束索引
		index.Fields("course_id"),
	}
}
