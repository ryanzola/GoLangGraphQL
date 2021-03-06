package schema

import (
	"github.com/graphql-go/graphql"
	"github.com/ryanzola/GoLangGraphQL/server/data"
	"github.com/ryanzola/GoLangGraphQL/server/models"
)

// Schema for all query objects
var Schema graphql.Schema

// BookType graphQL object
var BookType *graphql.Object

// AuthorType graphQL object
var AuthorType *graphql.Object

func init() {
	BookType = graphql.NewObject(graphql.ObjectConfig{
		Name:        "Book",
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
							return data.GetAuthor(book.AuthorID)
						}
						return nil, nil
					},
				},
			}
		}),
	})

	AuthorType = graphql.NewObject(graphql.ObjectConfig{
		Name:        "Author",
		Description: "Author object",
		Fields: graphql.FieldsThunk(func() graphql.Fields {
			return graphql.Fields{
				"id": &graphql.Field{
					Type:        graphql.ID,
					Description: "Author Id"},
				"name": &graphql.Field{
					Type:        graphql.String,
					Description: "Author name"},
				"age": &graphql.Field{
					Type:        graphql.Int,
					Description: "Author age"},
				"books": &graphql.Field{
					Type:        graphql.NewList(BookType),
					Description: "All books by author",
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						if author, ok := p.Source.(*models.Author); ok {
							return data.GetAllBooksByAuthor(author.ID)
						}
						return nil, nil
					},
				},
			}
		}),
	})

	var rootQuery = graphql.NewObject(graphql.ObjectConfig{
		Name:        "RootQuery",
		Description: "Root for all query objects on the GraphQL server",
		Fields: graphql.Fields{
			"books": &graphql.Field{
				Type:        graphql.NewList(BookType),
				Description: "Get a list of all books",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return data.GetBooks()
				},
			},
			"book": &graphql.Field{
				Type:        BookType,
				Description: "Get a single book",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					id := ""
					if v, ok := p.Args["id"].(string); ok {
						id = v
					}
					return data.GetBook(id)
				},
			},
			"authors": &graphql.Field{
				Type:        graphql.NewList(AuthorType),
				Description: "Get a list of all authors",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return data.GetAuthors()
				},
			},
			"author": &graphql.Field{
				Type:        AuthorType,
				Description: "Get a single author",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					id := ""
					if v, ok := p.Args["id"].(string); ok {
						id = v
					}
					return data.GetAuthor(id)
				},
			},
		},
	})

	var rootMutation = graphql.NewObject(graphql.ObjectConfig{
		Name: "RootMutation",
		Fields: graphql.Fields{
			"addAuthor": &graphql.Field{
				Type:        AuthorType,
				Description: "Add a new author",
			},
		},
	})

	var err error
	Schema, err = graphql.NewSchema(graphql.SchemaConfig{
		Query:    rootQuery,
		Mutation: rootMutation,
	})
	if err != nil {
		panic(err)
	}
}
