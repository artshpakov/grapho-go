package schema

import "github.com/graphql-go/graphql"

func UserProfileType() graphql.Type {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "UserProfile",
		Fields: graphql.Fields{
			"firstname": &graphql.Field{
				Type: graphql.String,
			},

			"lastname": &graphql.Field{
				Type: graphql.String,
			},

			"location": &graphql.Field{
				Type: graphql.String,
			},

			"bio": &graphql.Field{
				Type: graphql.String,
			},
		},
	})
}
