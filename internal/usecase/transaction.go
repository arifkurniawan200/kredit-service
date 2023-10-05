package usecase

import (
	models "kredit-service/internal/model"
	"kredit-service/internal/repository"
)

type TransactionHandler struct {
	u repository.TransactionRepository
}

func (t TransactionHandler) GetTenorList() ([]models.Tenor, error) {
	return t.u.GetTenorList()
}

func NewTransactionsUsecase(u repository.TransactionRepository) TransactionUcase {
	return &TransactionHandler{u}
}
