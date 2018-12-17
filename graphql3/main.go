package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/fatih/structs"
	"github.com/graphql-go/graphql"
)

func Filter(data []Pirate, args map[string]interface{}) []Pirate {
	vsf := make([]Pirate, 0)
	if len(args) != 0 {

		for _, p := range data {
			fmt.Println(p)
			for k, v := range args {
				structMap := structs.Map(p)
				if structMap[strings.Title(k)] != v {
					break
				}
				vsf = append(vsf, p)
			}

		}
		return vsf
	}
	return data
}

func main() {

	// GraphQL model
	pirateType := graphql.NewObject(
		graphql.ObjectConfig{
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

	rootQuery := graphql.NewObject(

		graphql.ObjectConfig{
			Name: "Query",
			Fields: graphql.Fields{
				"pirates": &graphql.Field{
					Type: graphql.NewList(pirateType),
					Args: graphql.FieldConfigArgument{
						"id": &graphql.ArgumentConfig{
							Type: graphql.String,
						},
						"name": &graphql.ArgumentConfig{
							Type: graphql.String,
						},
						"life": &graphql.ArgumentConfig{
							Type: graphql.String,
						},
						"yearsactive": &graphql.ArgumentConfig{
							Type: graphql.String,
						},
						"country": &graphql.ArgumentConfig{
							Type: graphql.String,
						},
						"comments": &graphql.ArgumentConfig{
							Type: graphql.String,
						},
					},
					Resolve: func(params graphql.ResolveParams) (interface{}, error) {
						return Filter(pirates, params.Args), nil
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
