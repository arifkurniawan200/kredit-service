package repository

const (
	insertNewCostumer  = `INSERT INTO customer(NIK, full_name, legal_name, born_place, born_date, salary, is_admin, email, password, foto_selfie, foto_ktp) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
	getCostumerByEmail = `
        SELECT
            id,
            NIK,
            full_name,
            legal_name,
            born_place,
            born_date,
            salary,
            is_admin,
            email,
            password,
            foto_selfie,
            foto_ktp,
            created_at,
            updated_at,
            deleted_at
        FROM
            customer
        WHERE
            email = ?
    `
)
