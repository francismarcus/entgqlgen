package graph

import (
	"github.com/99designs/gqlgen/graphql"
	"github.com/francismarcus/eg/ent"
	"github.com/francismarcus/eg/graph/generated"
	"github.com/francismarcus/eg/pkg/directives"
)

// Resolver is the resolver root.
type Resolver struct{ client *ent.Client }

// NewSchema creates a graphql executable schema.
func NewSchema(client *ent.Client) graphql.ExecutableSchema {

	c := generated.Config{
		Resolvers: &Resolver{client},
	}

	c.Directives = generated.DirectiveRoot{
		Log:   directives.Log,
		HasID: directives.HasID,
	}

	return generated.NewExecutableSchema(c)
}
