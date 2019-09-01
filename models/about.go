/*
@Time : 2019-09-01 18:49
@Author : biglight
@File : about
@Software: GoLand
*/
package models

type About struct {
	Id int
	Content string
	Img string
}

func (m *About) TableName() string {
	return TableName("about")
}