package schema

import "github.com/graphql-go/graphql"

func PostType() graphql.Type {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "Post",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.NewNonNull(graphql.Int),
			},

			"title": &graphql.Field{
				Type: graphql.String,
			},

			"text": &graphql.Field{
				Type: graphql.String,
			},

			"slug": &graphql.Field{
				Type: graphql.String,
			},

			"user_id": &graphql.Field{
				Type: graphql.Int,
			},

			"created_at": &graphql.Field{
				Type: graphql.String,
			},
		},
	})
}
