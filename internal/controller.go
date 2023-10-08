package internal

import (
	"github.com/labstack/echo"
)

func InitAPI() {
	e := echo.New()

	Routes(e)

	e.Start(":8080")
}
