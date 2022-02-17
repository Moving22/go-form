package controller

import (
	"github.com/Masterminds/squirrel"
	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"
	"go-form/dao"
	mod "go-form/models"
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
	forms, err := mod.QueryPageForm(pageNum, pageSize, name)

	//响应数据
	if err == nil {
		c.JSON(http.StatusOK, gin.H{
			"code":        0,
			"data":        forms,
			"pageSize":    pageSize,
			"pageNumber":  pageNum,
			"totalNumber": len(forms),
		})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 1,
			"msg":  "QueryPageForm获取数据时出错",
		})
	}
}

//获得某表单
func GetForm(c *gin.Context) {
	id := c.Param("id")
	var form mod.Form
	if query, arg, err := squirrel.Select("*").From("test").Where(squirrel.Eq{"id": id}).ToSql(); err == nil {
		if err := dao.Db.QueryRowx(query, arg...).StructScan(&form); err == nil {
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"data": form,
			})
			return
		} else {
			panic(err)
		}

	}

	c.JSON(http.StatusInternalServerError, gin.H{
		"code": 1,
		"msg":  "QueryFormById出错",
	})

}

type UpdateFormBody struct {
	Formdata interface{}              `json:"formdata" `
	Rule     []map[string]interface{} `json:"rule"`
	Name     string                   `json:"name" `
}

//编辑页面的保存表单，更新，注：？
func UpdateForm(c *gin.Context) {
	var updateFormBody UpdateFormBody
	query := squirrel.Update("test")
	id := c.Param("id")
	if err := c.ShouldBindJSON(&updateFormBody); err == nil {
		if updateFormBody.Formdata != nil {
			if jsonString, err := jsoniter.MarshalToString(updateFormBody.Formdata); err == nil {
				query = query.Set("formdata", jsonString)
			}
		}
		if updateFormBody.Rule != nil {
			if jsonString, err := jsoniter.MarshalToString(updateFormBody.Rule); err == nil {
				query = query.Set("rule", jsonString)
			}
		}
		if updateFormBody.Name != "" {
			query = query.Set("name", updateFormBody.Name)
		}
		if queryString, arg, err := query.Where(squirrel.Eq{"id": id}).ToSql(); err == nil {
			if _, err := dao.Db.Exec(queryString, arg...); err == nil {
				c.JSON(http.StatusOK, gin.H{
					"code": 0,
				})
			} else {
				c.Status(http.StatusBadRequest)
			}
		} else {
			c.Status(http.StatusBadRequest)
		}
	} else {

	}

}

//删除表单
func DeleteForm(c *gin.Context) {
	id := c.Param("id")

	if query, arg, err := squirrel.Delete("test").Where(squirrel.Eq{"id": id}).ToSql(); err == nil {
		if _, err := dao.Db.Exec(query, arg...); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code": 1,
				"msg":  "DeleteForm err",
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "删除成功",
			})
		}
	} else {
		c.Status(http.StatusInternalServerError)
	}

}

type CopyFormBody struct {
	FormId string
}

//创建表单，添加一个默认表单，再返回该表单
func CopyForm(c *gin.Context) {
	var copyFormBody CopyFormBody
	if err := c.ShouldBindJSON(&copyFormBody); err == nil { //更优的获取json参数，ShouldBlindJSON 只绑定结构体tag有binding:required的参数
		if record, err := mod.QueryFormById(copyFormBody.FormId); err != nil {
			if newFormId, err := mod.AddForm(record); err != nil {
				record.Id = int(newFormId)
				c.JSON(200, gin.H{
					"code": 0,
					"data": record,
				})
			} else {
				c.JSON(200, gin.H{
					"code":    1,
					"message": err,
				})
			}
		} else {
			c.JSON(200, gin.H{
				"code":    1,
				"message": "没有该表单项",
			})
		}
	} else {
		c.JSON(500, gin.H{
			"code": 1,
			"msg":  "AddDefaultForm err",
		})
	}

}

//创建表单，添加一个默认表单，再返回该表单
func CreateForm(c *gin.Context) {
	var m map[string]string
	if err := c.ShouldBindJSON(&m); err == nil { //更优的获取json参数，ShouldBlindJSON 只绑定结构体tag有binding:required的参数
		if id, err := mod.AddDefaultForm(m["name"]); err != nil {
			c.JSON(500, gin.H{
				"code": 1,
				"msg":  "AddDefaultForm err",
			})
		} else {
			c.JSON(200, gin.H{
				"code": 0,
				"id":   id,
			})
		}
	} else {
		c.JSON(500, gin.H{
			"code": 1,
			"msg":  "AddDefaultForm err",
		})
	}

}
