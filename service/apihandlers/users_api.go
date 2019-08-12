package apihandlers

import (
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
	handler.PUT("/update-user", api.updateUserHandler)
}

func (api *UsersAPI) getUserHandler(c echo.Context) error {
	//api.service.GetAllUsers();
	return c.String(http.StatusOK, string(":::"))
}

func (api *UsersAPI) getUsersHandler(c echo.Context) error {
	return c.String(http.StatusOK, string(":::"))
}

func (api *UsersAPI) addUserHandler(c echo.Context) error {
	return c.String(http.StatusOK, string(":::"))
}

func (api *UsersAPI) deleteUserHandler(c echo.Context) error {
	return c.String(http.StatusOK, string(":::"))
}

func (api *UsersAPI) updateUserHandler(c echo.Context) error {
	return c.String(http.StatusOK, string(":::"))
}
