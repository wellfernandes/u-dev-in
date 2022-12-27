package dtos

import "api/rest/models"

type CreateUserResponse struct {
	ID uint64 `json:"id"`
}

func (r *CreateUserResponse) ParseFromUserObject(user *models.User) {
	r.ID = user.ID
}
