package repository

import (
	"database/sql"
	models "kredit-service/internal/model"
)

type TransactionHandler struct {
	db *sql.DB
}

func (h TransactionHandler) GetTenorByID(id int) (models.Tenor, error) {
	var (
		data models.Tenor
		err  error
	)
	rows, err := h.db.Query(getTenorByID, id)
	if err != nil {
		return data, err
	}
	defer rows.Close()

	// Iterate hasil query
	for rows.Next() {
		if err = rows.Scan(&data.ID, &data.Tenor, &data.Value); err != nil {
			return data, err
		}
	}

	// Check for errors from iterating over rows
	if err = rows.Err(); err != nil {
		return data, err
	}
	return data, err
}

func NewTransactionRepository(db *sql.DB) TransactionRepository {
	return &TransactionHandler{db}
}

func (h TransactionHandler) GetUserTransactionByUserID(userID int) ([]models.Transaction, error) {
	//TODO implement me
	panic("implement me")
}

func (h TransactionHandler) GetTenorList() ([]models.Tenor, error) {
	var (
		data []models.Tenor
		err  error
	)
	rows, err := h.db.Query(getTenorList)
	if err != nil {
		return data, err
	}
	defer rows.Close()

	// Iterate hasil query
	for rows.Next() {
		var t models.Tenor
		if err = rows.Scan(&t.ID, &t.Tenor, &t.Value); err != nil {
			return data, err
		}
		data = append(data, t)
	}

	// Check for errors from iterating over rows
	if err = rows.Err(); err != nil {
		return data, err
	}
	return data, err
}
