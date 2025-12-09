package postgres

import (
	"langbrv/internal/core/model"

	"github.com/jmoiron/sqlx"
)

type UserRepo struct {
	db *sqlx.DB
}

func NewUserRepo(db *sqlx.DB) *UserRepo {
	return &UserRepo{
		db: db,
	}
}

func (r *UserRepo) CreateOrUpdate(user *model.User) error {
	query := `
		INSERT INTO users (id, username, created_at) 
		VALUES (:id, :username, :created_at)
		ON CONFLICT (id) DO UPDATE 
		SET created_at = EXCLUDED.created_at
	`
	_, err := r.db.NamedQuery(query, user)
	return err
}

func (r *UserRepo) GetByID(userID string) (*model.User, error) {
	var user model.User

	query := `SELECT * FROM users WHERE id = $1`

	err := r.db.Get(&user, query, userID)
	return &user, err
}
