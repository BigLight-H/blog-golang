/*
@Time : 2019-09-01 13:24
@Author : biglight
@File : tags
@Software: GoLand
*/
package models

type Tags struct {
	Id int
	TagName string
}

func (m *Tags) TableName() string {
	return TableName("tags")
}
