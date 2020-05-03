package responses

import (
	"net/http"

	"github.com/labstack/echo"
)

func Response(c echo.Context, statusCode int, data interface{}) error {
	// nolint // context.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	// nolint // context.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
	// nolint // context.Writer.Header().Set("Access-Control-Allow-Headers", "Authorization")
	return c.JSON(statusCode, data)
}

func SuccessResponse(c echo.Context, data interface{}) error {
	return Response(c, http.StatusOK, data)
}

func ErrorResponse(c echo.Context, statusCode int, message string) error {
	return Response(c, statusCode, struct {
		Code  int
		Error string
	}{
		Code:  statusCode,
		Error: message,
	})
}
