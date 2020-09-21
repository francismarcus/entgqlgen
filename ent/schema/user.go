package schema

// TODO: Ordering fields should normally be indexed to avoid full table DB scan. https://entgo.io/docs/schema-indexes/
import (
	"time"

	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
	"github.com/facebookincubator/ent-contrib/entgql"

	"github.com/badoux/checkmail"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
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
		field.String("username").
			Unique(),
		field.String("email").
			Unique().
			Validate(checkmail.ValidateFormat),
		field.String("password"),

		field.Int("follows_count").
			Default(0),
		field.Int("followers_count").
			Default(0),
		field.Int("shouts_count").
			Default(0),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("following", User.Type).
			From("followers"),
		edge.To("programs", Program.Type),
		edge.To("shouts", Shout.Type),
		edge.From("liked_shouts", Shout.Type).
			Ref("liked_by"),
		edge.To("settings", UserSettings.Type).
			Unique(),
	}
}
