package schema

import (
	"time"

	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
	"github.com/facebookincubator/ent-contrib/entgql"
)

// Diet holds the schema definition for the Diet entity.
type Diet struct {
	ent.Schema
}

// Fields of the Diet.
func (Diet) Fields() []ent.Field {
	return []ent.Field{
		field.Time("created_at").
			Default(time.Now).
			Immutable().
			Annotations(
				entgql.OrderField("CREATED_AT"),
			),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now).
			Annotations(
				entgql.OrderField("UPDATED_AT"),
			),
		field.String("name"),
		field.Int("goal_weight"),
		field.Int("length"),
	}
}

// Edges of the Diet.
func (Diet) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("author", User.Type).
			Ref("diets").
			Unique().
			// We add the "Required" method to the builder
			// to make this edge required on entity creation.
			// i.e. Settings cannot be created without its User.
			Required(),
	}
}
