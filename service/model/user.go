package model

import (
	"time"
)

//User model
type User struct {
	ID          int
	FirstName   string
	LastName    string
	CreatedDate time.Time
	IsActive    bool
}
