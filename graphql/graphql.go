package graphapi

import (
	"github.com/graphql-go/graphql"
)

var queryType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Query",
	Fields: graphql.Fields{
		"test": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return "Hello World!", nil
			},
		},
	},
})

// Schema - outputted gql schema
var Schema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query: queryType,
})
