package repository

import (
	"database/sql"
	models "kredit-service/internal/model"
)

type TransactionHandler struct {
	db *sql.DB
}

func NewTransactionRepository(db *sql.DB) TransactionRepository {
	return &TransactionHandler{db}
}

func (h TransactionHandler) GetUserTransactionByUserID(userID int) ([]models.Transaction, error) {
	//TODO implement me
	panic("implement me")
}
