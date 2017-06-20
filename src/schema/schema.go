package schema

import (
	"github.com/artshpakov/grapho/src/api"
	"github.com/go-pg/pg"
	"github.com/graphql-go/graphql"
	_ "github.com/lib/pq"
)

func DefineSchema(db *pg.DB) (graphql.Schema, error) {
	postType := PostType()
	userProfileType := UserProfileType()
	userType := UserType(db, &userProfileType, &postType)

	fields := make(map[string]*graphql.Field)
	for field, definition := range api.Users(db, userType) {
		fields[field] = definition
	}
	for field, definition := range api.Posts(db, postType) {
		fields[field] = definition
	}

	rootQuery := graphql.NewObject(graphql.ObjectConfig{
		Name:   "RootQuery",
		Fields: graphql.Fields(fields),
	})

	schemaConfig := graphql.SchemaConfig{Query: rootQuery}
	return graphql.NewSchema(schemaConfig)
}
