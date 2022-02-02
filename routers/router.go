package routers

import (
	"github.com/gin-gonic/gin"
	ctrl "go-form/controller"
)

func SetupRouter() (*gin.Engine){
	r := gin.Default()

	//表单工具
	{
		//查询所有
		r.GET("/getform", ctrl.GetPageForm)
		//查看一个
		r.GET("/getform/:id", ctrl.GetForm)
		//更新
		r.PUT("/getform/:id", ctrl.UpdateForm)
		//创建
		r.POST("/getform", ctrl.CreateForm)
		//删除
		r.DELETE("/getform/:id", ctrl.DeleteForm)
	}

	return r
}
