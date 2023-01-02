package facades

import (
	"api/rest/models"
	"context"
)

type userService interface {
	Create(ctx context.Context, user models.User) (*models.User, error)
}
