package schema

import (
	"github.com/artshpakov/grapho/src/models"
	"github.com/go-pg/pg"
	"github.com/graphql-go/graphql"
)

func UserType(db *pg.DB, userProfileType *graphql.Type, postType *graphql.Type) graphql.Type {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "User",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.NewNonNull(graphql.Int),
			},

			"email": &graphql.Field{
				Type: graphql.String,
			},

			"slug": &graphql.Field{
				Type: graphql.String,
			},

			"name": &graphql.Field{
				Type: graphql.String,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if user, ok := p.Source.(models.User); ok {
						return user.Name(), nil
					}
					return nil, nil
				},
			},

			"profile": &graphql.Field{
				Type: *userProfileType,
			},

			"posts": &graphql.Field{
				Type: graphql.NewList(*postType),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if user, ok := p.Source.(models.User); ok {
						return user.Posts(db)
					}
					return nil, nil
				},
			},
		},
	})
}
