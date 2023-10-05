package app

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func (u handler) TenorList(c echo.Context) error {
	data, err := u.Transaction.GetTenorList()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ResponseFailed{
			Messages: "failed to connect database",
			Error:    err.Error(),
		})
	}
	return c.JSON(http.StatusOK, ResponseSuccess{
		Messages: "success fetch tenor list",
		Data:     data,
	})
}
