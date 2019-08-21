package model

import "time"

//Group model
type Group struct {
	ID          int64     `xorm:"PK autoincr bigint"`
	Name        string    `xorm:"varchar(200)"`
	CreatedDate time.Time `xorm:"created"`
	Updated     time.Time `xorm:"updated"`
	IsActive    bool      `xorm:"not null"`
}
