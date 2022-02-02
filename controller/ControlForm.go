package controller

import (
	"github.com/gin-gonic/gin"
	mod "go-form/models"
	"strconv"
)

//获得一页表单
func GetPageForm(c *gin.Context)  {
	//获取pageNumber和pageSize
	pageNum ,_ := strconv.Atoi(c.DefaultQuery("pageNumber","1"))
	pageSize ,_ := strconv.Atoi(c.DefaultQuery("pageSize","10"))

	//从mysql查询
	forms := mod.QueryPageForm(pageNum, pageSize)

	//响应数据
	c.JSON(200, gin.H{
		"code" : 200,
		//"msg" : "GetPageForm",
		"data" : forms,
		"pageSize" : pageSize,
		"pageNumber" : pageNum,
		"totalNumber" : pageNum * pageSize,
	})
}

//获得某表单
func GetForm(c *gin.Context)  {
	c.JSON(200, gin.H{
		"msg":"GetForm",
	})
}

//更新表单
func UpdateForm(c *gin.Context)  {
	c.JSON(200, gin.H{
		"msg":"UpdateForm",
	})
}

//删除表单
func DeleteForm(c *gin.Context)  {
	c.JSON(200, gin.H{
		"msg":"DeleteForm",
	})
}

//创建表单
func CreateForm(c *gin.Context)  {
	c.JSON(200, gin.H{
		"msg":"CreateForm",
	})
}