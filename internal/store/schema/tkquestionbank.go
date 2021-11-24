package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

//题库管理
type TkQuestionBank struct {
	ent.Schema
}

func (TkQuestionBank) Mixin() []ent.Mixin {
	return []ent.Mixin{
		CommonMixin{},
	}
}

func (TkQuestionBank) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Default("").Comment("题库名称"),
		field.Uint8("status").Default(1).Comment("状态，1：启用 2:禁用"),
		field.Int("question_count").Default(0).Comment("题目数量"),
		//field.Float("correct_rate").Default(0).Comment("正确率"),
		field.Int("created_admin_id").Optional().Comment("创建人id"),
		field.Int("item_category_id").Optional().Comment("项目id"),
		field.Int("level_id").Optional().Comment("层次id"),
		field.Int("sort_order").Default(0).Comment("排序"),
	}
}

func (TkQuestionBank) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("item_category", ItemCategory.Type).Ref("item_question_banks").Field("item_category_id").Unique(),
		edge.From("level", Level.Type).Ref("level_question_banks").Field("level_id").Unique(),
		edge.From("admin", Admin.Type).Ref("admin_question_banks").Field("created_admin_id").Unique(),
		edge.To("question_chapters", TkChapter.Type),
		edge.To("question_bank_courses", KcCourse.Type),
		edge.To("questions", TkQuestion.Type), //题库题目（多）
		edge.To("exam_papers", TkExamPaper.Type),
		edge.To("exam_question_types", TkExamQuestionType.Type),
		edge.To("user_question_bank", TkUserQuestionBankRecord.Type), //用户题库刷题记录（一对一）
		edge.To("user_bank_records", TkUserQuestionRecord.Type),      //用户做题记录
		edge.To("knowledge_points", TkKnowledgePoint.Type),           //题库知识点
		edge.To("city_question_banks", TkQuestionBankCity.Type),      //题库地区中间表
		edge.To("major_question_banks", TkQuestionBankMajor.Type),    //题库专业中间表
	}
}
