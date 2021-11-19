package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

//资讯分类
type InformationClassify struct {
	ent.Schema
}

func (InformationClassify) Mixin() []ent.Mixin {
	return []ent.Mixin{
		CommonMixin{},
	}
}

func (InformationClassify) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Default("").Comment("名称"),
		field.Uint8("status").Default(1).Comment("状态：1、锁定。2:生效"),
	}
}

func (InformationClassify) Edges() []ent.Edge {
	return []ent.Edge{}
}
