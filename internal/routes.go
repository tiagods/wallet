package internal

import (
	"github.com/labstack/echo"
	"github.com/tiagods/wallet/internal/web/handler"
)

func Routes(e *echo.Echo) {
	e.GET("/health", handler.Health)
}
