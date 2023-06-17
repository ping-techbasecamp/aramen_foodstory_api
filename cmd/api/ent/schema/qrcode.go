package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Qrcode holds the schema definition for the Qrcode entity.
type Qrcode struct {
	ent.Schema
}

// Fields of the Qrcode.
func (Qrcode) Fields() []ent.Field {
	return []ent.Field{
		field.String("qrcode_id"),
		field.String("order_id"),
		field.String("order_status"),
		field.String("branch_name"),
		field.String("order_channel"),
		field.Int("total_paid"),
		field.Int("user_id"),
	}
}

// Edges of the Qrcode.
func (Qrcode) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("qrcode").Field("user_id").Required().Unique(),
	}

}
