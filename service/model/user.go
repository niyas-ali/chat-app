package model

import (
	"time"
)

//User model
type User struct {
	//
	ID          int64     `xorm:"PK autoincr bigint"`
	FirstName   string    `xorm:"varchar(200)"`
	LastName    string    `xorm:"varchar(200)"`
	CreatedDate time.Time `xorm:"created"`
	UpdatedDate time.Time `xorm:"updated"`
	IsActive    bool      `xorm:"not null"`
}
