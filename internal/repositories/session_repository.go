package repositories

import (
	"context"
	"database/sql"
	"go-sqap/internal/models"
	"go-sqap/internal/utils"
)

type SessionRepository interface {
	CreateSession(ctx context.Context, session *models.Session) error
}

type sessionRepository struct {
	db     *sql.DB
	logger *utils.Logger
}

func NewSessionRepository(db *sql.DB, logger *utils.Logger) SessionRepository {
	return &sessionRepository{
		db:     db,
		logger: logger,
	}
}

func (r *sessionRepository) CreateSession(ctx context.Context, session *models.Session) error {
	query := "INSERT INTO sessions (user_id, token, created_at, updated_at) VALUES (?, ?, ?, ?)"

	stmt, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(
		ctx,
		&session.UserID,
		&session.Token,
		&session.CreatedAt,
		&session.UpdatedAt,
	)
	if err != nil {
		return err
	}

	return nil
}
