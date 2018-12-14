package main

import (
	"encoding/json"
	"net/http"

	"github.com/graphql-go/graphql"
)

func main() {

	// GraphQL model
	pirateType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Pirate",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.String,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"life": &graphql.Field{
				Type: graphql.String,
			},
			"yearsactive": &graphql.Field{
				Type: graphql.String,
			},
			"country": &graphql.Field{
				Type: graphql.String,
			},
			"comments": &graphql.Field{
				Type: graphql.String,
			},
			"wikipedia": &graphql.Field{
				Type: graphql.String,
			},
		},
	})

	rootQuery := graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"pirates": &graphql.Field{
				Type: graphql.NewList(pirateType),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return pirates, nil
				},
			},
		},
	})

	schema, _ := graphql.NewSchema(graphql.SchemaConfig{
		Query: rootQuery,
	})
	http.HandleFunc("/graphql", func(w http.ResponseWriter, r *http.Request) {
		result := graphql.Do(graphql.Params{
			Schema:        schema,
			RequestString: r.URL.Query().Get("query"),
		})
		json.NewEncoder(w).Encode(result)
	})
	http.ListenAndServe(":3000", nil)
}
