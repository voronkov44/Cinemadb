package film

import (
	"context"
)

type Repository interface {
	FindAll(ctx context.Context) (u []Film, err error)
}
