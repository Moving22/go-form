package service

import (
	"go-form/dao"
	. "go-form/models"
)

/**
	感觉结构太简单了，分service显得很刻意，直接用dao了
 */

func QueryForms(num int, size int, name string) ([]Form, error){
	forms, err := dao.QueryPageForm(num, size, name)
	if err != nil {
		panic(err)
		return nil, err
	}
	return forms, nil
}