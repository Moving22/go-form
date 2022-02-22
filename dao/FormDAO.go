package dao

import (
	"fmt"
	"github.com/Masterminds/squirrel"
	. "go-form/models"
	"go-form/utils"
)


//func QueryForms(num,size int) []Form {
//
//}

//func QueryFormsByName(num,size int, name string) []Form {
//
//	return nil
//}
//
//func QueryFormsById(id string) *Form {
//
//}

//按页查找，id降序
func QueryFormsByName(num int, size int, name string) (forms []Form, err error) {
	begin := (num - 1) * size
	end := num * size
	sql := fmt.Sprintf(`select * from test where name='%s' order by id desc limit ?,?`, name)
	err = utils.Db.Select(&forms, sql, begin, end)
	return forms, err
}

func QueryForms(num, size int) (forms []Form, err error) {
	begin := (num - 1) * size
	end := num * size
	sql := "select * from test order by id desc limit ?,?"
	err = utils.Db.Select(&forms, sql, begin, end)
	return forms, err
}

func QueryFormById(id int) (form *Form, err error) {
	sql := "select * from test where id=?"
	err = utils.Db.Get(&form, sql, id) //Get查一个
	if err != nil {
		fmt.Printf("QueryFormById err：%v\n", err)
	}
	return form, err
}

//根据id查找表单
func SelectFormById(id int) (*Form,error) {
	var form Form
	if query, arg, err := squirrel.Select("*").From("test").Where(squirrel.Eq{"id": id}).ToSql(); err == nil {
		if err := utils.Db.QueryRowx(query, arg...).StructScan(&form); err == nil {
			return &form, nil
		} else {
			return nil,err
		}
	}else {
		return nil, err
	}
}

//增加一个默认表单
func AddDefaultForm(name string) (int64, error) {
	if qeuryString, args, err := squirrel.Insert("test").Columns("name", "formdata", "rule").Values(name, "{\"source\": []}", "[]").ToSql(); err == nil {
		res, err := utils.Db.Exec(qeuryString, args...)
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
func AddForm(form Form) (int64, error) {
	if qeuryString, args, err := squirrel.Insert("test").Columns("name", "formdata", "rule").Values(form.Name, form.Formdata, form.Rule).ToSql(); err == nil {
		if res, err := utils.Db.Exec(qeuryString, args...); err != nil {
			fmt.Printf("AddDefaultForm err: %v\n", err)
			return -1, err
		} else {
			id, _ := res.LastInsertId()
			return id, nil
		}
	}
	return 0, nil
}


func UpdateFormById(id string, updateFormBody UpdateFormBody) error {
	query := squirrel.Update("test")
	if updateFormBody.Formdata != nil {
		//if jsonString, err := jsoniter.MarshalToString(updateFormBody.Formdata); err == nil {
			query = query.Set("formdata", updateFormBody.Formdata)
		//}
	}
	if updateFormBody.Rule != nil {
		//if jsonString, err := jsoniter.MarshalToString(updateFormBody.Rule); err == nil {
			query = query.Set("rule", updateFormBody.Rule)
		//}
	}
	if updateFormBody.Name != "" {
		query = query.Set("name", updateFormBody.Name)
	}
	if queryString, arg, err := query.Where(squirrel.Eq{"id": id}).ToSql(); err == nil {
		if _, err := utils.Db.Exec(queryString, arg...); err == nil {
			return nil
		} else {
			return err
		}
	} else {
		return err
	}
}

func DeleteFormById(id int) error {
	if query, arg, err := squirrel.Delete("test").Where(squirrel.Eq{"id": id}).ToSql(); err == nil {
		if _, err := utils.Db.Exec(query, arg...); err == nil{
			return nil
		}else {
			return err
		}
	}else {
		return err
	}
}