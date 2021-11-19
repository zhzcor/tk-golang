package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

//app协议表
type AppAgreement struct {
	ent.Schema
}

func (AppAgreement) Mixin() []ent.Mixin {
	return []ent.Mixin{
		CommonMixin{},
	}
}

// Fields of the User.
func (AppAgreement) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Default("").Comment("协议名称"),
		field.Uint8("type").Default(0).Comment(
			"协议类型：1：服务协议，2：隐私政策，3：ios充值协议，4：账户注销协议，5：App温馨提示，" +
				"6：优惠卷使用规则，7：关于我们，8：加入我们，9：版权声明，10：联系我们，11：常见问题"),
		field.Text("detail").Optional().Comment("协议详情"),
	}
}

// Edges of the User.
func (AppAgreement) Edges() []ent.Edge {
	return []ent.Edge{}
}
