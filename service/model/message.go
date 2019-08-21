package model

import "time"

//Message model
type Message struct {
	ID              int64     `xorm:"PK autoincr bigint"`
	Subject         string    `xorm:"text"`
	MessageBody     string    `xorm:"text"`
	CreatorID       int64     `xorm:"bigint"`
	ParentMessageID int64     `xorm:"bigint"`
	CreatedDate     time.Time `xorm:"created"`
	UpdatedDate     time.Time `xorm:"updated"`
	ExpiryDate      time.Time
}
