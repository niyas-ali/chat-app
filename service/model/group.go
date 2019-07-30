package model

import "time"

//Group model
type Group struct {
	ID          int
	Name        string
	CreatedDate time.Time
	IsActive    bool
}
