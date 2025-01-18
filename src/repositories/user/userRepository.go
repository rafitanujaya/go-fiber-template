package userRepository

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	Entity "github.com/rafitanujaya/go-fiber-template/src/model/entities/user"
	"github.com/samber/do/v2"
)

type UserRepository struct{}

func NewUserRepository() UserRepositoryInterface {
	return &UserRepository{}
}

func NewUserRepositoryInject(i do.Injector) (UserRepositoryInterface, error) {
	return NewUserRepository(), nil
}

func (ur *UserRepository) CreateUser(ctx context.Context, pool *pgxpool.Pool, user Entity.User) (userId string, err error) {
	query := `INSERT INTO users(email, password) VALUES($1, $2) RETURNING id`
	fmt.Printf("Email: %s, Password %s", user.Email, user.Password)

	row := pool.QueryRow(ctx, query, user.Email, user.Password)
	err = row.Scan(&userId)
	if err != nil {
		return "", err
	}
	return userId, nil

}
