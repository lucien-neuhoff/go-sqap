package repositories

import (
	"context"
	"database/sql"
	"errors"
	"go-sqap/internal/models"
	"go-sqap/internal/utils"
	"time"

	"github.com/google/uuid"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *models.User) error
	GetUserById(ctx context.Context, id uuid.UUID) (*models.User, error)
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
	UpdateUser(ctx context.Context, user *models.User) error
	DeleteUser(ctx context.Context, user *models.User) error
}

type userRepository struct {
	db     *sql.DB
	logger *utils.Logger
}

func NewUserRepository(db *sql.DB, logger *utils.Logger) UserRepository {
	return &userRepository{
		db:     db,
		logger: logger,
	}
}

func (r *userRepository) CreateUser(ctx context.Context, user *models.User) error {
	now := time.Now().UTC()
	user.CreatedAt.Time = now
	user.UpdatedAt.Time = now

	query := "INSERT INTO user (uuid, email, password) VALUES (?, ?, ?)"

	stmt, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.ExecContext(
		ctx,
		user.UUID,
		user.Email,
		user.Password,
	)

	if err != nil {
		return err
	}

	return nil
}

func (r *userRepository) GetUserById(ctx context.Context, id uuid.UUID) (*models.User, error) {
	query := "SELECT (uuid, email, password, created_at, updated_at) FROM users WHERE uuid='?'"

	row := r.db.QueryRowContext(ctx, query, id)

	var user models.User

	err := row.Scan(
		&user.UUID,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}

func (r *userRepository) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	query := "SELECT (uuid, email, password, created_at, updated_at) FROM user WHERE email='?'"

	row := r.db.QueryRowContext(ctx, query, email)

	var user models.User

	err := row.Scan(
		&user.UUID,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}

func (r *userRepository) UpdateUser(ctx context.Context, user *models.User) error {
	query := "UPDATE users SET email='?', password='?', updated_at='?'"

	stmt, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.ExecContext(
		ctx,
		user.Email,
		user.Password,
		user.UpdatedAt,
	)

	if err != nil {
		return err
	}

	return nil
}

func (r *userRepository) DeleteUser(ctx context.Context, user *models.User) error {
	query := "DELETE FROM users WHERE uuid='?'"

	stmt, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.ExecContext(
		ctx,
		user.UUID,
	)
	if err != nil {
		return nil
	}

	return nil
}
