package repository

import (
	"database/sql"
	models "kredit-service/internal/model"
)

type UserRepository interface {
	RegisterUser(user models.CustomerParam) error
	GetUserByEmail(email string) (models.Customer, error)
	BeginTx() (*sql.Tx, error)
}

type TransactionRepository interface {
	GetUserTransactionByUserID(userID int) ([]models.Transaction, error)
}
