package model

import "time"

//Message model
type Message struct {
	ID              int
	Subject         string
	MessageBody     string
	CreatorID       int
	ParentMessageID int
	CreatedDate     time.Time
	ExpiryDate      time.Time
}
