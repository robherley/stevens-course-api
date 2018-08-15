package graphapi

import (
	"github.com/graphql-go/graphql"
)

// Schema - outputted gql schema
var Schema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query: queryType,
})
