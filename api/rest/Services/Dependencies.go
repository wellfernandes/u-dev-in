package Services

import (
	"api/rest/models"
	"context"
)

type userRepository interface {
	create(ctx context.Context, user models.User) (uint64, error)
}
