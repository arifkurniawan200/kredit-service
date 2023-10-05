package repository

import (
	"database/sql"
	models "kredit-service/internal/model"
)

type UserRepository interface {
	RegisterUser(user models.CustomerParam) error
	GetUserByEmail(email string) (models.Customer, error)
	BeginTx() (*sql.Tx, error)
	RequestLoan(user models.CustomerLoan) error
	GetUserLimit(userID int) (models.CustomerLoan, error)
	CustomerLoanRequest(status string) ([]models.CustomerLoan, error)
	CustomerLoanRequestByIds(ids []int, status string) ([]models.CustomerLoan, error)
	UpdateLoanRequest(loan models.CustomerLoan) error
}

type TransactionRepository interface {
	GetUserTransactionByUserID(userID int) ([]models.Transaction, error)
	GetTenorList() ([]models.Tenor, error)
	GetTenorByID(id int) (models.Tenor, error)
}
