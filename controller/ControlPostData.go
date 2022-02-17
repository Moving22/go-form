package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"go-form/dao"
	mod "go-form/models"
	."go-form/utils"
)


func GetTestField(c *gin.Context)  {
	params := c.Request.URL.Query()

	var sql string
	if params["column_value"] == nil{
		sql = fmt.Sprintf(`SELECT id FROM user_data WHERE table_key=%s AND status='1' AND JSON_CONTAINS_PATH(data,'one','$.%s');`, params["table_key"][0],params["column_head"][0])
	}else {
		sql = fmt.Sprintf(`SELECT id FROM user_data WHERE table_key=%s AND status='1' AND JSON_CONTAINS(data,'%s','$.%s');`,params["table_key"][0],params["column_value"][0],params["column_head"][0])
	}

	var res []int
	err := dao.Db.Select(&res, sql)

	if err != nil {
		c.JSON(200, gin.H{
			"code" : 1,
			"msg" : "TestField err",
		})
	}else {
		m := make(map[string]int)
		m["result"] = len(res)
		c.JSON(200, gin.H{
			"code" : 0,
			"data" : m,
		})
	}
}


func GetFormField(c *gin.Context)  {
	pam := c.Request.URL.Query()

	sql := `select id from user_data where userId=? and orderNumber=? and status=1`

	var res []int
	dao.Db.Select(&res, sql, pam["userId"][0], pam["orderNumber"][0])
	if res == nil {
		c.JSON(500, gin.H{
			"code" : 1,
		})
	}else {
		c.JSON(200, gin.H{
			"code" : 0,
		})
	}
}


func GetDataFiled(c *gin.Context)  {
	pam := c.Request.URL.Query()

	sql := `select * from user_data where table_key=? and userId=? and orderNumber=? and status=1`

	var data mod.UserData
	err := dao.Db.Get(&data, sql, pam["table_key"][0],pam["userId"][0],pam["orderNumber"][0])

	var m map[string]interface{}
	json.Unmarshal([]byte(data.Data), &m)
	if err == nil{
		c.JSON(200, gin.H{
			"code" : 0,
			"data" : m,
		})
	}else {
		fmt.Println(err)
		c.JSON(500, gin.H{
			"code" : 1,
			"msg" : "系统错误",
		})
	}
}


func PutFormField(c *gin.Context)  {
	id := c.Param("id")
	var m map[string]interface{}
	c.BindJSON(&m)

	data,_ := json.Marshal(m)
	sql := `update user_data set data=? where id=? and status=1`
	_,err := dao.Db.Exec(sql, string(data), id)

	if err == nil {
		c.JSON(200, gin.H{
			"code" : 0,
			"msg" : "更新成功",
		})
	}else {
		c.JSON(500, gin.H{
			"code" : 1,
			"msg" : "更新错误",
		})
	}

}


func DelFormField(c *gin.Context)  {
	id := c.Param("id")
	sql := `delete from user_data where id=? and status=1`
	_,err := dao.Db.Exec(sql,id)
	if err != nil {
		c.JSON(500, gin.H{
			"code" : 1,
			"msg" : "删除失败",
		})
	}else {
		c.JSON(200, gin.H{
			"code" : 0,
			"msg" : "删除成功",
		})
	}
}

//废弃
func PostFormField(c *gin.Context)  {
	m := Map{
		"data" : make(Map),
		"orderNumber" : "testorder",
	}	//设置请求参数的默认值
	c.BindJSON(&m)

	//sql := `update user_data set status=0 where userId=? and orderNumber=?`
	//dao.Db.Exec(sql, m["userId"], m["orderNumber"])

	//authorization := fmt.Sprintf("Bearer %s", m["token"])



	delivery_address := m["data"].(map[string]interface{})["delivery_addresses"].(map[string]interface{})	//注：转换时不能用tpye的Map?
	for k,v := range delivery_address{
		if v == nil {
			delete(delivery_address, k)
		}
	}
	personInfo := m["data"].(map[string]interface{})["personInfo"].(map[string]interface{})
	for k,v := range personInfo{
		if v == nil {
			delete(personInfo, k)
		}
	}


	c.JSON(200, gin.H{
		"ret":delivery_address,
	})

}
