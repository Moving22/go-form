package models

import (
	"fmt"
	"go-form/dao"
)


type Form struct {
	Id		int	   `json:"id"`
	Name 	string `json:"name"`
	Data 	string `json:"formdata" db:"formdata"`	//db:sqlx模块中对应的数据库字段名
	Rule 	string `json:"rule"`
}

//功能测试
func TestQuery() []Form{
	var forms []Form
	sql := "select * from test limit 1,5"
	err := dao.Db.Select(&forms, sql)
	if err != nil {
		fmt.Printf("select err：%v\n",err)
		return nil
	}
	return forms
}

//查找一页，id降序
func QueryPageForm(num, size int) ([]Form){
	var forms []Form
	sql := "select * from test order by id desc limit ?,?"
	begin := (num-1) * size
	end := num * size
	err := dao.Db.Select(&forms, sql, begin, end)
	if err != nil {
		fmt.Printf("QueryPageForm err：%v\n",err)
		return nil
	}
	return forms
}

//查找一个
func QueryForm()  {

}
