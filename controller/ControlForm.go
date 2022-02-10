package controller

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	mod "go-form/models"
	"strconv"
)

//获得一页表单
func GetPageForm(c *gin.Context)  {
	//获取pageNumber和pageSize
	pageNum ,_ := strconv.Atoi(c.DefaultQuery("pageNumber","1"))
	pageSize ,_ := strconv.Atoi(c.DefaultQuery("pageSize","10"))

	//获取数据库中的数据
	forms, err := mod.QueryPageForm(pageNum, pageSize)

	//响应数据
	if err == nil {
		c.JSON(200, gin.H{
			"code" : 0,
			"data" : forms,
			"pageSize" : pageSize,
			"pageNumber" : pageNum,
			"totalNumber" : len(forms),
		})
	} else {
		c.JSON(500, gin.H{
			"code" : 1,
			"msg" : "QueryPageForm获取数据时出错",
		})
	}
}


//获得某表单
func GetForm(c *gin.Context)  {
	id := c.Param("id")

	forms, err := mod.QueryFormById(id)

	if err != nil {
		c.JSON(500, gin.H{
			"code" : 1,
			"msg" : "QueryFormById出错",
		})
	}else {
		c.JSON(200, gin.H{
			"code" : 0,
			"data" : forms,
		})
	}
}


//编辑页面的保存表单，更新
func UpdateForm(c *gin.Context)  {
	id := c.Param("id")
	bytes,err := c.GetRawData()		//获取request.body参数

	var m map[string]interface{}
	json.Unmarshal(bytes, &m)	//字节转json

	data, _ := json.Marshal(m["formdata"])	//取出json中的数据，转为字节
	if data != nil{
		err = mod.UpdateFormData(id, "formdata", string(data))
	}
	data, _ = json.Marshal(m["rule"])
	if data != nil{
		err = mod.UpdateFormData(id, "rule", string(data))
	}
	data, _ = json.Marshal(m["name"])
	if data != nil{
		err = mod.UpdateFormData(id, "name", string(data))
	}

	if err != nil {
		c.JSON(500, gin.H{
			"code":1,
			"msg":"UpdateForm更新出错",
		})
	}else {
		c.JSON(200, gin.H{
			"code":0,
		})
	}
}


//删除表单
func DeleteForm(c *gin.Context)  {
	id := c.Param("id")

	err := mod.DeleteForm(id)

	if err != nil {
		c.JSON(500, gin.H{
			"code":1,
			"msg":"DeleteForm err",
		})
	}else {
		c.JSON(200, gin.H{
			"code":0,
		})
	}
}


//创建表单，添加一个默认表单，再返回该表单
func CreateForm(c *gin.Context)  {
	var m map[string]string
	c.BindJSON(&m)			//更优的获取json参数

	id,err := mod.AddDefaultForm(m["name"])

	if err != nil {
		c.JSON(500, gin.H{
			"code" : 1,
			"msg" : "AddDefaultForm err",
		})
	}else {
		c.JSON(200, gin.H{
			"code" : 0,
			"id" : id,
		})
	}
}