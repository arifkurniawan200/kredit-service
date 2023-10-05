package app

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func (u handler) ListCostumerLoan(c echo.Context) error {
	data, err := u.User.ListRequestLoan(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ResponseFailed{
			Messages: "failed to connect database",
			Error:    err.Error(),
		})
	}
	return c.JSON(http.StatusCreated, ResponseSuccess{
		Messages: "success fetch loan request list",
		Data:     data,
	})
}
