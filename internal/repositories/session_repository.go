package repositories

import (
	"context"
	"database/sql"
	"go-sqap/internal/models"
	"go-sqap/internal/utils"
	"time"

	"github.com/google/uuid"
)

type SessionRepository interface {
	CreateSession(ctx context.Context, user_id string) (*models.Session, error)
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

func (r *sessionRepository) CreateSession(ctx context.Context, user_id string) (*models.Session, error) {
	var session *models.Session

	session.UUID = uuid.New().String()

	now := time.Now().UTC()
	session.CreatedAt = now
	session.UpdatedAt = now

	session.Token = utils.GenerateToken(255)

	query := "INSERT INTO session (uuid, user_id, token, created_at, updated_at) VALUES (?, ?, ?, ?, ?)"

	stmt, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(
		ctx,
		session.UUID,
		session.UserID,
		session.Token,
		session.CreatedAt,
		session.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return session, nil
}
