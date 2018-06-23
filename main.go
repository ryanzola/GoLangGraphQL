package main

import (
	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"github.com/unrolled/render"
	"golang.org/x/net/context"

	"github.com/ryanzola/GoLangGraphQL/schema"
)

var R *render.Render

var PORT = "4000"

func init() {
	R = render.New(render.Options{
		Directory:     "views",
		IsDevelopment: true,
		Extensions:    []string{".html"},
	})
}

func serveGraphQL(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	// helper to parse request query
	opts := handler.NewRequestOptions(r)

	//execute graphQL query
	params := graphql.Params{
		Schema: schema.Root,
		RequestString: opts.Query,
		VariableValues: opts.Variables,
		OperationName: opts.OperationName,
		Context: ctx
	}
}

func main() {

}
