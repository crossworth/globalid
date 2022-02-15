package bug

import (
	"bytes"

	"entgo.io/bug/ent"
	"github.com/99designs/gqlgen/graphql"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	client *ent.Client
}

func NewSchema(client *ent.Client) graphql.ExecutableSchema {
	return NewExecutableSchema(Config{
		Resolvers: &Resolver{client},
	})
}

func MarshalCursor(cursor *ent.Cursor) *string {
	var buf bytes.Buffer
	cursor.MarshalGQL(&buf)
	cur := buf.String()
	return &cur
}

func UnmarshalCursor(cursor *string) (*ent.Cursor, error) {
	if cursor == nil {
		return nil, nil
	}
	var cur ent.Cursor
	err := cur.UnmarshalGQL(*cursor)
	if err != nil {
		return nil, err
	}
	return &cur, nil
}
