package connections

import (
	"chat-app/service/model"
	"fmt"

	"github.com/go-xorm/xorm"
	_ "github.com/mattn/go-sqlite3"
)

// ConnectionManager for managing all db related connections
type ConnectionManager struct {
	Engine *xorm.Engine
}

// Init initialize new db connection manager
func (c *ConnectionManager) Init() (*xorm.Engine, error) {
	_engine, err := xorm.NewEngine("sqlite3", "./chat.db")
	c.Engine = _engine
	if err != nil {
		fmt.Printf("database initialization failed %v", err)
		return nil, err
	}
	if c.Engine == nil {
		fmt.Println("engine initialization failed..")
		return nil, err
	}
	c.createTableCheck()
	return _engine, nil
}

// CreateTableCheck executes the script for the first time
func (c *ConnectionManager) createTableCheck() {
	fmt.Println("syncing all table schemas...")
	if err := c.Engine.Sync(new(model.User)); err != nil {
		fmt.Printf("table syncing failed %v", err)
	}
	if err := c.Engine.Sync(new(model.Message)); err != nil {
		fmt.Printf("table syncing failed %v", err)
	}
	if err := c.Engine.Sync(new(model.Group)); err != nil {
		fmt.Printf("table syncing failed %v", err)
	}
	if err := c.Engine.Sync(new(model.UserGroup)); err != nil {
		fmt.Printf("table syncing failed %v", err)
	}
	if err := c.Engine.Sync(new(model.MessageRecipient)); err != nil {
		fmt.Printf("table syncing failed %v", err)
	}
}
