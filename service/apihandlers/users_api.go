package apihandlers

import (
	"chat-app/service/model"
	"chat-app/service/services"
	"net/http"

	"github.com/go-xorm/xorm"
	"github.com/labstack/echo"
)

//UsersAPI struct
type UsersAPI struct {
	service services.UsersService
}

//InitializeHandler method
func (api *UsersAPI) InitializeHandler(engine *xorm.Engine, handler *echo.Echo) {
	api.service = services.UsersService{Engine: engine}
	handler.GET("/get-user", api.getUserHandler)
	handler.GET("/get-users", api.getUsersHandler)
	handler.POST("/add-user", api.addUserHandler)
	handler.DELETE("/delete-user", api.deleteUserHandler)
	handler.DELETE("/delete-users", api.deleteUsersHandler)
	handler.PUT("/update-user", api.updateUserHandler)
}

func (api *UsersAPI) getUserHandler(c echo.Context) error {
	var user model.User
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	_user, err := api.service.GetUser(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSONPretty(http.StatusOK, _user, " ")
}

func (api *UsersAPI) getUsersHandler(c echo.Context) error {
	users, err := api.service.GetAllUsers()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSONPretty(http.StatusOK, users, " ")
}

func (api *UsersAPI) addUserHandler(c echo.Context) error {
	var user model.User
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	result, err := api.service.AddNewUser(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSONPretty(http.StatusOK, result, " ")
}

func (api *UsersAPI) deleteUserHandler(c echo.Context) error {
	var user model.User
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	result, err := api.service.DeleteUser(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSONPretty(http.StatusOK, result, " ")
}

func (api *UsersAPI) deleteUsersHandler(c echo.Context) error {
	var users []model.User
	if err := c.Bind(&users); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	result, err := api.service.DeleteUsers(users)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSONPretty(http.StatusOK, result, " ")
}

func (api *UsersAPI) updateUserHandler(c echo.Context) error {
	var user model.User
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	result, err := api.service.UpdateUser(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSONPretty(http.StatusOK, result, " ")
}
