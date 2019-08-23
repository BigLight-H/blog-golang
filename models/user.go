package models

import "time"

type User struct {
	Id int
	Username string
	Password string
	Email string
	LoginCount int
	LastTime time.Time
	LastIp string
	Status int
	Mobile string
	HeadImg *HeadImg  `orm:"rel(one)"`  //头像一对一
	Created time.Time
	Updated time.Time
	Role *Role   `orm:"rel(one)"` //等级一对一
}

func (m *User) TableName() string {
	return TableName("user")
}