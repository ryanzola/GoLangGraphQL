package schema

import (
	"github.com/graphql-go/graphql"
	"github.com/ryanzola/GoLangGraphQL/models"
)

var Root graphql.Schema
var BookType graphql.Object
var AuthorType graphql.Object

func init() {
	BookType = graphql.NewObject(graphql.ObjectConfig{
		Name:        "book",
		Description: "Book object",
		Fields: graphql.FieldsThunk(func() graphql.Fields {
			return graphql.Fields{
				"id": &graphql.Field{
					Type:        graphql.ID,
					Description: "Book Id"},
				"name": &graphql.Field{
					Type:        graphql.String,
					Description: "Book title"},
				"genre": &graphql.Field{
					Type:        graphql.String,
					Description: "Book genre"},
				"author": &graphql.Field{
					Type:        AuthorType,
					Description: "Author of the book",
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						if book, ok := p.Source.(*models.Book); ok {
							// Retrieve the author of the book
							return data.GetAuthor(author.ID)
						}
						return nil, nil
					},
				},
			}
		}),
	})

	AuthorType = graphql.NewObject(graphql.ObjectConfig{
		Name:        "author",
		Description: "Author object",
		Fields: graphql.FieldsThunk(func() graphql.Fields {
			return graphql.Fields{
				"id": &graphql.Field{Type: graphql.ID, Description: "Author Id"},
			}
		}),
	})
}
