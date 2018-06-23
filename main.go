package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/mnmtanish/go-graphiql"
	"github.com/unrolled/render"

	"github.com/ryanzola/GoLangGraphQL/schema"
)

// R renderer instace
var R *render.Render

// PORT number to be used
var PORT = "4000"

func init() {
	R = render.New(render.Options{
		Directory:     "views",
		IsDevelopment: true,
		Extensions:    []string{".html"},
	})
}

func serveGraphQL(query string, schema graphql.Schema) *graphql.Result {
	// helper to parse request query

	// execute graphql query
	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})
	if len(result.Errors) > 0 {
		fmt.Printf("wrong result, unexpected errors: %v", result.Errors)
	}

	return result
}

func main() {

	// GraphQL Endpoint
	http.HandleFunc("/graphql", func(w http.ResponseWriter, r *http.Request) {
		result := serveGraphQL(r.URL.Query().Get("query"), schema.schema)
		json.NewEncoder(w).Encode(result)
	})
	http.HandleFunc("/", graphiql.ServeGraphiQL)

	bind := fmt.Sprintf(":%s", PORT)

	log.Printf("Starting server on port %s", bind)

	log.Fatal(http.ListenAndServe(bind, nil))
}
