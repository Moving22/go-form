package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-form/lib"
	. "go-form/models"
	"go-form/service"
	"go-form/utils"
	"net/http"
)

var userDataService service.UserDataService

func GetTestField(c *gin.Context)  {
	var tableData TableData
	if err := c.ShouldBind(&tableData); err ==nil{
		if cnt, err := userDataService.GetTableCount(tableData); err == nil{
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
		if res, err := userDataService.GetDataIds(userData); err == nil{
			if res == nil {
				c.JSON(http.StatusOK, lib.Fail(0,"未查到"))
			}else {
				c.JSON(http.StatusOK, lib.Success(utils.JSON{"id":res}))
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
		if data, err := userDataService.GetUserData(userTable); err == nil{
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
	var params UserData
	if err := c.ShouldBind(&params); err == nil{
		fmt.Printf("%s\n", params.Data)
		if err := userDataService.UpdateById(id, params); err == nil {
			c.JSON(http.StatusOK, lib.Success(nil))
		}else {
			fmt.Printf("%+v\n",err)
			c.JSON(http.StatusInternalServerError, lib.Fail(1,"更新失败"))
		}
	}else {
		c.JSON(http.StatusBadRequest, lib.Fail(1,"参数错误"))
	}

}


func DelFormField(c *gin.Context)  {
	var params UserData
	if err := c.ShouldBindUri(&params); err == nil{
		if err := userDataService.Delete(params); err == nil{
			c.JSON(http.StatusOK, lib.Success(nil))
		}else {
			c.JSON(http.StatusInternalServerError, lib.Fail(1,"删除失败"))
		}
	}else {
		c.JSON(http.StatusBadRequest, lib.Fail(1,"参数错误"))
	}

}

