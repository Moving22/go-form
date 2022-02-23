package service

import (
	"encoding/json"
	"fmt"
	."go-form/models"
	."go-form/utils"
)

type UserDataService struct{}

func (u *UserDataService)GetTableCount(params TableData) (res []int,err error) {
	var sql string
	if params.ColumnValue == ""{
		sql = fmt.Sprintf(`SELECT id FROM user_data WHERE table_key=%s AND status='1' AND JSON_CONTAINS_PATH(data,'one','$.%s');`,params.TableKey,params.ColumnHead)
	}else {
		sql = fmt.Sprintf(`SELECT id FROM user_data WHERE table_key=%s AND status='1' AND JSON_CONTAINS(data,'%s','$.%s');`,params.TableKey,params.ColumnValue,params.ColumnHead)
	}
	err = Db.Select(&res, sql)
	return res, err
}

func (u *UserDataService) GetDataIds(params UserData) (res []int,err error) {
	sql := `select id from user_data where userId=? and orderNumber=? and status=1`
	err = Db.Select(&res, sql, params.UserId, params.OrderNumber)
	return res, err
}

func (u *UserDataService) GetUserData(params UserData) (res UserData,err error) {
	sql := `select * from user_data where table_key=? and userId=? and orderNumber=? and status=1`
	err = Db.Get(&res, sql, params.TableKey, params.UserId, params.OrderNumber)
	return res, err
}

func (u *UserDataService) UpdateById(id string, params UserData) error {
	data,_ := json.Marshal(params.Data)
	sql := `update user_data set data=? where id=? and status=1`
	_,err := Db.Exec(sql, string(data), id)
	return err
}

func (u *UserDataService) Delete(params UserData) error {
	sql := `delete from user_data where id=? and status=1`
	 _,err := Db.Exec(sql,params.Id)
	return err
}