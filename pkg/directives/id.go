package directives

import (
	"context"
	"fmt"

	"github.com/99designs/gqlgen/graphql"
	"github.com/francismarcus/eg/pkg/middlewares"
	"go.uber.org/zap"
)

// HasID directive
func HasID(ctx context.Context, obj interface{}, next graphql.Resolver) (interface{}, error) {
	sugar := zap.NewExample().Sugar()
	defer sugar.Sync()
	u := middlewares.UserContext(ctx)
	g := graphql.GetFieldContext(ctx)
	id := g.Args["id"]

	if id != u.ID {
		return nil, fmt.Errorf("You do not have permission to do that")
	}

	sugar.Infow("@hasID %+v\n",
		"id", id,
		"uID", u.ID,
	)

	return next(ctx)
}
