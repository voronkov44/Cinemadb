package actor

import "context"

type Repository interface {
	Create(ctx context.Context, actor *Actor) error
	FindAll(ctx context.Context) (u []Actor, err error)
	FindOne(ctx context.Context, id string) (Actor, error)
	Update(ctx context.Context, user Actor) error
	Delete(ctx context.Context, id string) error
}
