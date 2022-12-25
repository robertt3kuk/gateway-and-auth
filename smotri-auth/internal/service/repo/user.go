package repo

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"

	"smotri-auth/internal/entity"
	"smotri-auth/pkg/postgres"
)

const _defaultUserEntityCap = 64

// UserPostgres -.
type UserPostgres struct {
	*postgres.Postgres
}

// New -.
func NewUser(pg *postgres.Postgres) *UserPostgres {
	return &UserPostgres{pg}
}

// squirrel builder
func (r *UserPostgres) GetUser(ctx context.Context, email, password string) (entity.User, error) {
	var result entity.User

	sql, args, err := r.Builder.
		Select("id, email, password").
		From("users").
		Where(sq.Eq{"email": email, "password": password}).
		ToSql()
	if err != nil {
		return result, fmt.Errorf("UserPostgres - GetUser - r.Builder: %w", err)
	}

	row := r.Pool.QueryRow(ctx, sql, args...)

	err = row.Scan(
		&result.ID,
		&result.Email,
		&result.Password,
	)
	if err != nil {
		return result, fmt.Errorf("UserPostgres - GetUser - row.Scan: %w", err)
	}

	return result, nil
}

// CreateUser -.
func (r *UserPostgres) CreateUser(ctx context.Context, user entity.User) (int, error) {
	var result int

	sql, args, err := r.Builder.
		Insert("users").
		Columns("email", "password").
		Values(user.Email, user.Password).
		Suffix("RETURNING \"id\"").
		ToSql()
	if err != nil {
		return result, fmt.Errorf("UserPostgres - CreateUser - r.Builder: %w", err)
	}

	row := r.Pool.QueryRow(ctx, sql, args...)

	err = row.Scan(&result)
	if err != nil {
		return result, fmt.Errorf("UserPostgres - CreateUser - row.Scan: %w", err)
	}

	return result, nil
}
