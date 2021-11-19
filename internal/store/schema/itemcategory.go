package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

//项目分类
type ItemCategory struct {
	ent.Schema
}

func (ItemCategory) Mixin() []ent.Mixin {
	return []ent.Mixin{
		CommonMixin{},
	}
}

func (ItemCategory) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Default("").Comment("名称"),
		field.Uint8("status").Default(1).Comment("状态：1、锁定。2:生效"),
		field.String("code").Default("").Comment("编码"),
		field.String("desc").Default("").Comment("描述"),
		field.Int("pid").Optional().Comment("父级id"),
		field.Int("sort_order").Default(0).Comment("排序"),
	}
}

func (ItemCategory) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("item_question_banks", TkQuestionBank.Type),                         //题库所属项目
		edge.To("kc_class", KcClass.Type).Unique(),                                  //班级
		edge.To("course", KcCourse.Type),                                            //课程
		edge.To("user_item_cate", User.Type),                                        //课程
		edge.To("children", ItemCategory.Type).From("parent").Field("pid").Unique(), //pid

	}
}
