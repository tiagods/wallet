package handler

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/tiagods/wallet/internal/web/presenter"
)

func Health(c echo.Context) error {
	return c.JSON(http.StatusOK, &presenter.Health{Status: "OK"})
}
