package main

import (
	"fmt"
	"net/http"

	"github.com/graphql-go/graphql/testutil"
	"github.com/graphql-go/handler"
)

func main() {

	h := handler.New(&handler.Config{
		// Schema: &starwars.Schema,
		Schema:     &testutil.StarWarsSchema,
		Pretty:     true,
		GraphiQL:   true,
		Playground: true,
	})

	fs := http.FileServer(http.Dir("static"))

	http.Handle("/", fs)
	http.Handle("/graphql", h)

	port := 3000
	fmt.Printf("Server is running on port %d\n", port)
	fmt.Printf("Open http://localhost:%d\n", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}

// "github.com/graphql-go/relay/examples/starwars"
// "github.com/graphql-go/graphql/testutil"
