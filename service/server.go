package main

import (
	"chat-app/service/apihandlers"
	"chat-app/service/connections"
	"fmt"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	fmt.Println("Server started ...")
	e := echo.New()
	e.Use(middleware.Logger())
	//e.Use(middleware.Recover())
	database := &connections.ConnectionManager{}
	database.Init()
	_usersAPI := apihandlers.UsersAPI{}
	_groupsAPI := apihandlers.GroupsAPI{}
	_messageAPI := apihandlers.MessageAPI{}
	_usersAPI.InitializeHandler(database.Engine, e)
	_groupsAPI.InitializeHandler(database.Engine, e)
	_messageAPI.InitializeHandler(database.Engine, e)
	e.Logger.Fatal(e.Start(":1323"))
}
