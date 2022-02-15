package bug

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"entgo.io/bug/ent"
)

func (r *pageInfoResolver) StartCursor(ctx context.Context, obj *ent.PageInfo) (*string, error) {
	return MarshalCursor(obj.StartCursor), nil
}

func (r *pageInfoResolver) EndCursor(ctx context.Context, obj *ent.PageInfo) (*string, error) {
	return MarshalCursor(obj.EndCursor), nil
}

func (r *queryResolver) Node(ctx context.Context, id ent.GlobalID) (ent.Noder, error) {
	return r.client.Noder(ctx, id)
}

func (r *queryResolver) Nodes(ctx context.Context, ids []*ent.GlobalID) ([]ent.Noder, error) {
	return r.client.Noders(ctx, ids)
}

func (r *queryResolver) Users(ctx context.Context, after *string, first *int, before *string, last *int, orderBy *ent.UserOrder, where *ent.UserWhereInput) (*ent.UserConnection, error) {
	afterCur, err := UnmarshalCursor(after)
	if err != nil {
		return nil, err
	}
	beforeCur, err := UnmarshalCursor(before)
	if err != nil {
		return nil, err
	}

	return r.client.User.Query().Paginate(ctx, afterCur, first, beforeCur, last,
		ent.WithUserOrder(orderBy),
		ent.WithUserFilter(where.Filter),
	)
}

func (r *userEdgeResolver) Cursor(ctx context.Context, obj *ent.UserEdge) (string, error) {
	return *MarshalCursor(&obj.Cursor), nil
}

// PageInfo returns PageInfoResolver implementation.
func (r *Resolver) PageInfo() PageInfoResolver { return &pageInfoResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

// UserEdge returns UserEdgeResolver implementation.
func (r *Resolver) UserEdge() UserEdgeResolver { return &userEdgeResolver{r} }

type pageInfoResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type userEdgeResolver struct{ *Resolver }
