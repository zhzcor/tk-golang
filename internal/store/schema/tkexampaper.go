package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

//试卷管理
type TkExamPaper struct {
	ent.Schema
}

func (TkExamPaper) Mixin() []ent.Mixin {
	return []ent.Mixin{
		CommonMixin{},
	}
}

func (TkExamPaper) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Default("").Comment("试卷名称"),
		field.String("desc").Default("").Comment("试卷描述"),
		field.Uint8("exam_question_type").Default(0).Comment("试卷题目分类，1：模拟考试，2：章节练习，3：历年真题，4：通关必做300题，5：考前密押卷"),
		field.Uint8("exam_type").Default(0).Comment("考试类型，1：固定卷，2：随机卷"),
		field.Uint8("difficulty").Default(0).Comment("难易度，1：易，2：较易，3：较难，4：难，5：一般"),
		//field.Time("practice_exam_at").Optional().Comment("模拟考试时间"),
		field.Int("question_count").Default(0).Comment("总题数"),
		//field.Int("answered_count").Default(0).Comment("已答题数"),
		field.Int("answered_user_count").Default(0).Comment("总答题人数"),
		field.Int("score").Default(0).Comment("总分数"),
		field.Int("pass_score").Default(0).Comment("通过分数"),
		field.Int("duration").Default(0).Comment("总时长"),
		field.Int("duration_type").Default(0).Comment("1:按照作答时间计时,2:设置试卷的答题时间"),
		field.Time("start_at").Optional().Comment("模拟考试开始时间"),
		field.Time("end_at").Optional().Comment("模拟考试结束时间"),
		//field.Uint8("exam_status").Default(0).Comment("考试状态：1：未开考，2：考试中，3：已结束"),
		field.Uint8("enable_status").Default(1).Comment("启用状态 1：启用，2：未启用"),
		field.Int("created_admin_id").Optional().Comment("创建人id"),
		field.Int("question_bank_id").Optional().Comment("题库id"),
	}
}

func (TkExamPaper) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("question_bank", TkQuestionBank.Type).Ref("exam_papers").Field("question_bank_id").Unique(),
		edge.From("admin", Admin.Type).Ref("admin_exam_papers").Field("created_admin_id").Unique(),
		edge.To("exam_partitions", TkExamPaperPartition.Type),
		edge.To("make_user_question_exam", MakeUserQuestionRecord.Type), //做题记录
		edge.To("user_exam_papers", TkUserExamScoreRecord.Type),
		edge.To("course_exam_papers", KcSmallCategoryExamPaper.Type), //课时试卷作业
		edge.To("exam_paper_records", TkUserQuestionRecord.Type),     //用户答题记录-试卷
		edge.To("user_random_exam", TkUserRandomExamRecode.Type),     //用户随机卷记录
	}
}
