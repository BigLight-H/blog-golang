/*
@Time : 2019-08-24 14:49
@Author : biglight
@File : type
@Software: GoLand
*/
package models

type Type struct {
	Id int `orm:"auto"`
	Pid int
	TName string
	Url string
	Dir int
	Icon string
	Status int
}

func (m *Type) TableName() string {
	return TableName("type")
}
