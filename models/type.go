/*
@Time : 2019-08-24 14:49
@Author : biglight
@File : type
@Software: GoLand
*/
package models

type Type struct {
	Id int `orm:"auto"`
	Pid string
	TName string
	Status int
}

func (m *Type) TableName() string {
	return TableName("type")
}
