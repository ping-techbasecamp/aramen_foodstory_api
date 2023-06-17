package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Member holds the schema definition for the Member entity.
type User struct {
	ent.Schema
}

// Fields of the Member.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("first_name"),
		field.String("last_name"),
		field.String("phone_number"),
		field.Time("birthdate"),
		field.String("email"),
		field.Time("created_at"),
		field.Int("point"),
		field.Int("star"),
	}
}

// Edges of the Member.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("otp", UserOtp.Type),
		edge.To("qrcode", Qrcode.Type),
	}
}
