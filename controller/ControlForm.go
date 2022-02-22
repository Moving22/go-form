package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-form/dao"
	"go-form/lib"
	. "go-form/models"
	"go-form/service"
	"net/http"
)


var formService service.FormService

//获得一页表单
func GetPageForm(c *gin.Context) {
	params := FormParam{
		PageNumber: 1,
		PageSize: 10,
	}
	if err := c.ShouldBind(&params); err == nil{
		if forms, err := formService.GetForms(params); err == nil{
			c.JSON(http.StatusOK, lib.Success(forms))
		} else {
			c.JSON(http.StatusInternalServerError, lib.Fail(1,"QueryPageForm出错"))
		}
	}else {
		c.JSON(http.StatusBadRequest, lib.Fail(1,"参数传递错误"))
	}
}

//获得某表单
func GetForm(c *gin.Context) {
	var params FormParam
	if err := c.ShouldBind(&params); err == nil{
		if form,err := formService.GetForm(params.FormId, 0); err==nil{
			c.JSON(http.StatusOK, lib.Success(form))
		}else {
			fmt.Printf("err: %v", err)
			c.JSON(http.StatusInternalServerError, lib.Fail(1,"GetForm错误"))
		}
	}else {
		c.JSON(http.StatusBadRequest, lib.Fail(1,"参数传递错误"))
	}
}


//编辑页面的保存表单，更新，注: 接收参数待修改
func UpdateForm(c *gin.Context) {
	var updateFormBody UpdateFormBody
	id := c.Param("id")
	if err := c.ShouldBindJSON(&updateFormBody); err == nil {
		if err := formService.UpdateForm(id, updateFormBody); err == nil {
			c.JSON(http.StatusOK, lib.Success(nil))
		} else {
			c.JSON(http.StatusInternalServerError, lib.Fail(1, "UpdateFormById err"))
		}
	} else {
		c.JSON(http.StatusBadRequest, lib.Fail(1,"参数传递错误"))
	}
}


//删除表单
func DeleteForm(c *gin.Context) {
	var params FormParam
	if err := c.ShouldBindUri(&params); err == nil{
		fmt.Printf("%+v\n", params)
		if err := formService.DeleteForm(params); err != nil {
			c.JSON(http.StatusInternalServerError, lib.Fail(1,"DeleteForm错误"))
		} else {
			c.JSON(http.StatusOK, lib.Success(nil))
		}
	}else {
		c.JSON(http.StatusBadRequest, lib.Fail(1,"参数传递错误"))
	}
}


//创建表单，添加一个默认表单，再返回该表单
func CopyForm(c *gin.Context) {
	var copyFormBody CopyFormBody
	if err := c.ShouldBindJSON(&copyFormBody); err == nil { //更优的获取json参数，ShouldBlindJSON 只绑定结构体tag有binding:required的参数
		if record, err := formService.GetForm(copyFormBody.FormId, 1); err != nil {
			if newFormId, err := dao.AddForm(*record); err != nil {
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
	var params FormParam
	if err := c.ShouldBindJSON(&params); err == nil { //更优的获取json参数，ShouldBlindJSON 只绑定结构体tag有binding:required的参数
		if id, err := formService.AddForm(params, nil, 0); err != nil {
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
