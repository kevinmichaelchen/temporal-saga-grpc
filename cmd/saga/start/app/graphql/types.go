package graphql

import "github.com/graphql-go/graphql"

type License struct {
	Name string `json:"name"`
}

var licenseType = graphql.NewObject(graphql.ObjectConfig{
	Name: "License",
	Fields: graphql.Fields{
		"name": &graphql.Field{
			Type:        graphql.String,
			Description: `The name of the License`,
		},
	},
})

var orgType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Org",
	Fields: graphql.Fields{
		"name": &graphql.Field{
			Type:        graphql.String,
			Description: `The name of the Org`,
		},
	},
})

var profileType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Profile",
	Fields: graphql.Fields{
		"name": &graphql.Field{
			Type:        graphql.String,
			Description: `The name of the Profile`,
		},
	},
})
