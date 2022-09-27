package helpers

import "github.com/labstack/echo/v4"

type ResponseModel struct {
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
	Status  bool        `json:"status"`
}

func Response(c echo.Context, code int, model ResponseModel) error {
	return c.JSON(code, model)
}
