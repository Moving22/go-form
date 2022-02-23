package models

import "go-form/utils"

type UserData struct {
	Id 			int `json:"id" db:"id" form:"id" uri:"id"`
	TableKey 	int `json:"table_key" db:"table_key" form:"table_key"`
	Data		utils.JSON `json:"data" db:"data" form:"data"`
	UserId		int `json:"userId" db:"userId" form:"userId"`
	CourseId	int `json:"courseId" db:"courseId"`
	OrderNumber	string `json:"orderNumber" db:"orderNumber" form:"orderNumber"`// GET中的·ShouldBind·需要·form·
	CreateTime	string `json:"createTime" db:"createTime"`
	Status		byte `json:"status" db:"status"`
	Type 		byte `json:"type" db:"type"`
}

type TableData struct {
	TableKey 	string `json:"table_key" db:"table_key" form:"table_key"`
	ColumnHead 	string `json:"column_head" db:"table_key" form:"column_head"`
	ColumnValue string `json:"column_value" db:"table_key" form:"column_value"`
}

