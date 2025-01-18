package userRepository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	Entity "github.com/rafitanujaya/go-fiber-template/src/model/entities/user"
)

type UserRepositoryInterface interface {
	CreateUser(ctx context.Context, pool *pgxpool.Pool, user Entity.User) (userId string, err error)
}
