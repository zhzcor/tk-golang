package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

//热门搜索
type HotSearch struct {
	ent.Schema
}

func (HotSearch) Mixin() []ent.Mixin {
	return []ent.Mixin{
		CommonMixin{},
	}
}

func (HotSearch) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Default("").Comment("名称"),
		field.Uint8("status").Default(1).Comment("状态：1、锁定。2:生效"),
		field.Int("search_count").Default(0).Comment("搜索次数"),
	}
}

func (HotSearch) Edges() []ent.Edge {
	return []ent.Edge{}
}
