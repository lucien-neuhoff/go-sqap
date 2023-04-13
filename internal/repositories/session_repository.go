package repositories

import (
	"database/sql"
	"go-sqap/internal/models"
	"go-sqap/internal/utils"
)

type SessionRepository interface {
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

func (r *sessionRepository) Create(user_id string) (models.Session, error) {
	var session models.Session

	return session, nil
}
