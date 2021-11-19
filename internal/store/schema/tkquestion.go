package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

//题目
type TkQuestion struct {
	ent.Schema
}

func (TkQuestion) Mixin() []ent.Mixin {
	return []ent.Mixin{
		CommonMixin{},
	}
}

func (TkQuestion) Fields() []ent.Field {
	return []ent.Field{
		field.Text("name").Optional().Comment("题干"),
		field.Uint8("difficulty").Default(0).Comment("难易度，1：易，2：较易，3：较难，4：难，5：一般"),
		field.Uint8("type").Default(0).Comment("题目类型，1：单选题，2：多选题，3：判断题，4：简答题，5：材料题"),
		field.Text("desc").Optional().Comment("题目解析"),
		field.Int("created_admin_id").Optional().Comment("创建人id"),
		field.Int("question_bank_id").Optional().Comment("题库id"),
		field.Int("answer_count").Default(0).Comment("答题次数"),
		field.Int("right_count").Default(0).Comment("答对次数"),
		field.Int("pid").Optional().Comment("父级id"),
	}
}

func (TkQuestion) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("question_bank", TkQuestionBank.Type).Ref("questions").Field("question_bank_id").Unique(),
		edge.From("admin", Admin.Type).Ref("admin_Questions").Field("created_admin_id").Unique(),
		edge.From("user_rand_dom", TkUserRandomExamRecode.Type).Ref("random_exam_question"),
		edge.To("answer_options", TkQuestionAnswerOption.Type),
		edge.To("knowledge_points", TkKnowledgePoint.Type),        //题目知识点（多）
		edge.To("question_section_links", TkQuestionSection.Type), //题目节
		edge.To("exam_partition_questions", TkExamPartitionQuestionLink.Type),
		edge.To("question_error_feedbacks", TkQuestionErrorFeedback.Type),
		edge.To("user_records", TkUserQuestionRecord.Type),
		edge.To("small_category_questions", KcSmallCategoryQuestion.Type), //课程课时练习
		edge.To("user_exam_questions", TkUserSimulationTeacherMark.Type),
		edge.To("user_recode_wrong", TkUserWrongQuestionRecode.Type),
		edge.To("children", TkQuestion.Type).From("parent").Field("pid").Unique(), //pid
		edge.To("collection_question", Collection.Type),                           //收藏
	}
}
