package graphql

import (
	"github.com/graphql-go/graphql"
	"github.com/sirupsen/logrus"
)

var numberHolderType = graphql.NewObject(graphql.ObjectConfig{
	Name: "NumberHolder",
	Fields: graphql.Fields{
		"theNumber": &graphql.Field{
			Type: graphql.Int,
		},
	},
})

var rootMutation = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootMutation",
	Fields: graphql.Fields{
		"createLicense": &graphql.Field{
			Type: numberHolderType,
			Args: graphql.FieldConfigArgument{
				"newNumber": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				logrus.Info("graphQL mutation executing...")
				return License{Name: "foobar"}, nil
			},
		},
	},
})
