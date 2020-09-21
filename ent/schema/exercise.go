package schema

import "github.com/facebook/ent"

// Exercise holds the schema definition for the Exercise entity.
type Exercise struct {
	ent.Schema
}

// Fields of the Exercise.
func (Exercise) Fields() []ent.Field {
	return nil
}

// Edges of the Exercise.
func (Exercise) Edges() []ent.Edge {
	return nil
}
