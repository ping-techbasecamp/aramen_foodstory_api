package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// UserOtp holds the schema definition for the UserOtp entity.
type UserOtp struct {
	ent.Schema
}

// Fields of the UserOtp.
func (UserOtp) Fields() []ent.Field {
	return []ent.Field{
		field.Int("user_id"),
		field.String("reference").Unique(),
		field.String("hashed_otp"),
	}
}

// Edges of the UserOtp.
func (UserOtp) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("otp").Field("user_id").Required().Unique(),
	}
}
