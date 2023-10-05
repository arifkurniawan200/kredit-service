package repository

import (
	"context"
	"database/sql"
	models "kredit-service/internal/model"
)

type UserHandler struct {
	db *sql.DB
}

// `INSERT INTO customer (NIK, full_name, legal_name, born_place, born_date, salary, is_admin, email, password, foto_selfie, foto_ktp)
// VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11);`
func (h UserHandler) RegisterUser(c models.CustomerParam) error {
	_, err := h.db.Exec(insertNewCostumer, c.NIK, c.FullName, c.LegalName, c.BornPlace, c.BornDate, c.Salary, false, c.Email, c.Password, c.FotoSelfie, c.FotoKTP)
	if err != nil {
		return err
	}
	return err
}

func (h UserHandler) GetUserByEmail(email string) (models.Customer, error) {
	var (
		data models.Customer
		err  error
	)
	rows, err := h.db.Query(getCostumerByEmail, email)
	if err != nil {
		return data, err
	}
	defer rows.Close()

	// Iterate hasil query
	for rows.Next() {
		if err = rows.Scan(&data.ID, &data.NIK, &data.FullName, &data.LegalName, &data.BornPlace, &data.BornDate, &data.Salary, &data.IsAdmin,
			&data.Email, &data.Password, &data.FotoSelfie, &data.FotoKTP, &data.CreatedAt, &data.UpdatedAt, &data.DeletedAt,
		); err != nil {
			return data, err
		}
	}

	// Check for errors from iterating over rows
	if err = rows.Err(); err != nil {
		return data, err
	}
	return data, err
}

func (h UserHandler) BeginTx() (*sql.Tx, error) {
	return h.db.BeginTx(context.Background(), &sql.TxOptions{
		Isolation: sql.LevelSerializable,
	})
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &UserHandler{db}
}
