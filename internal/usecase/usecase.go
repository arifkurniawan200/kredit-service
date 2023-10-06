package usecase

import (
	"github.com/labstack/echo/v4"
	models "kredit-service/internal/model"
)

type UserUcase interface {
	RegisterCustomer(ctx echo.Context, customer models.CustomerParam) error
	GetUserInfoByEmail(ctx echo.Context, email string) (models.Customer, error)
	GetUserLimit(ctx echo.Context, userID int) (models.LimitInformation, error)
	RequestLoan(ctx echo.Context, loan models.LoanRequestParam) error
	ListRequestLoan(ctx echo.Context) ([]models.CustomerLoan, error)
	BulkApproveLoanRequest(ctx echo.Context, ids []int) (res []int, err error)
}

type TransactionUcase interface {
	GetTenorList() ([]models.Tenor, error)
	CreateTransaction(ctx echo.Context, transaction models.TransactionParam) error
	GetUserSchedulePayment(ctx echo.Context, userID int) ([]models.MonthPayments, error)
	PayTransaction(ctx echo.Context, param models.PaymentParam) error
}
