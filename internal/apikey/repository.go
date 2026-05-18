package apikey

import "context"

type Repository interface {
	Create(ctx context.Context, email, keyHash string) (APIKey, error)
	FindByHash(ctx context.Context, keyHash string) (*APIKey, error)
	RevokeAll(ctx context.Context, email string) error
}
