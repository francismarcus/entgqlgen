package schema

import (
	"time"

	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
	"github.com/facebookincubator/ent-contrib/entgql"
)

// Program holds the schema definition for the Program entity.
type Program struct {
	ent.Schema
}

// Fields of the Program.
func (Program) Fields() []ent.Field {
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
	}
}

// Edges of the Program.
func (Program) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("author", User.Type).
			Ref("programs").
			Unique(),
		edge.To("workouts", Workout.Type),
	}
}
