package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
	"github.com/facebookincubator/ent-contrib/entgql"
)

// UserSettings holds the schema definition for the UserSettings entity.
type UserSettings struct {
	ent.Schema
}

// Fields of the UserSettings.
func (UserSettings) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("gender").
			NamedValues(
				"Male", "MALE",
				"Female", "FEMALE",
				"None", "NONE",
			).
			Annotations(
				entgql.OrderField("GENDER"),
			).Optional().
			Nillable(),
		field.Int("age").Optional().
			Nillable(),
		field.Int("weight").Optional().
			Nillable(),
		field.Int("height").Optional().
			Nillable(),
		field.Enum("level").
			NamedValues(
				"Beginner", "BEGINNER",
				"Intermediate", "INTERMEDIATE",
				"Advanced", "Advanced",
			).
			Annotations(
				entgql.OrderField("LEVEL"),
			),
	}
}

// Edges of the UserSettings.
func (UserSettings) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("belongs_to", User.Type).
			Ref("settings").
			Unique().
			// We add the "Required" method to the builder
			// to make this edge required on entity creation.
			// i.e. Settings cannot be created without its User.
			Required(),
	}
}
