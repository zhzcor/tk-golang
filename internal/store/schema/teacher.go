package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

//教师
type Teacher struct {
	ent.Schema
}

func (Teacher) Mixin() []ent.Mixin {
	return []ent.Mixin{
		CommonMixin{},
	}
}

func (Teacher) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Default("").Comment("教师名称"),
		field.Uint8("sex").Default(0).Comment("性别，1：男，2：女，3：未知"),
		field.String("email").Default("").Comment("邮箱"),
		field.String("phone").Default("").Comment("手机号"),
		field.String("nickname").Default("").Comment("昵称"),
		field.String("sub_title").Default("").Comment("副标题"),
		field.Text("detail").Optional().Comment("详情"),
		field.Uint8("status").Default(1).Comment("是否启用  1：启用，2：未启用"),
		field.Uint8("teaching_age").Default(0).Comment("教龄"),
		field.Int("avatar_id").Optional().Comment("教师头像id"),
		field.Int("sort_order").Default(0).Comment("排序"),
	}
}

func (Teacher) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("majors", Major.Type), //老师-专业（多）
		edge.To("teacher_tags", TeacherTag.Type),
		edge.To("teacher_classes", KcClassTeacher.Type),            //关联班级老师
		edge.To("kc_class_masters", KcClass.Type),                  //关联班级班主任
		edge.To("teacher_courses", KcCourseTeacher.Type),           //关联课程老师
		edge.To("user_exams_teachers", TkUserExamScoreRecord.Type), //老师模拟考试成绩表
		edge.To("ask_teachers", UserAskAnswer.Type),                //用户老师问答

		edge.From("attachment", Attachment.Type).Ref("teacher_attachments").Field("avatar_id").Unique(), //讲师头像
	}
}
