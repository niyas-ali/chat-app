package services

import (
	"chat-app/service/model"
	"fmt"

	"github.com/go-xorm/xorm"
)

// UsersService struct
type UsersService struct {
	Engine *xorm.Engine
}

//AddNewUser method
func (u *UsersService) AddNewUser(user model.User) (bool, error) {
	_, err := u.Engine.Insert(&user)
	if err != nil {
		fmt.Printf("failed to add new user %v", err)
		return false, err
	}
	return true, nil

}

//UpdateUser method
func (u *UsersService) UpdateUser(user model.User) (bool, error) {
	_, err := u.Engine.Update(&user)
	if err != nil {
		fmt.Printf("failed to update the user %v", err)
		return false, err
	}
	return true, nil
}

//DeleteUser method
func (u *UsersService) DeleteUser(user model.User) (bool, error) {
	_, err := u.Engine.Delete(&user)
	if err != nil {
		fmt.Printf("failed to delete user %v", err)
		return false, err
	}
	return true, nil
}

//GetUser method
func (u *UsersService) GetUser() (model.User, error) {
	var user model.User
	err := u.Engine.Find(&user)
	if err != nil {
		fmt.Printf("failed to get user %v", err)
		return user, err
	}
	return user, nil
}

//GetAllUsers method
func (u *UsersService) GetAllUsers() ([]model.User, error) {
	var users []model.User
	err := u.Engine.Find(&users)
	if err != nil {
		fmt.Printf("failed to get users %v", err)
		return nil, err
	}
	return users, nil
}
