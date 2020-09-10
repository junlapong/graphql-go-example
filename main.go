package main

import (
	"fmt"
	"log"
	"net/http"

	// "github.com/graphql-go/relay/examples/starwars"
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

	http.Handle("/", http.FileServer(http.Dir("static")))
	http.Handle("/graphql", h)

	port := "3000"
	fmt.Printf("Open http://localhost:%s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
