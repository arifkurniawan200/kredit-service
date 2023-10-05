package repository

import (
	"context"
	"database/sql"
	models "kredit-service/internal/model"
)

type Handler struct {
	db *sql.DB
}

func (h Handler) GetUserByEmail(email string) (models.Customer, error) {
	//TODO implement me
	panic("implement me")
}

func (h Handler) BeginTx() (*sql.Tx, error) {
	return h.db.BeginTx(context.Background(), &sql.TxOptions{
		Isolation: sql.LevelSerializable,
	})
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &Handler{db}
}
