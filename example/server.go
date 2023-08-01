package main

import (
	"example/graph"
	"example/graph/model"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	embeddedpostgres "github.com/fergusstrange/embedded-postgres"
	customContext "github.com/juunini/gorm-custom-context"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const defaultPort = "8080"

func main() {
	runEmbeddedPostgres()
	go interruptEmbedded()

	db, err := gorm.Open(
		postgres.Open("host=127.0.0.1 user=postgres password=1234 dbname=postgres port=5432 sslmode=disable TimeZone=UTC"),
		&gorm.Config{},
	)
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&model.User{}, &model.Todo{})

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))
	ctx := customContext.CreateContext(&customContext.CustomContext{
		Database: db,
	}, srv)

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", ctx)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

var embeddedPostgres *embeddedpostgres.EmbeddedPostgres

func runEmbeddedPostgres() {
	embeddedPostgres = embeddedpostgres.NewDatabase(
		embeddedpostgres.DefaultConfig().
			Database("postgres").
			Username("postgres").
			Password("1234").
			Port(5432),
	)
	if err := embeddedPostgres.Start(); err != nil {
		panic(err)
	}
}

func interruptEmbedded() {
	sig := make(chan os.Signal, 1)
	signal.Notify(
		sig,
		syscall.SIGTERM,
		syscall.SIGINT,
		os.Interrupt,
	)

	<-sig

	embeddedPostgres.Stop()
	os.Exit(0)
}
