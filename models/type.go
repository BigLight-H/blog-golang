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
	//Article []*Article  `orm:"reverse(many)"` // 设置一对多的反向关系
}

func (m *Type) TableName() string {
	return TableName("type")
}
