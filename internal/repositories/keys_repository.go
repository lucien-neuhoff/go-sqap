package repositories

import (
	"context"
	"database/sql"
	"errors"
	"go-sqap/internal/models"
	"go-sqap/internal/utils"
)

type KeysRepository interface {
	SaveUserPublicKey(ctx context.Context, public_key models.PublicKey) error
	GetUserPublicKey(ctx context.Context, user *models.User) (publicKey *string, err error)
}

type keysRepository struct {
	db     *sql.DB
	logger *utils.Logger
}

func NewKeysRepository(db *sql.DB, logger *utils.Logger) KeysRepository {
	return &keysRepository{
		db:     db,
		logger: logger,
	}
}

func (r *keysRepository) SaveUserPublicKey(ctx context.Context, publicKey models.PublicKey) error {
	exists, err := r.UserPublicKeyExists(ctx, &models.User{UUID: publicKey.UserID})
	if err != nil {
		return err
	}

	query := ""
	if !exists {
		query = "INSERT INTO public_keys (public_key, user_id) VALUES (?, ?)"
	} else {
		query = "UPDATE public_keys SET public_key=? WHERE user_id=?"
	}

	stmt, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(
		ctx,
		publicKey.Key,
		&publicKey.UserID,
	)

	if err != nil {
		return err
	}

	return nil
}

func (r *keysRepository) UserPublicKeyExists(ctx context.Context, user *models.User) (bool, error) {
	publicKey, err := r.GetUserPublicKey(ctx, user)
	if err != nil {
		return false, err
	}

	if publicKey == nil {
		return false, nil
	}

	return true, nil
}

func (r *keysRepository) GetUserPublicKey(ctx context.Context, user *models.User) (*string, error) {
	query := "SELECT public_key FROM public_keys WHERE user_id=?"

	row := r.db.QueryRowContext(ctx, query, user.UUID)

	var publicKey string

	err := row.Scan(
		&publicKey,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &publicKey, nil
}
