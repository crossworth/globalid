package bug

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"entgo.io/bug/ent"
	"entgo.io/bug/ent/user"
	"github.com/google/uuid"
)

func (r *entityResolver) FindUserByID(ctx context.Context, id ent.GlobalID) (*ent.User, error) {
	return r.client.User.Query().Where(user.ID(uuid.MustParse(id.ID))).Only(ctx)
}

// Entity returns EntityResolver implementation.
func (r *Resolver) Entity() EntityResolver { return &entityResolver{r} }

type entityResolver struct{ *Resolver }
