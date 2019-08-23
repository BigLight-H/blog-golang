/*
@Time : 2019-08-23 20:59
@Author : biglight
@File : role
@Software: GoLand
*/
package models

type Role struct {
	Id    int
	Role  string
}

func (m *Role) TableName() string {
	return TableName("role")
}
