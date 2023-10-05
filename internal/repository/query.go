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
	getTenorList     = `SELECT id,tenor,value FROM tenor`
	getTenorByID     = `SELECT id,tenor,value FROM tenor WHERE id = ?`
	queryRequestLoan = `
		INSERT INTO customer_loan (customer_id, tenor, loan_date, loan_amount, status)
		VALUES (?, ?, ?, ?, ?)
	`
	queryGetUserLimit = `  SELECT
            id,
            customer_id,
            status,
            used_amount,
            loan_amount,
            tenor,
            expired_at
        FROM
            customer_loan WHERE customer_id=?`
	queryGetListCustomerRequest = `
        SELECT
            id,
            customer_id,
            status,
            loan_amount,
            used_amount,
            tenor,
            created_at,
            updated_at,
            loan_date
        FROM
            customer_loan
        WHERE
            status = ? ORDER BY created_at desc 
    `
	queryGetListCustomerRequestByIds = `
        SELECT
            id,
            customer_id,
            status,
            loan_amount,
            used_amount,
            tenor,
            created_at,
            updated_at,
            loan_date
        FROM
            customer_loan
        WHERE
            status = ? AND id IN (%s)
    `
	queryApproveLoanRequest = `
			UPDATE customer_loan
			SET
			  status = ?,
			  approved_date = ?,
			  expired_at = ?
			WHERE
			  id = ?
`
	queryUseLimitRequest = `
			UPDATE customer_loan
			SET
			  status = ?,
			  used_amount = ?
			WHERE
			  id = ?`

	queryCreateTransaction = `INSERT INTO transaction (customer_id, contract_number, OTR, admin_fee, total_installment, interest, asset_name, status)
							VALUES (?, ?, ?, ?, ?, ?, ?, ?);`

	queryCreateSchedulePayment = `INSERT INTO schedule_payment (transaction_id, amount, status, due_date)
								VALUES(?,?,?,?)`
)
