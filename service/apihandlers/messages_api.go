package apihandlers

import (
	"chat-app/service/services"
	"net/http"

	"github.com/go-xorm/xorm"
	"github.com/labstack/echo"
)

//MessageAPI struct
type MessageAPI struct {
	service services.MessageService
}

//InitializeHandler method
func (api *MessageAPI) InitializeHandler(engine *xorm.Engine, handler *echo.Echo) {
	api.service = services.MessageService{Engine: engine}
	//message handlers
	handler.GET("/get-messages", getMessagesHandler)
	handler.POST("/add-message", addMessageHandler)
	handler.DELETE("/delete-message", deleteMessageHandler)
	// message recipient handlers
	handler.GET("/get-message-recipients", getMessageRecipientsHandler)
	handler.POST("/add-message-recipient", addMessageRecipient)
	handler.DELETE("/delete-message-recipient", deleteMessageRecipient)
	handler.DELETE("/delete-many-message-recipients", deletManyMessageRecipienta)
}

func getMessagesHandler(c echo.Context) error {
	return c.String(http.StatusOK, string(":::"))
}

func addMessageHandler(c echo.Context) error {
	return c.String(http.StatusOK, string(":::"))
}

func deleteMessageHandler(c echo.Context) error {
	return c.String(http.StatusOK, string(":::"))
}

func getMessageRecipientsHandler(c echo.Context) error {
	return c.String(http.StatusOK, string(":::"))
}
func addMessageRecipient(c echo.Context) error {
	return c.String(http.StatusOK, string(":::"))
}
func deleteMessageRecipient(c echo.Context) error {
	return c.String(http.StatusOK, string(":::"))
}

func deletManyMessageRecipienta(c echo.Context) error {
	return c.String(http.StatusOK, string(":::"))
}
