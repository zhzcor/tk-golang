package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"time"
)

// -------------------------------------------------
// Mixin definition

type CommonMixin struct {
	// We embed the `mixin.Schema` to avoid
	// implementing the rest of the methods.
	mixin.Schema
}

func (CommonMixin) Fields() []ent.Field {
	return []ent.Field{
		field.String("uuid"),
		field.Time("created_at").
			Immutable().
			Optional().
			Nillable().
			Default(time.Now),
		field.Time("updated_at").
			Default(time.Now).
			Nillable().
			Optional().
			UpdateDefault(time.Now),
		field.Time("deleted_at").
			Nillable().
			Optional(),
	}
}

type OrderMixin struct {
	mixin.Schema
}

func (m OrderMixin) Fields() []ent.Field {
	return []ent.Field{
		field.String("order_id"),
		field.Int("version").Default(0),
	}
}

type PasswordMixin struct {
	mixin.Schema
}

func (PasswordMixin) Fields() []ent.Field {
	return []ent.Field{
		field.String("password"),
		field.String("salt"),
	}
}
