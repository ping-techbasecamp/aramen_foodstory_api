package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Branches holds the schema definition for the Branches entity.
type Branches struct {
	ent.Schema
}

// Fields of the Branches.
func (Branches) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.String("short_description"),
		field.String("full_description"),
		field.String("telephone"),
		field.Int("latitude"),
		field.Int("longitude"),
		field.String("goole_map_url"),
		field.Bool("dine_in"),
		field.Bool("delivery"),
		field.Bool("take_away"),
	}
}

// Edges of the Branches.
func (Branches) Edges() []ent.Edge {
	return nil
}
