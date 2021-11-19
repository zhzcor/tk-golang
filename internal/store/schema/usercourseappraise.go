package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

//用户 课程 评价
type UserCourseAppraise struct {
	ent.Schema
}

func (UserCourseAppraise) Mixin() []ent.Mixin {
	return []ent.Mixin{
		CommonMixin{},
	}
}

func (UserCourseAppraise) Fields() []ent.Field {
	return []ent.Field{
		field.Float("teach_attitude_score").Default(0).Comment("授课态度评分"),
		field.Float("teach_content_score").Default(0).Comment("授课内容评分"),
		field.Float("teach_atmosphere_score").Default(0).Comment("授课氛围评分"),
		field.Float("composite_score").Default(0).Comment("综合评分"),
		field.Int("user_id").Optional().Comment("用户id"),
		field.Int("small_cate_id").Optional().Comment("课时id"),
		field.Int("course_id").Optional().Comment("课程id"),
		field.Uint8("type").Default(0).Comment("评价类型：1:普通课程，2:直播课程，3:直播公开课，4:录播公开课"),
		field.Uint8("show_status").Default(2).Comment("是否显示 1：显示 2：不显示"),
		field.String("teacher_impression").Default("").Comment("老师印象"),
		field.String("desc").Default("").Comment("评论描述"),
		field.String("teacher_reply").Default("").Comment("老师回复"),
	}
}

func (UserCourseAppraise) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("small_cate", KcCourseSmallCategory.Type).Ref("course_appraise_smalls").Field("small_cate_id").Unique(),
		edge.From("course", KcCourse.Type).Ref("course_appraise").Field("course_id").Unique(),
		edge.From("user", User.Type).Ref("course_appraise_users").Field("user_id").Unique(),
	}
}


func (UserCourseAppraise) Indexes() []ent.Index {
	return []ent.Index{
		// 非唯一约束索引
		//index.Fields("major_id"),
		index.Fields("small_cate_id"),
		index.Fields("course_id"),
		index.Fields("user_id"),
		//index.Fields("live_small_next_id"),
	}
}
