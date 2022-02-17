package models

import (
	"encoding/json"
	"fmt"
	"go-form/dao"
)

type BaseFormInfo struct {
	Id		int	   `json:"id"`
	Name 	string `json:"name"`
}

type RawForm struct {
	BaseFormInfo
	Data 	string `json:"formdata" db:"formdata"`	//db:sqlx模块中对应的数据库字段名
	Rule 	string `json:"rule" db:"rule"`
}

type Form struct {
	BaseFormInfo
	Data 	map[string]interface{} `json:"formdata"`	//formdata是{}
	Rule 	[]map[string]interface{} `json:"rule"`		//rule是[]
}


//按页查找，id降序
func QueryPageForm(num int, size int, name string) ([]Form,error){
	var rforms []RawForm
	var sql string
	if name !="" {
		sql = fmt.Sprintf(`select * from test where name=%s order by id desc limit ?,?`, name)
	}else {
		sql = "select * from test order by id desc limit ?,?"
	}

	begin := (num-1) * size
	end := num * size
	err := dao.Db.Select(&rforms, sql, begin, end)			//Select查多个
	if err != nil {
		fmt.Printf("QueryPageForm err：%v\n",err)
		return nil, err
	}

	//解析原生表单的JSON字符串
	size = len(rforms)
	forms := make([]Form, size)
	for i := 0; i < size; i++ {
		forms[i].Id = rforms[i].Id
		forms[i].Name = rforms[i].Name

		err := json.Unmarshal([]byte(rforms[i].Data), &forms[i].Data)
		if err != nil {
			fmt.Printf("QueryPageForm parse data err：%v\n",err)
			return nil, err
		}
		err = json.Unmarshal([]byte(rforms[i].Rule), &forms[i].Rule)
		if err != nil {
			fmt.Printf("QueryPageForm parse rule err：%v\n",err)
			return nil, err
		}
	}

	return forms, nil
}


//按名称查找
func QueryFormByName()  {

}


//按id查看详情
func QueryFormById(id string) (*Form, error){
	var rform RawForm
	sql := "select * from test where id=?"
	err := dao.Db.Get(&rform, sql, id)			//Get查一个
	if err != nil {
		fmt.Printf("QueryFormById err：%v\n",err)
		return nil, err
	}

	//解析原生表单的JSON字符串
	var form Form
	form.Id = rform.Id
	form.Name = rform.Name

	err = json.Unmarshal([]byte(rform.Data), &form.Data)
	if err != nil {
		fmt.Printf("QueryPageForm parse err：%v\n",err)
		return nil, err
	}
	err = json.Unmarshal([]byte(rform.Rule), &form.Rule)
	if err != nil {
		fmt.Printf("QueryPageForm parse err：%v\n",err)
		return nil, err
	}

	return &form, nil
}


//更新col字段的数据
func UpdateFormData(id, col, data string) error {
	sql := fmt.Sprintf("update test set %s=? where id=?",col)
	_, err := dao.Db.Exec(sql, data, id)
	if err != nil {
		fmt.Printf("UpdateFormData err: %v\n",err)
		return err
	}

	return nil
}


//删除
func DeleteForm(id string) error {
	sql := "delete from test where id=?"
	_,err := dao.Db.Exec(sql, id)
	if err != nil {
		fmt.Printf("UpdateFormData err: %v\n",err)
		return err
	}
	return nil
}


//增加一个默认表单
func AddDefaultForm(name string) (int64,error) {
	sql := "insert into test(name,formdata,rule) values(?,?,?)"
	formdata := `{"source": []}`
	rule := "[]"

	res,err := dao.Db.Exec(sql, name, formdata, rule)
	if err != nil {
		fmt.Printf("AddDefaultForm err: %v\n",err)
		return -1,err
	}

	id,_ := res.LastInsertId()
	return id,nil
}