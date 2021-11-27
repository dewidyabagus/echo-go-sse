package message

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	echo "github.com/labstack/echo/v4"

	"go-sse/api/v1/message/response"
	"go-sse/business/message"
)

type Controller struct {
	service message.Service
}

func NewController(service message.Service) *Controller {
	return &Controller{service}
}

func (c *Controller) FindMessageByCreatedDate(ctx echo.Context) error {
	dashboard := &response.Dashboard{Name: ctx.Request().RemoteAddr, Events: make(chan *response.Message, 10)}

	go func(db *response.Dashboard) {
		for {
			message, _ := c.service.FindMessageByCreatedDate(time.Now())

			db.Events <- response.ResponseMessage(message)
		}
	}(dashboard)

	writer := ctx.Response()
	writer.Header().Set("Access-Control-Allow-Origin", "*")
	writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	writer.Header().Set("Content-Type", "text/event-stream")
	writer.Header().Set("Cache-Control", "no-cache")
	writer.Header().Set("Connection", "keep-alive")
	writer.WriteHeader(http.StatusOK)

	timeout := time.After(1 * time.Second)
	select {
	case ev := <-dashboard.Events:
		var buf bytes.Buffer
		enc := json.NewEncoder(&buf)
		enc.Encode(ev)

		// return ctx.String(http.StatusOK, fmt.Sprintf("data: %v", buf.String()))
		writer.Write([]byte(fmt.Sprintf("data: %v\n\n", buf.String())))

	case <-timeout:
		writer.Write([]byte("data: nothing data found\n\n"))

	}

	writer.Flush()

	return nil
}
