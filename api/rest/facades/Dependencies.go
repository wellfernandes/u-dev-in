package facades

import (
	"api/rest/models"
	"context"
)

type userService interface {
	create(ctx context.Context, user models.User) (uint64, error)
}
