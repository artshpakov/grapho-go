package api

import (
	"github.com/artshpakov/grapho/src/models"
	"github.com/go-pg/pg"
	"github.com/graphql-go/graphql"
)

const DEFAULT_POSTS_LIMIT = 30

func Posts(db *pg.DB, postType graphql.Type) map[string]*graphql.Field {
	fields := make(map[string]*graphql.Field)

	fields["posts"] = &graphql.Field{
		Type:        graphql.NewList(postType),
		Description: "List of posts",
		Args: graphql.FieldConfigArgument{
			"limit": &graphql.ArgumentConfig{
				Type:         graphql.Int,
				DefaultValue: DEFAULT_POSTS_LIMIT,
			},
			"id": &graphql.ArgumentConfig{
				Type:         graphql.NewList(graphql.Int),
				DefaultValue: make([]interface{}, 0),
			},
			"user_id": &graphql.ArgumentConfig{
				Type:         graphql.Int,
				DefaultValue: nil,
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			return models.FetchPosts(db, p.Args["limit"].(int), p.Args["id"].([]interface{}), p.Args["user_id"])
		},
	}

	return fields
}
