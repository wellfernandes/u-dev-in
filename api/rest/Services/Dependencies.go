package Services

import (
	"api/rest/models"
	"context"
)

type userRepository interface {
	Create(ctx context.Context, user *models.User) (*models.User, error)
}
