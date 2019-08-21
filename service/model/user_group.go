package model

import "time"

// UserGroup model
type UserGroup struct {
	ID          int64     `xorm:"PK autoincr bigint"`
	UserID      int64     `xorm:"bigint"`
	GroupID     int64     `xorm:"bigint"`
	CreatedDate time.Time `xorm:"created"`
	UpdatedDate time.Time `xorm:"updated"`
	IsActive    bool      `xorm:"not null"`
}
