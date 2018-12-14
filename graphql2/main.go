package main

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/graphql-go/graphql"
)

func Filter(pirates []Pirate, f func(Pirate) bool) []Pirate {
	vsf := make([]Pirate, 0)
	for _, v := range pirates {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}

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
				Args: graphql.FieldConfigArgument{
					"country": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					// Optional argument country. `country, _`does the trick.
					country, _ := params.Args["country"].(string)
					filtered := Filter(pirates, func(v Pirate) bool {
						return strings.Contains(v.Country, country)
					})
					return filtered, nil
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
