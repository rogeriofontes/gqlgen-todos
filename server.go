package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gorilla/handlers"
	"github.com/rogeriofontes/gqlgen-todos/graph"
)

const defaultPort = "8080"

func main() {
	// Definir a porta
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	// Criar o servidor GraphQL
	srv := handler.New(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

	// Configurar os transportes suportados
	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})

	// Configurar cache de queries
	srv.Use(extension.Introspection{})
	srv.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New[string](100),
	})

	// Middleware COtir requisições de qualquer origem
	corsMiddleware := handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}), // Pode restringir para ["http://seu-frontend.com"]
		handlers.AllowedMethods([]string{"GET", "POST", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
	)

	// Configurar handlers
	mux := http.NewServeMux()
	mux.Handle("/", playground.Handler("GraphQL Playground", "/query"))
	mux.Handle("/query", srv)

	// Aplicar middleware CORS a todas as rotas
	handlerWithCORS := corsMiddleware(mux)

	// Iniciar o servidor
	log.Printf("🚀 Servidor rodando em http://localhost:%s/", port)
	log.Fatal(http.ListenAndServe(":"+port, handlerWithCORS))
}
