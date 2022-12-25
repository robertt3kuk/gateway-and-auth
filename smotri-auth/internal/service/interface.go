package service

import (
	"context"

	"smotri-auth/internal/entity"
	"smotri-auth/internal/service/repo"
)

type UseCase struct {
	User
}

func NewUseCase(r *repo.Repo) *UseCase {
	return &UseCase{
		User: NewUser(r.User),
	}

}

type User interface {
	GenerateToken(ctx context.Context, email, password string) (string, error)
	ParseToken(accessToken string) (int, error)
	CreateUser(ctx context.Context, user entity.User) (int, error)
}
