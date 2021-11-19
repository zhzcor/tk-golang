package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Class holds the schema definition for the Class entity.
//班级管理
type KcClass struct {
	ent.Schema
}

func (KcClass) Mixin() []ent.Mixin {
	return []ent.Mixin{
		CommonMixin{},
	}
}

// Fields of the Class.
func (KcClass) Fields() []ent.Field {
	return []ent.Field{
		field.String("class_title").Default("").Comment("班级标题"),
		field.String("class_code").Default("").Comment("班级编号"),
		field.Text("class_desc").Default("").Comment("班级介绍"),
		field.Uint8("is_display").Default(1).Comment("是否展示1：展示 0：不展示"),
		field.Uint8("is_buy").Default(1).Comment("是否购买1：可以购买 0：不可以购买"),
		field.Uint8("class_period_type").Default(3).Comment("1:随到随学，2： 固定周期，3：长期有效"),
		field.Time("class_start_date").Optional().Comment("开始日期"),
		field.Time("class_end_date").Optional().Comment("结束日期"),
		field.Time("closing_date").Nillable().Optional().Comment("截止日期"),
		field.Int("days_validity").Default(0).Comment("有效期天数"),
		field.Int("class_head_master_id").Optional().Comment("班主任老师id"),
		field.Int("class_cover_img_id").Optional().Comment("封面图片id"),
		field.Uint8("status").Default(2).Comment("状态 1：已发布 2：未发布 3：已关闭"),
		field.Float("price").Default(0).Comment("价格"),
		field.Int("student_count").Default(0).Comment("班级人数"),
		field.Int("course_count").Default(0).Comment("课程数"),
		//field.Int("major_id").Optional().Comment("专业id"),//多专业
		field.Int("cate_id").Optional().Comment("项目id"),
		field.Int("city_id").Optional().Comment("地区id"),
		field.Int("created_admin_id").Optional().Comment("创建人id"),
	}
}

// Edges of the Class.
func (KcClass) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("majors", Major.Type).Ref("kc_classes"),                                    //班级专业（多）
		edge.From("item", ItemCategory.Type).Ref("kc_class").Field("cate_id").Unique(),       //项目id
		edge.From("city", City.Type).Ref("kc_class").Field("city_id").Unique(),               //地区id
		edge.From("admin", Admin.Type).Ref("class_admin").Field("created_admin_id").Unique(), //创建人id

		edge.From("attachment", Attachment.Type).
			Ref("class_cover_attachments").Field("class_cover_img_id").Unique(), //班级封面图
		edge.From("master_teachers", Teacher.Type).
			Ref("kc_class_masters").Field("class_head_master_id").Unique(), //班主任老师
		edge.To("class_teachers", KcClassTeacher.Type), //班级授课老师

		edge.To("kc_class_courses", KcCourse.Type),   //班级课程（多）
		edge.To("kc_user_classes", KcUserClass.Type), //用户班级
		edge.To("message_classes", Message.Type),     //班级消息

	}
}

func (KcClass) Indexes() []ent.Index {
	return []ent.Index{
		// 非唯一约束索引
		index.Fields("class_code"),
	}
}
