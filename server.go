package main

import (
	"context"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/debug"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/alecthomas/kong"
	"github.com/facebookincubator/ent-contrib/entgql"
	"github.com/francismarcus/eg/ent"
	"github.com/francismarcus/eg/ent/migrate"
	"github.com/francismarcus/eg/graph"
	"github.com/francismarcus/eg/pkg/cache"
	"github.com/francismarcus/eg/pkg/middlewares"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"go.uber.org/zap"

	_ "github.com/francismarcus/eg/ent/runtime"
	_ "github.com/lib/pq"
)

const defaultPort = "8080"

func main() {
	var cli struct {
		Addr  string `name:"address" default:":8080" help:"Address to listen on."`
		Debug bool   `name:"debug" help:"Enable debugging mode."`
	}

	kong.Parse(&cli)

	log, _ := zap.NewDevelopment()

	client, err := ent.Open(
		"postgres",
		"postgres://marcusmagnusson:password@pg:5432/getmyprogram?sslmode=disable",
	)

	if err != nil {
		log.Fatal("opening ent client", zap.Error(err))
	}

	if err := client.Schema.Create(
		context.Background(),
		migrate.WithGlobalUniqueID(true),
	); err != nil {
		log.Fatal("running schema migration", zap.Error(err))
	}

	redis, redisError := cache.NewCache("redis:6379")
	if redisError != nil {
		log.Fatal("cannot create APQ redis cache:", zap.Error(redisError))
	}

	router := chi.NewRouter()
	router.Use(middlewares.AuthMiddleware())
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	srv := handler.NewDefaultServer(graph.NewSchema(client))
	srv.Use(entgql.Transactioner{TxOpener: client})

	srv.SetErrorPresenter(entgql.DefaultErrorPresenter)
	if cli.Debug {
		srv.Use(&debug.Tracer{})
	}
	srv.AddTransport(transport.POST{})
	srv.Use(extension.AutomaticPersistedQuery{Cache: redis})

	router.Handle("/",
		playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)

	log.Info("listening on", zap.String("address", cli.Addr))
	if err := http.ListenAndServe(cli.Addr, router); err != nil {
		log.Error("http server terminated", zap.Error(err))
	}
}
