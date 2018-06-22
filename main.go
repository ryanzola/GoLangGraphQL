package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/BurntSushi/toml"
)

type Book struct {
	Name 			string `json:"text"`
	Genre 		string `json:"text"`
	AuthorId 	string `json:"text"`
}

type Author struct {
	Name 			string `json:"text"`
	Age 			int    `json:"number"`
	Books 		[]Book `json:books`
}


func main() {

	http.HandleFunc("/graphql", func(w http.ResponseWriter, r *http.Request) {
		
		
		})
		json.NewEncoder(w).Encode(result)
	})


	bookType := &graphql.NewObject(
		graphql.ObjectConfig{
			Name: "book",
			Fields: graphql.Fields{
				"id": &graphql.Field{ Type: graphql.ID },
				"name": &graphql.Field{ Type: graphql.String },
				"genre": &graphql.Field{ Type: graphql.String },
				"author": &graphql.Field{ 
					Type: authorType,
					Resolve: func(p) (interface{}, error) {
						return Author.findById(p.AuthorId), nil
					},
				},
			},
	})

	authorType := &graphql.NewObject(graphql.ObjectConfig{
		Name: "author",
		Fields: graphql.Fields{
			"id": &graphql.Field{ Type: graphql.ID },
			"name": &graphql.Field{ Type: graphql.String },
			"age": &graphql.Field{ Type: graphql.Int },
			"books": &graphql.Field{ 
				Type: graphql.List(bookType),
				Resolve: func(p)
			}
		}
	})

	schemaConfig := graphql.SchemaConfig{Query: graphql.NewObject(rootQuery)}
	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		log.Fatalf("Failed to create new schema, error %v", err)
	}

	rJSON, _ := json.Marshal(r)
	fmt.Printf("%s \n", rJSON)

	http.ListenAndServe(":4000", nil)
}
