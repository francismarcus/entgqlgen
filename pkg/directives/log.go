package directives

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	"go.uber.org/zap"
)

type key string

const resolverCtx key = "resolver_context"

// Log logs to the console
func Log(ctx context.Context, obj interface{}, next graphql.Resolver) (interface{}, error) {
	sugar := zap.NewExample().Sugar()
	// r := ctx.Value(chi.RouteCtxKey)
	g := graphql.GetFieldContext(ctx)
	defer sugar.Sync()
	// user := middlewares.UserContext(ctx)
	sugar.Infow("@log %+v\n",
		"object", g.Object,
		"Arguments", g.Args,
	)

	return next(ctx)
}
