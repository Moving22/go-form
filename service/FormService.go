package service

import (
	"go-form/dao"
	. "go-form/models"
)

type FormService struct{}

func (formService *FormService) GetForms(params FormParam) (forms []Form, err error){
	if params.FormName == ""{
		forms, err = dao.QueryForms(params.PageNumber, params.PageSize)
	}else {
		forms, err = dao.QueryFormsByName(params.PageNumber, params.PageSize, params.FormName)
	}
	return forms, err
}

func (formService *FormService) DeleteForm(params FormParam) (error) {
	return dao.DeleteFormById(params.FormId)
}

func (formService *FormService) GetForm(id int, Type int) (form *Form, err error) {
	if Type == 0 {
		form, err = dao.SelectFormById(id)
	}
	if Type == 1 {
		form, err = dao.QueryFormById(id)
	}
	return form, err
}

func (formService *FormService) AddForm(params FormParam, form *Form, Type int) (id int64, err error) {
	if Type == 0 {
		id, err = dao.AddDefaultForm(params.FormName)
	}
	if Type == 1 {
		id, err = dao.AddForm(*form)
	}
	return id, err
}

func (formService *FormService) UpdateForm(id string, updateFormBody UpdateFormBody) error {
	return dao.UpdateFormById(id, updateFormBody)
}
