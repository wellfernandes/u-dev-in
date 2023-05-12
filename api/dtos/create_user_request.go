package dtos

import "api/rest/models"

type CreateUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (r *CreateUserRequest) ParseToUserObject() *models.User {
	return &models.User{
		Email:    r.Email,
		Password: r.Password,
	}
}
