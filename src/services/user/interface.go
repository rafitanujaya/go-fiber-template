package userService

import (
	"context"

	"github.com/rafitanujaya/go-fiber-template/src/model/dtos/request"
	"github.com/rafitanujaya/go-fiber-template/src/model/dtos/response"
)

type UserServiceInterface interface {
	Register(ctx context.Context, input request.UserRegister) (response.UserRegister, error)
}
