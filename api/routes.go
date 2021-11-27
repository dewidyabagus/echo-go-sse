package api

import (
	"net/http"

	echo "github.com/labstack/echo/v4"

	"go-sse/api/v1/message"
)

func RegisterRoutes(e *echo.Echo, message *message.Controller) {
	// ======= ROUTES =======
	e.GET("", func(c echo.Context) error {
		return c.HTML(http.StatusOK, "<h1>Hello World</h1>")
	})

	e.GET("/dashboard", message.FindMessageByCreatedDate)
	// ======= END ROUTES =======
}
