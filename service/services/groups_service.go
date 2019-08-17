package services

import (
	"chat-app/service/model"
	"fmt"

	"github.com/go-xorm/xorm"
)

//GroupService struct
type GroupService struct {
	Engine *xorm.Engine
}

//NewGroup method
func (g *GroupService) NewGroup(group model.Group) (bool, error) {
	_, err := g.Engine.Insert(&group)
	if err != nil {
		fmt.Printf("failed to add new group %v", err)
		return false, err
	}
	return true, nil

}

//DeleteGroup method
func (g *GroupService) DeleteGroup(group model.Group) (bool, error) {
	_, err := g.Engine.Delete(&group)
	if err != nil {
		fmt.Printf("failed to delete group %v", err)
		return false, err
	}
	return true, nil
}
//DeleteGroups method
func (g *GroupService) DeleteGroups(group []model.Group) (bool, error) {
	_, err := g.Engine.Delete(&group)
	if err != nil {
		fmt.Printf("failed to delete groups %v", err)
		return false, err
	}
	return true, nil
}

//GetAllGroup method
func (g *GroupService) GetAllGroup() ([]model.Group, error) {
	var groups []model.Group
	err := g.Engine.Find(&groups)
	if err != nil {
		fmt.Printf("failed to get all groups %v", err)
		return nil, err
	}
	return groups, nil
}
//GetGroup method
func (g *GroupService) GetGroup(group model.Group) (model.Group, error) {
	_,err := g.Engine.Get(&group)
	if err != nil {
		fmt.Printf("failed to get all group %v", err)
		return group, err
	}
	return group, nil
}
//AddUserGroup method
func (g *GroupService) AddUserGroup(userGroup model.UserGroup) (bool, error) {
	_, err := g.Engine.Insert(&userGroup)
	if err != nil {
		fmt.Printf("failed to add new user group %v", err)
		return false, err
	}
	return true, nil
}

//GetAllUserGroup method
func (g *GroupService) GetAllUserGroup() ([]model.UserGroup, error) {
	var userGroups []model.UserGroup
	err := g.Engine.Find(&userGroups)
	if err != nil {
		fmt.Printf("failed to get all user groups %v", err)
		return nil, err
	}
	return userGroups, nil
}

//DeleteUserGroup method
func (g *GroupService) DeleteUserGroup(userGroup model.UserGroup) (bool, error) {
	_, err := g.Engine.Delete(&userGroup)
	if err != nil {
		fmt.Printf("failed to delete user group %v", err)
		return false, err
	}
	return true, nil
}

//DeleteUserGroups method
func (g *GroupService) DeleteUserGroups(userGroups []model.UserGroup) (bool, error) {
	_, err := g.Engine.Delete(&userGroups)
	if err != nil {
		fmt.Printf("failed to delete user groups %v", err)
		return false, err
	}
	return true, nil
}
