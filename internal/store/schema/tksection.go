package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

//题库节
type TkSection struct {
	ent.Schema
}

func (TkSection) Mixin() []ent.Mixin {
	return []ent.Mixin{
		CommonMixin{},
	}
}

func (TkSection) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Default("").Comment("节名称"),
		field.Int("chapter_id").Optional().Comment("章id"),
		field.Int("question_count").Default(0).Comment("节题目数量"),
	}
}

func (TkSection) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("chapter", TkChapter.Type).Ref("sections").Field("chapter_id").Unique(),
		edge.To("tk_section_links", TkQuestionSection.Type),
		edge.To("section_records", TkUserQuestionRecord.Type),
		edge.To("user_section_exam", TkUserExamScoreRecord.Type),
		edge.To("make_user_question_sec", MakeUserQuestionRecord.Type),
	}
}
