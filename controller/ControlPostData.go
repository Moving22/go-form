package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"go-form/dao"
	"go-form/lib"
	. "go-form/models"
	. "go-form/utils"
	"net/http"
)



func GetTestField(c *gin.Context)  {
	var tableData TableData
	if err := c.ShouldBind(&tableData); err ==nil{
		if cnt, err := dao.GetTableCount(tableData); err == nil{
			m := make(map[string]int)
			m["result"] = len(cnt)
			c.JSON(http.StatusOK, lib.Success(m))
		}else {
			fmt.Printf("%+v\n",err)
			c.JSON(http.StatusInternalServerError, lib.Fail(1,"GetTableCount err"))
		}
	}else {
		c.JSON(http.StatusInternalServerError, lib.Fail(1,"TestField parameter err"))
	}
}


func GetFormField(c *gin.Context)  {
	var userData UserData
	if err := c.ShouldBind(&userData); err == nil{
		if res, err := dao.GetIds(userData); err == nil{
			if res == nil {
				c.JSON(http.StatusOK, lib.Fail(1,"未查到"))
			}else {
				c.JSON(http.StatusOK, lib.Success(res))
			}
		}else {
			c.JSON(http.StatusInternalServerError, lib.Fail(1,"GetIds err"))
		}
	}else {
		c.JSON(http.StatusInternalServerError, lib.Fail(1,"获取参数错误"))
	}
}


func GetDataFiled(c *gin.Context)  {
	var userTable UserData
	if err := c.ShouldBind(&userTable); err == nil{
		if data, err := dao.GetUserData(userTable); err == nil{
			c.JSON(http.StatusOK, lib.Success(data))
		}else {
			fmt.Printf("%+v\n",err)
			c.JSON(http.StatusInternalServerError, lib.Fail(1,"GetUserData错误"))
		}
	}else {
		fmt.Printf("%+v\n",err)
		c.JSON(http.StatusInternalServerError, lib.Fail(1,"参数获取错误"))
	}
}


func PutFormField(c *gin.Context)  {
	id := c.Param("id")
	var m JSON
	c.BindJSON(&m)
	data,_ := json.Marshal(m)
	if err := dao.UpdateUserDataById(id, string(data)); err == nil {
		c.JSON(http.StatusOK, lib.Success(nil))
	}else {
		fmt.Printf("%+v\n",err)
		c.JSON(http.StatusInternalServerError, lib.Fail(1,"更新失败"))
	}
}


func DelFormField(c *gin.Context)  {
	id := c.Param("id")
	if err := dao.DeleteUserDataById(id); err == nil{
		c.JSON(http.StatusOK, lib.Success(nil))
	}else {
		c.JSON(http.StatusInternalServerError, lib.Fail(1,"删除失败"))
	}
}

