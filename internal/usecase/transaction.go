package usecase

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"kredit-service/internal/consts"
	models "kredit-service/internal/model"
	"kredit-service/internal/repository"
	"time"
)

type TransactionHandler struct {
	t repository.TransactionRepository
	u repository.UserRepository
}

func (t TransactionHandler) CreateTransaction(ctx echo.Context, trx models.TransactionParam) error {
	limit, err := t.u.GetUserLimit(trx.UserID)
	if err != nil {
		return err
	}

	if limit.ID == 0 {
		return fmt.Errorf("you don't have a loan limit yet. Please apply for a loan")
	}

	trx.GenerateContractNumber()

	switch limit.Status {
	case consts.LoanRequestStatusExpired:
		err = fmt.Errorf("your limit already expired, please contact admin")
	case consts.LoanRequestStatusRequested:
		err = fmt.Errorf("your limit need approval, please wait before doing transaction")
	}
	if err != nil {
		return err
	}

	amount := limit.LoanAmount - limit.UsedAmount

	if trx.OTR > amount {
		return fmt.Errorf("available amount not enough to doing transaction")
	}

	amountWithInterest := trx.OTR + trx.AdminFee + (trx.Interest * trx.OTR)
	amountInstallment := amountWithInterest / float64(trx.TotalInstallment)
	trx.Status = consts.TransactionStatusSuccess

	//start operation using database transaction
	tx, err := t.u.BeginTx()
	if err != nil {
		return err
	}

	err = t.u.UpdateLoanRequestTx(tx, models.CustomerLoan{
		ID:         limit.ID,
		UsedAmount: limit.UsedAmount + amountWithInterest,
		Status:     consts.LoanRequestStatusUsed,
	})
	if err != nil {
		tx.Rollback()
		return err
	}

	id, err := t.t.CreateTransactionTx(tx, trx)
	if err != nil {
		tx.Rollback()
		return err
	}

	for i := 1; i <= trx.TotalInstallment; i++ {
		now := time.Now()
		err = t.t.CreateSchedulePaymentTx(tx, models.SchedulePayment{
			TransactionID: id,
			Status:        consts.ScheduleStatusOnGoing,
			DueDate:       now.AddDate(0, i, 0),
			Amount:        amountInstallment,
		})
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	err = tx.Commit()
	if err != nil {
		return err
	}
	return err

}

func (t TransactionHandler) GetTenorList() ([]models.Tenor, error) {
	return t.t.GetTenorList()
}

func NewTransactionsUsecase(t repository.TransactionRepository, u repository.UserRepository) TransactionUcase {
	return &TransactionHandler{t, u}
}
