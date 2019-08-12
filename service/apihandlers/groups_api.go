package apihandlers

import (
	"chat-app/service/services"
	"net/http"

	"github.com/go-xorm/xorm"
	"github.com/labstack/echo"
)

//GroupsAPI struct
type GroupsAPI struct {
	service services.GroupService
}

//InitializeHandler method
func (api *GroupsAPI) InitializeHandler(engine *xorm.Engine, handler *echo.Echo) {
	api.service = services.GroupService{Engine: engine}
	//group handlers
	handler.GET("/get-groups", getGroupsHandler)
	handler.POST("/add-group", addMessageHandler)
	handler.DELETE("/delete-group", deleteMessageHandler)
	// message recipient handlers
	handler.GET("/get-user-groups", getUserGroupsHandler)
	handler.POST("/add-user-group", addUserGroup)
	handler.DELETE("/delete-user-group", deleteUserGroup)
	handler.DELETE("/delete-many-user-groups", deletManyUserGroups)
}

func getGroupsHandler(c echo.Context) error {
	return c.String(http.StatusOK, string(":::"))
}

func addGroupHandler(c echo.Context) error {
	return c.String(http.StatusOK, string(":::"))
}

func deleteGroupHandler(c echo.Context) error {
	return c.String(http.StatusOK, string(":::"))
}

func getUserGroupsHandler(c echo.Context) error {
	return c.String(http.StatusOK, string(":::"))
}
func addUserGroup(c echo.Context) error {
	return c.String(http.StatusOK, string(":::"))
}
func deleteUserGroup(c echo.Context) error {
	return c.String(http.StatusOK, string(":::"))
}

func deletManyUserGroups(c echo.Context) error {
	return c.String(http.StatusOK, string(":::"))
}
