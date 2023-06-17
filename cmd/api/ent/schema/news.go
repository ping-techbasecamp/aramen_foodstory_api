package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// News holds the schema definition for the News entity.
type News struct {
	ent.Schema
}

// Fields of the News.
func (News) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.String("description"),
		field.String("tag"),
		field.String("place"),
		field.String("date"),
	}
}

// Edges of the News.
func (News) Edges() []ent.Edge {
	return nil
}
