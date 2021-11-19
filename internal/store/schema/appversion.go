package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

//app版本表
type AppVersion struct {
	ent.Schema
}

func (AppVersion) Mixin() []ent.Mixin {
	return []ent.Mixin{
		CommonMixin{},
	}
}

func (AppVersion) Fields() []ent.Field {
	return []ent.Field{
		field.Uint16("ip").Default(0).Comment("登录ip"),
		field.String("name").Default("").Comment("版本名称"),
		field.String("sn").Default("").Comment("版本编号"),
		field.Text("remark").Optional().Comment("版本描述"),
		field.String("url").Default("").Comment("下载地址"),
		field.Uint8("phone_type").Default(0).Comment("适用机型 1:android 2:ios"),
		field.Uint8("is_force_update").Default(2).Comment("是否强制更新 1:是,2:否"),
	}
}

func (AppVersion) Edges() []ent.Edge {
	return []ent.Edge{}
}
