package handler

import (
	"net/http"

	"github.com/labstack/echo"
)

func GetAccount(e echo.Context) error {
	return e.JSON(http.StatusOK, "GetAccount")
}
