package repositories

import (
	"api/rest/models"
	"context"
	"gorm.io/gorm"
)

type UserRepository struct {
	*BaseRepository
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	baseRepo := NewBaseRepository(db)
	return &UserRepository{
		baseRepo,
	}
}

func (u *UserRepository) Create(ctx context.Context, user *models.User) (*models.User, error) {
	conn, err := u.getConnection(ctx)
	if err != nil {
		return nil, err
	}

	err = conn.Create(user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}
