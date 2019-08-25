package models

import "time"

type Permissions struct {
	Id 			  int
	Pid           int
	Status        int
	WebUrl       string
	DisplayName  string
	Sort          int
	IsMenu       int
	Description   string
	Icon          string
	CreatedAt 	  time.Time
	UpdatedAt 	  time.Time
}

func (m *Permissions) TableName() string {
	return TableName("permissions")
}