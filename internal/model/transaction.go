package models

import "time"

// Transaction adalah model struct untuk tabel transaction
type Transaction struct {
	ID               int        `json:"id"`
	UserID           int        `json:"user_id"`
	ContractNumber   string     `json:"contract_number"`
	OTR              int        `json:"OTR"`
	AdminFee         float64    `json:"admin_fee"`
	TotalInstallment int        `json:"total_installment"`
	Interest         float64    `json:"interest"`
	AssetName        string     `json:"asset_name"`
	Status           string     `json:"status"`
	CreatedAt        time.Time  `json:"created_at"`
	UpdatedAt        time.Time  `json:"updated_at"`
	DeletedAt        *time.Time `json:"deleted_at,omitempty"`
}
