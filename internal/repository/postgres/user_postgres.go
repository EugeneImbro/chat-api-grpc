package postgres

import (
	"github.com/jmoiron/sqlx"

	"github.com/EugeneImbro/chat-backend/internal/model"
)

type UserPostgres struct {
	db *sqlx.DB
}

func NewUserPostgres(db *sqlx.DB) *UserPostgres {
	return &UserPostgres{db: db}
}

func (r *UserPostgres) GetById(id int32) (*model.User, error) {
	var user model.User
	query := "SELECT * FROM users WHERE id=$1"
	if err := r.db.Get(&user, query, id); err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserPostgres) GetByNickName(nickName string) (*model.User, error) {
	var user model.User
	query := "SELECT * FROM users WHERE nickname=$1"
	if err := r.db.Get(&user, query, nickName); err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserPostgres) GetAll() ([]*model.User, error) {
	var users []*model.User
	query := "SELECT * FROM users"
	if err := r.db.Select(users, query); err != nil {
		return nil, err
	}
	return users, nil
}
