package model

import "time"

// UserGroup model
type UserGroup struct {
	ID          int
	UserID      int
	GroupID     int
	CreatedDate time.Time
	IsActive    bool
}
