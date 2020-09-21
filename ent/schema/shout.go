package schema

import (
	"time"

	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
	"github.com/facebookincubator/ent-contrib/entgql"
)

// Shout holds the schema definition for the Shout entity.
type Shout struct {
	ent.Schema
}

// Fields of the Shout.
func (Shout) Fields() []ent.Field {
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
		field.String("message"),
		field.Int("likes").
			Default(0).
			Annotations(
				entgql.OrderField("LIKES"),
			),
	}
}

// Edges of the Shout.
func (Shout) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("author", User.Type).
			Ref("shouts").
			Unique(),
		edge.To("liked_by", User.Type),
	}
}
