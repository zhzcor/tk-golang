package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Course holds the schema definition for the Course entity.
//课程表
type KcCourse struct {
	ent.Schema
}

func (KcCourse) Mixin() []ent.Mixin {
	return []ent.Mixin{
		CommonMixin{},
	}
}

// Fields of the KcCourse.
func (KcCourse) Fields() []ent.Field {
	return []ent.Field{
		field.String("course_name").Default("").Comment("名称"),
		field.Uint8("course_type").Default(1).Comment("1:普通课程，2:直播课程，3:直播公开课，4:录播公开课"),
		field.Int("people_num").Default(0).Comment("学习人数"),
		field.Float("course_price").Default(0).Comment("价格"),
		field.Text("course_desc").Default("").Comment("简介"),
		field.Uint8("push_status").Default(2).Comment("1，发布 2：未发布 3 关闭"),
		//field.Uint8("period_type").Default(3).Comment("1:随到随学，2： 固定周期，3：长期有效"),
		//field.Time("open_live_start").Optional().Comment("公开课的直播开始时间"),
		//field.Int("open_live_duration").Optional().Comment("公开课直播时长"),
		//field.Time("start_date").Optional().Comment("开始日期"),
		//field.Time("end_date").Optional().Comment("结束日期"),
		//field.Time("closing_date").Optional().Comment("截止日期"),
		//field.Int("days_validity").Default(0).Comment("有效期天数"),
		//field.Int("live_small_next_id").Optional().Comment("当前小节id"),
		//field.Int("major_id").Optional().Comment("专业id"),//多选
		field.Int("cate_id").Optional().Comment("项目id"),
		field.Int("city_id").Optional().Comment("地区id"),
		field.Int("question_bank_id").Optional().Comment("题库id"),
		field.Int("course_cover_img_id").Optional().Comment("课程封面图id"),
		field.Int("created_admin_id").Optional().Comment("创建人id"),
	}
}

// Edges of the Course.
func (KcCourse) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("major", Major.Type).Ref("courses"),                                   //专业 (多)
		edge.From("item", ItemCategory.Type).Ref("course").Field("cate_id").Unique(),    //项目id
		edge.From("city", City.Type).Ref("course").Field("city_id").Unique(),            //地区id
		edge.From("admin", Admin.Type).Ref("course").Field("created_admin_id").Unique(), //创建人id
		//edge.From("liveNextSmall", KcCourseSmallCategory.Type).Ref("live_nexts_smalls").Field("live_small_next_id").Unique(),//当前直播小节id

		edge.From("question_bank", TkQuestionBank.Type).
			Ref("question_bank_courses").Field("question_bank_id").Unique(), //课程题库
		edge.From("attachment", Attachment.Type).
			Ref("course_attachments").Field("course_cover_img_id").Unique(), //课程封面图

		edge.To("course_teachers", KcCourseTeacher.Type),           //老师
		edge.From("classes", KcClass.Type).Ref("kc_class_courses"), //班级课程（多）

		edge.To("course_small_categorys", KcCourseSmallCategory.Type), //课程课时
		edge.To("course_appraise", UserCourseAppraise.Type), //课程id
		edge.To("course_chapters", KcCourseChapter.Type),
		edge.To("kc_user_courses", KcUserCourse.Type),
		edge.To("message_courses", Message.Type), //课程消息
	}
}

func (KcCourse) Indexes() []ent.Index {
	return []ent.Index{
		// 非唯一约束索引
		//index.Fields("major_id"),
		index.Fields("cate_id"),
		//index.Fields("live_small_next_id"),
	}
}
