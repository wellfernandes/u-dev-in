package controllers

import (
	"api/dtos"
	"context"
)

type userFacade interface {
	Create(ctx context.Context, userRequest *dtos.CreateUserRequest) (*dtos.CreateUserResponse, error)
}
