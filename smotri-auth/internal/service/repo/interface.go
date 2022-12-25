package repo

import (
	"context"

	"smotri-auth/internal/entity"
	"smotri-auth/pkg/postgres"
)

type Repo struct {
	User
}

func NewRepo(pg *postgres.Postgres) *Repo {
	return &Repo{
		User: NewUser(pg),
	}
}

type User interface {
	GetUser(ctx context.Context, email, password string) (entity.User, error)
	CreateUser(ctx context.Context, user entity.User) (int, error)
}
