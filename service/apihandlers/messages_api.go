package apihandlers

import (
	"chat-app/service/model"
	"chat-app/service/services"
	"net/http"

	"github.com/go-xorm/xorm"
	"github.com/labstack/echo"
)

//MessageAPI struct
type MessageAPI struct {
	service                 services.MessageService
}

//InitializeHandler method
func (api *MessageAPI) InitializeHandler(engine *xorm.Engine, handler *echo.Echo) {
	api.service = services.MessageService{Engine: engine}
	//message handlers
	handler.GET("/get-message", api.getMessageHandler)
	handler.GET("/get-messages", api.getMessagesHandler)
	handler.POST("/add-message", api.addMessageHandler)
	handler.DELETE("/delete-message", api.deleteMessageHandler)
	handler.DELETE("/delete-messages", api.deleteMessagesHandler)
	// message recipient handlers
	handler.GET("/get-message-recipients", api.getMessageRecipientsHandler)
	handler.POST("/add-message-recipient", api.addMessageRecipient)
	handler.DELETE("/delete-message-recipient", api.deleteMessageRecipient)
	handler.DELETE("/delete-message-recipients", api.deleteMessageRecipients)
}

func (api *MessageAPI) getMessagesHandler(c echo.Context) error {
	messages, err := api.service.GetAllMessages()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSONPretty(http.StatusOK, messages, " ")
}
func (api *MessageAPI) getMessageHandler(c echo.Context) error {
	var message model.Message
	if  err:=c.Bind(&message); err!=nil{
		return c.JSON(http.StatusInternalServerError, err)
	}
	_message, err := api.service.GetMessage(message)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSONPretty(http.StatusOK, _message, " ")
}
func (api *MessageAPI) addMessageHandler(c echo.Context) error {
	var message model.Message
	if err:=c.Bind(&message);err!=nil{
		return c.JSON(http.StatusInternalServerError, err)
	}
	result,err:=api.service.NewMessage(message)
	if err!= nil{
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSONPretty(http.StatusOK, result, " ")
}

func (api *MessageAPI) deleteMessageHandler(c echo.Context) error {
	var message model.Message
	if err := c.Bind(&message); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	result, err := api.service.DeleteMessage(message)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSONPretty(http.StatusOK, result, " ")
}

func (api *MessageAPI) deleteMessagesHandler(c echo.Context) error {
	var messages []model.Message
	if err := c.Bind(&messages); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	result, err := api.service.DeleteMessages(messages)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSONPretty(http.StatusOK, result, " ")
}

func (api *MessageAPI) getMessageRecipientsHandler(c echo.Context) error {
	messageRecipients, err := api.service.GetAllMessageRecipient()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSONPretty(http.StatusOK, messageRecipients, " ")
}
func (api *MessageAPI) addMessageRecipient(c echo.Context) error {
	var recipient model.MessageRecipient
	if err:=c.Bind(&recipient);err!=nil{
		return c.JSON(http.StatusInternalServerError, err)
	}
	result,err:=api.service.AddMessageRecipient(recipient)
	if err!= nil{
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSONPretty(http.StatusOK, result, " ")
}
func (api *MessageAPI) deleteMessageRecipient(c echo.Context) error {
	var recipient model.MessageRecipient
	if err := c.Bind(&recipient); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	result, err := api.service.DeleteMessageRecipient(recipient)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSONPretty(http.StatusOK, result, " ")
}
func (api *MessageAPI) deleteMessageRecipients(c echo.Context) error {
	var recipients []model.MessageRecipient
	if err := c.Bind(&recipients); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	result, err := api.service.DeleteMessageRecipients(recipients)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSONPretty(http.StatusOK, result, " ")
}