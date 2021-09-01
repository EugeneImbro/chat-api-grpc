package repository

import (
	"fmt"
	"github.com/EugeneImbro/chat-backend/internal/model"
	"github.com/jmoiron/sqlx"
)

type UserPostgres struct {
	db *sqlx.DB
}

func NewUserPostgres(db *sqlx.DB) *UserPostgres {
	return &UserPostgres{db: db}
}

func (r *UserPostgres) GetById(id int32) (*model.User, error) {
	var user model.User
	query := fmt.Sprintf("SELECT * FROM %s WHERE id=$1", usersTable)
	if err := r.db.Get(&user, query, id); err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserPostgres) GetByNickName(nickName string) (*model.User, error) {
	var user model.User
	query := fmt.Sprintf("SELECT * FROM %s WHERE nickname=$1", usersTable)
	if err := r.db.Get(&user, query, nickName); err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserPostgres) GetAll() ([]*model.User, error) {
	var users []*model.User
	query := fmt.Sprintf("SELECT * FROM %s", usersTable)
	if err := r.db.Select(users, query); err != nil {
		return nil, err
	}
	return users, nil
}
