package usecase

import (
	"github.com/labstack/echo/v4"
	models "kredit-service/internal/model"
)

type UserUcase interface {
	RegisterCustomer(ctx echo.Context, customer models.CustomerParam) error
	GetUserInfoByEmail(ctx echo.Context, email string) (models.Customer, error)
}
