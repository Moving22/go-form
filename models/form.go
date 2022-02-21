package models

import (
	."go-form/utils"
)


type BaseFormInfo struct {
	Id   int    `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
}

type RawForm struct {
	BaseFormInfo
	Data JSON `json:"formdata" db:"formdata"` //db:sqlx模块中对应的数据库字段名
	Rule JSON `json:"rule" db:"rule"`
}

type Form struct {
	BaseFormInfo
	Formdata JSON `db:"formdata" json:"formdata"` //formdata是{}
	Rule     JSON `db:"rule" json:"rule"`         //rule是[]
}


type UpdateFormBody struct {
	Formdata interface{} `json:"formdata" `
	Rule     interface{} `json:"rule"`
	Name     string      `json:"name" `
}


type CopyFormBody struct {
	FormId string
}