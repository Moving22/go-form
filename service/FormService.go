package service

import (
	"go-form/dao"
	. "go-form/models"
)

type FormService struct{}

func (formService *FormService) GetForms(num int, size int, name string) (forms []Form, err error){
	if name == ""{
		forms, err = dao.QueryForms(num, size)
	}else {
		forms, err = dao.QueryFormsByName(num, size, name)
	}
	return forms, err
}

func (formService *FormService) DeleteForm(id string) (error) {
	return dao.DeleteFormById(id)
}

func (formService *FormService) GetForm(id string, Type int) (form *Form, err error) {
	if Type == 0 {
		form, err = dao.SelectFormById(id)
	}
	if Type == 1 {
		form, err = dao.QueryFormById(id)
	}
	return form, err
}

func (formService *FormService) AddForm(name string, form *Form, Type int) (id int64, err error) {
	if Type == 0 {
		id, err = dao.AddDefaultForm(name)
	}
	if Type == 1 {
		id, err = dao.AddForm(*form)
	}
	return id, err
}

func (formService *FormService) UpdateForm(id string, updateFormBody UpdateFormBody) error {
	return dao.UpdateFormById(id, updateFormBody)
}
