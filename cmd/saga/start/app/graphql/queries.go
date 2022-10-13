package graphql

import "github.com/graphql-go/graphql"

var rootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootQuery",
	Fields: graphql.Fields{
		"license": &graphql.Field{
			Type:        licenseType,
			Description: "Get a single license",
			Args: graphql.FieldConfigArgument{
				"name": &graphql.ArgumentConfig{
					Type:        graphql.String,
					Description: `Get license by name`,
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				return License{"foobar"}, nil
			},
		},
	},
})
