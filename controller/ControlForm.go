package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-form/dao"
	"go-form/lib"
	. "go-form/models"
	"go-form/service"
	"net/http"
	"strconv"
)

//获得一页表单
func GetPageForm(c *gin.Context) {
	//获取pageNumber和pageSize
	pageNum, _ := strconv.Atoi(c.DefaultQuery("pageNumber", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))
	name := c.Query("name")

	//获取数据库中的数据
	forms, err := service.QueryForms(pageNum, pageSize, name)

	//响应数据
	if err == nil {
		c.JSON(http.StatusOK, lib.Success(forms))
	} else {
		c.JSON(http.StatusInternalServerError, lib.Fail(1,"QueryPageForm出错"))
	}
}

//获得某表单
func GetForm(c *gin.Context) {
	id := c.Param("id")
	if form,err := dao.SelectFormById(id); err==nil{
		c.JSON(http.StatusOK, lib.Success(form))
	}else {
		fmt.Printf("err: %v", err)
		c.JSON(http.StatusInternalServerError, lib.Fail(1,"GetForm错误"))
	}
}


//编辑页面的保存表单，更新，注：？
func UpdateForm(c *gin.Context) {
	var updateFormBody UpdateFormBody
	id := c.Param("id")
	if err := c.ShouldBindJSON(&updateFormBody); err == nil {
		if err := dao.UpdateFormById(id, updateFormBody); err == nil {
			c.JSON(http.StatusOK, lib.Success(nil))
		} else {
			fmt.Printf("%+v\n",err)
			c.JSON(http.StatusInternalServerError, lib.Fail(1, "UpdateFormById err"))
		}
	} else {
		fmt.Printf("%+v\n",err)
		c.JSON(http.StatusInternalServerError, lib.Fail(1,"参数有误"))
	}
}


//删除表单
func DeleteForm(c *gin.Context) {
	id := c.Param("id")
	if err := dao.DeleteFormById(id); err != nil {
			c.JSON(http.StatusInternalServerError, lib.Fail(1,"DeleteForm错误"))
	} else {
		c.JSON(http.StatusOK, lib.Success(nil))
	}
}


//创建表单，添加一个默认表单，再返回该表单
func CopyForm(c *gin.Context) {
	var copyFormBody CopyFormBody
	if err := c.ShouldBindJSON(&copyFormBody); err == nil { //更优的获取json参数，ShouldBlindJSON 只绑定结构体tag有binding:required的参数
		if record, err := dao.QueryFormById(copyFormBody.FormId); err != nil {
			if newFormId, err := dao.AddForm(record); err != nil {
				record.Id = int(newFormId)
				c.JSON(http.StatusOK, lib.Success(record))
			} else {
				c.JSON(http.StatusOK, lib.Fail(1, err.Error()))
			}
		} else {
			c.JSON(http.StatusOK, lib.Fail(1,"没有该表单项"))
		}
	} else {
		c.JSON(http.StatusInternalServerError, lib.Fail(1,"formid错误"))
	}

}

//创建表单，添加一个默认表单，再返回该表单
func CreateForm(c *gin.Context) {
	var m map[string]string
	if err := c.ShouldBindJSON(&m); err == nil { //更优的获取json参数，ShouldBlindJSON 只绑定结构体tag有binding:required的参数
		if id, err := dao.AddDefaultForm(m["name"]); err != nil {
			c.JSON(http.StatusInternalServerError, lib.Fail(1, "AddDefaultForm err"))
		} else {
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"id":   id,
			})
		}
	} else {
		c.JSON(http.StatusInternalServerError, lib.Fail(1, "AddDefaultForm err"))
	}

}
