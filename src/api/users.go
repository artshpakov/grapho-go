package api

import (
	"github.com/artshpakov/grapho/src/models"
	"github.com/go-pg/pg"
	"github.com/graphql-go/graphql"
)

func Users(db *pg.DB, userType graphql.Type) map[string]*graphql.Field {
	fields := make(map[string]*graphql.Field)

	fields["users"] = &graphql.Field{
		Type:        graphql.NewList(userType),
		Description: "List of users",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type:         graphql.NewList(graphql.Int),
				DefaultValue: make([]interface{}, 0),
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			return models.UsersFetch(db, p.Args["id"].([]interface{}))
		},
	}

	return fields
}
