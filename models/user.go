package models

import "time"

type User struct {
	Id int
	Username string `validate:"required||string=2,10||unique"`
	Password string `validate:"required||string=6,20"`
	Email string `validate:"email"`
	LoginCount int
	LastTime time.Time
	LastIp string
	Status int
	Mobile string `validate:"required||string=11,11"`
	HeadImg *HeadImg  `orm:"rel(one)"`  //头像一对一
	Created time.Time
	Updated time.Time
	Role *Role   `orm:"rel(one)"` //等级一对一
}

func (m *User) TableName() string {
	return TableName("user")
}