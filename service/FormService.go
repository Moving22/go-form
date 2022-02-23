package service

import (
	"fmt"
	"github.com/Masterminds/squirrel"
	. "go-form/models"
	"go-form/utils"
)

type FormService struct{}

func (formService *FormService) GetForms(params FormParam) (forms []Form, err error){
	if params.FormName == ""{
		forms, err = formService.QueryForms(params.PageNumber, params.PageSize)
	}else {
		forms, err = formService.QueryFormsByName(params.PageNumber, params.PageSize, params.FormName)
	}
	return forms, err
}

func (formService *FormService) DeleteForm(params FormParam) (error) {
	if query, arg, err := squirrel.Delete("test").Where(squirrel.Eq{"id": params.FormId}).ToSql(); err == nil {
		_, err := utils.Db.Exec(query, arg...)
		return err
	}else {
		return err
	}
}

func (formService *FormService) GetForm(id int, Type int) (form *Form, err error) {
	if Type == 0 {
		form, err = formService.SelectFormById(id)
	}
	if Type == 1 {
		form, err = formService.QueryFormById(id)
	}
	return form, err
}

func (formService *FormService) AddForm(params FormParam, form *Form, Type int) (id int64, err error) {
	if Type == 0 {
		id, err = formService.AddDefaultForm(params.FormName)
	}
	if Type == 1 {
		id, err = formService.AddFormByForm(*form)
	}
	return id, err
}

func (formService *FormService) UpdateForm(id string, updateFormBody UpdateFormBody) error {
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
		_, err := utils.Db.Exec(queryString, arg...)
		return err
	} else {
		return err
	}
}

/*
 *	下方是一些sql语句的封装
 */
func (formService *FormService) QueryForms(num, size int) (forms []Form, err error)  {
	begin := (num - 1) * size
	end := num * size
	sql := "select * from test order by id desc limit ?,?"
	err = utils.Db.Select(&forms, sql, begin, end)
	return forms, err
}

func (formService *FormService) QueryFormsByName(num int, size int, name string) (forms []Form, err error) {
	begin := (num - 1) * size
	end := num * size
	sql := fmt.Sprintf(`select * from test where name='%s' order by id desc limit ?,?`, name)
	err = utils.Db.Select(&forms, sql, begin, end)
	return forms, err
}

func (formService *FormService) AddDefaultForm(name string) (int64, error) {
	if qeuryString, args, err := squirrel.Insert("test").Columns("name", "formdata", "rule").Values(name, "{\"source\": []}", "[]").ToSql(); err == nil {
		res, err := utils.Db.Exec(qeuryString, args...)
		if err != nil {
			return -1, err
		}
		id, _ := res.LastInsertId()
		return id, nil
	}
	return -1, nil
}

func (formService *FormService) AddFormByForm(form Form) (int64, error) {
	if qeuryString, args, err := squirrel.Insert("test").Columns("name", "formdata", "rule").Values(form.Name, form.Formdata, form.Rule).ToSql(); err == nil {
		if res, err := utils.Db.Exec(qeuryString, args...); err != nil {
			fmt.Printf("AddDefaultForm err: %v\n", err)
			return -1, err
		} else {
			id, _ := res.LastInsertId()
			return id, nil
		}
	}
	return -1, nil
}

func (formService *FormService) SelectFormById(id int) (*Form,error) {
	var form Form
	if query, arg, err := squirrel.Select("*").From("test").Where(squirrel.Eq{"id": id}).ToSql(); err == nil {
		err := utils.Db.QueryRowx(query, arg...).StructScan(&form)
		return &form,err
	}else {
		return &form, err
	}
}

func (formService *FormService) QueryFormById(id int) (form *Form, err error) {
	sql := "select * from test where id=?"
	err = utils.Db.Get(&form, sql, id) //Get查一个
	return form, err
}
