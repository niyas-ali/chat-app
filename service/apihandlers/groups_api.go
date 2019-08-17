package apihandlers

import (
	"chat-app/service/model"
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
	handler.GET("/get-groups", api.getGroupsHandler)
	handler.GET("/get-group", api.getGroupHandler)
	handler.POST("/add-group", api.addGroupHandler)
	handler.DELETE("/delete-group", api.deleteGroupHandler)
	handler.DELETE("/delete-groups", api.deleteGroupsHandler)
	// message recipient handlers
	handler.GET("/get-user-groups", api.getUserGroupsHandler)
	handler.POST("/add-user-group", api.addUserGroup)
	handler.DELETE("/delete-user-group", api.deleteUserGroup)
	handler.DELETE("/delete-user-groups", api.deleteUserGroups)
}

func (api *GroupsAPI) getGroupsHandler(c echo.Context) error {
	groups, err := api.service.GetAllGroup()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSONPretty(http.StatusOK, groups, " ")
}
func (api *GroupsAPI) getGroupHandler(c echo.Context) error {
	var group model.Group
	if  err:=c.Bind(&group); err!=nil{
		return c.JSON(http.StatusInternalServerError, err)
	}
	_message, err := api.service.GetGroup(group)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSONPretty(http.StatusOK, _message, " ")
}

func (api *GroupsAPI) addGroupHandler(c echo.Context) error {
	var group model.Group
	if err:=c.Bind(&group);err!=nil{
		return c.JSON(http.StatusInternalServerError, err)
	}
	result,err:=api.service.NewGroup(group)
	if err!= nil{
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSONPretty(http.StatusOK, result, " ")
}


func (api *GroupsAPI) deleteGroupHandler(c echo.Context) error {
	var group model.Group
	if err:=c.Bind(&group);err!=nil{
		return c.JSON(http.StatusInternalServerError, err)
	}
	result, err := api.service.DeleteGroup(group)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSONPretty(http.StatusOK, result, " ")
}
func (api *GroupsAPI) deleteGroupsHandler(c echo.Context) error {
	var group []model.Group
	if err := c.Bind(&group); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	result, err := api.service.DeleteGroups(group)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSONPretty(http.StatusOK, result, " ")
}

func (api *GroupsAPI) getUserGroupsHandler(c echo.Context) error {
	userGroups, err := api.service.GetAllUserGroup()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSONPretty(http.StatusOK, userGroups, " ")
}
func (api *GroupsAPI) addUserGroup(c echo.Context) error {
	var userGroup model.UserGroup
	if err:=c.Bind(&userGroup);err!=nil{
		return c.JSON(http.StatusInternalServerError, err)
	}
	result,err:=api.service.AddUserGroup(userGroup)
	if err!= nil{
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSONPretty(http.StatusOK, result, " ")
}
func (api *GroupsAPI) deleteUserGroup(c echo.Context) error {
	var userGroup model.UserGroup
	if err:=c.Bind(&userGroup);err!=nil{
		return c.JSON(http.StatusInternalServerError, err)
	}
	result, err := api.service.DeleteUserGroup(userGroup)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSONPretty(http.StatusOK, result, " ")
}

func (api *GroupsAPI) deleteUserGroups(c echo.Context) error {
	var userGroups []model.UserGroup
	if err := c.Bind(&userGroups); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	result, err := api.service.DeleteUserGroups(userGroups)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSONPretty(http.StatusOK, result, " ")
}
