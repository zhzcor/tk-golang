package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

//导入任务
type ImportTask struct {
	ent.Schema
}

func (ImportTask) Mixin() []ent.Mixin {
	return []ent.Mixin{
		CommonMixin{},
	}
}

func (ImportTask) Fields() []ent.Field {
	return []ent.Field{
		field.String("import_name").Default("").Comment("名称"),
		field.Uint8("status").Default(1).Comment("状态 1：未开始，2：导入中，3：完成"),
		field.String("path").Default("").Comment("路径"),
		field.Int("total").Default(0).Comment("总数"),
		field.Int("created_admin_id").Default(0).Comment("创建人id"),
		field.String("remark").Default("").Comment("备注"),
	}
}

func (ImportTask) Edges() []ent.Edge {
	return []ent.Edge{}
}
