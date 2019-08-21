package services

import (
	"chat-app/service/model"
	"fmt"

	"github.com/go-xorm/xorm"
)

// MessageService struct
type MessageService struct {
	Engine *xorm.Engine
}

//NewMessage method
func (m *MessageService) NewMessage(message model.Message) (bool, error) {
	_, err := m.Engine.Insert(&message)
	if err != nil {
		fmt.Printf("failed to add new message %v", err)
		return false, err
	}
	return true, nil

}

//DeleteMessage method
func (m *MessageService) DeleteMessage(message model.Message) (bool, error) {
	_, err := m.Engine.Delete(&message)
	if err != nil {
		fmt.Printf("failed to delete message %v", err)
		return false, err
	}
	return true, nil
}

//DeleteMessages method
func (m *MessageService) DeleteMessages(message []model.Message) (bool, error) {
	_, err := m.Engine.Delete(&message)
	if err != nil {
		fmt.Printf("failed to delete messages %v", err)
		return false, err
	}
	return true, nil
}

//GetAllMessage method
func (m *MessageService) GetMessage(message model.Message) (model.Message, error) {
	_, err := m.Engine.Get(&message)
	if err != nil {
		fmt.Printf("failed to get message %v", err)
		return message, err
	}
	return message, nil
}

//GetAllMessages method
func (m *MessageService) GetAllMessages() ([]model.Message, error) {
	var messages []model.Message
	err := m.Engine.Find(&messages)
	if err != nil {
		fmt.Printf("failed to get messages %v", err)
		return nil, err
	}
	return messages, nil
}

//AddMessageRecipient method
func (m *MessageService) AddMessageRecipient(recipient model.MessageRecipient) (bool, error) {
	_, err := m.Engine.Insert(&recipient)
	if err != nil {
		fmt.Printf("failed to add new message recipient %v", err)
		return false, err
	}
	return true, nil
}

//GetAllMessageRecipient method
func (m *MessageService) GetAllMessageRecipient() ([]model.MessageRecipient, error) {
	var recipients []model.MessageRecipient
	err := m.Engine.Find(&recipients)
	if err != nil {
		fmt.Printf("failed to get all message recipients %v", err)
		return nil, err
	}
	return recipients, nil
}

//DeleteMessageRecipient method
func (m *MessageService) DeleteMessageRecipient(recipient model.MessageRecipient) (bool, error) {
	_, err := m.Engine.Delete(&recipient)
	if err != nil {
		fmt.Printf("failed to delete message recipient %v", err)
		return false, err
	}
	return true, nil
}

//DeleteMessageRecipients method
func (m *MessageService) DeleteMessageRecipients(recipients []model.MessageRecipient) (bool, error) {
	_, err := m.Engine.Delete(&recipients)
	if err != nil {
		fmt.Printf("failed to delete message recipients %v", err)
		return false, err
	}
	return true, nil
}
