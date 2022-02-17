package models

type UserData struct {
	Id 			int `json:"id" db:"id"`
	TableKey 	int `json:"table_key" db:"table_key"`
	Data		string `json:"data" db:"data"`
	UserId		int `json:"userId" db:"userId"`
	CourseId	int `json:"courseId" db:"courseId"`
	OrderNumber	string `json:"orderNumber" db:"orderNumber"`
	CreateTime	string `json:"createTime" db:"createTime"`
	Status		byte `json:"status" db:"status"`
	Type 		byte `json:"type" db:"type"`
}