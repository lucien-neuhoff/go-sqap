package repositories

import (
	"database/sql"
	"go-sqap/internal/utils"
)

type UserRepository interface {
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
