package schema

import (
	"time"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Annotations(entgql.GlobalID()),
		field.Int("age"),
		field.String("name"),
		field.Time("created_at").
			Default(time.Now).
			Immutable().
			Annotations(entgql.OrderField("CREATED_AT")),
	}
}
