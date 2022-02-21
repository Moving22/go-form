package dao

import (
	"fmt"
	. "go-form/models"
	. "go-form/utils"
)

func GetTableCount(table TableData) ([]int,error) {
	var sql string
	if table.ColumnValue == ""{
		sql = fmt.Sprintf(`SELECT id FROM user_data WHERE table_key=%s AND status='1' AND JSON_CONTAINS_PATH(data,'one','$.%s');`,table.TableKey,table.ColumnHead)
	}else {
		sql = fmt.Sprintf(`SELECT id FROM user_data WHERE table_key=%s AND status='1' AND JSON_CONTAINS(data,'%s','$.%s');`,table.TableKey,table.ColumnValue,table.ColumnHead)
	}
	var res []int
	if err := Db.Select(&res, sql); err == nil{
		return res, nil
	}else {
		return res, err
	}
}

func GetIds(userData UserData) ([]int,error) {
	sql := `select id from user_data where userId=? and orderNumber=? and status=1`

	var res []int
	if err := Db.Select(&res, sql, userData.UserId, userData.OrderNumber); err == nil{
		return res, nil
	}else {
		return nil, err
	}
}

func GetUserData(data UserData) (*UserData,error) {
	var res UserData
	sql := `select * from user_data where table_key=? and userId=? and orderNumber=? and status=1`
	if err := Db.QueryRowx(sql, data.TableKey, data.UserId, data.OrderNumber).StructScan(&res); err == nil {
		return &res, nil
	} else {
		return nil,err
	}
}

func DeleteUserDataById(id string) error {
	sql := `delete from user_data where id=? and status=1`
	if _,err := Db.Exec(sql,id); err == nil{
		return nil
	}else {
		return err
	}
}

func UpdateUserDataById(id, data string) error {
	sql := `update user_data set data=? where id=? and status=1`
	if _,err := Db.Exec(sql, string(data), id); err == nil{
		return nil
	}else {
		return err
	}
}