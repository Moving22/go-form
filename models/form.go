package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Masterminds/squirrel"
	"go-form/dao"
)

type JSON map[string]interface{}

func (pc *JSON) Scan(val interface{}) error {
	switch v := val.(type) {
	case []byte:
		json.Unmarshal(v, &pc)
		return nil
	case string:
		json.Unmarshal([]byte(v), &pc)
		return nil
	default:
		return errors.New(fmt.Sprintf("Unsupported type: %T", v))
	}
}
func (pc *JSON) Value() (driver.Value, error) {
	return json.Marshal(pc)
}

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

//按页查找，id降序
func QueryPageForm(num int, size int, name string) ([]Form, error) {
	var rforms []Form
	var sql string
	if name != "" {
		sql = fmt.Sprintf(`select * from test where name=%s order by id desc limit ?,?`, name)
	} else {
		sql = "select * from test order by id desc limit ?,?"
	}

	begin := (num - 1) * size
	end := num * size
	err := dao.Db.Select(&rforms, sql, begin, end) //Select查多个
	if err != nil {
		fmt.Printf("QueryPageForm err：%v\n", err)
		return nil, err
	}
	return rforms, nil
}

//按id查看详情
func QueryFormById(id string) (*Form, error) {
	var rform Form
	sql := "select * from test where id=?"
	err := dao.Db.Get(&rform, sql, id) //Get查一个
	if err != nil {
		fmt.Printf("QueryFormById err：%v\n", err)
		return nil, err
	}
	//解析原生表单的JSON字符串
	var form Form
	form.Id = rform.Id
	form.Name = rform.Name
	return &form, nil
}

//增加一个默认表单
func AddDefaultForm(name string) (int64, error) {
	if qeuryString, args, err := squirrel.Insert("test").Columns("name", "formdata", "rule").Values(name, "{\"source\": []}", "[]").ToSql(); err == nil {
		res, err := dao.Db.Exec(qeuryString, args...)
		if err != nil {
			fmt.Printf("AddDefaultForm err: %v\n", err)
			return -1, err
		}
		id, _ := res.LastInsertId()
		return id, nil
	}
	return 0, nil
}

//增加一个表单
func AddForm(form *Form) (int64, error) {
	if qeuryString, args, err := squirrel.Insert("test").Columns("name", "formdata", "rule").Values(form.Name, form.Formdata, form.Rule).ToSql(); err == nil {
		if res, err := dao.Db.Exec(qeuryString, args...); err != nil {
			fmt.Printf("AddDefaultForm err: %v\n", err)
			return -1, err
		} else {
			id, _ := res.LastInsertId()
			return id, nil
		}
	}
	return 0, nil
}
