package routes

import (
	"github.com/gin-gonic/gin"
	controllers "go-web/app/http/controllers/api/v1"
	"go-web/app/http/middlewares"
)

func RegisterAPIRoutes(r *gin.Engine) {
	apiCtrl := new(controllers.ApiController)

	v1 := r.Group("/api/v1")
	v1.Use(middlewares.AuthJWT())
	{

		v1.GET("/index", apiCtrl.Index)
		//获取标签列表
		v1.GET("/tags", apiCtrl.GetTags)
		//增加标签
		v1.POST("/tags", apiCtrl.CreateTag)
		//编辑标签
		v1.PUT("/tags", apiCtrl.EditTag)
		//删除标签
		v1.DELETE("/tags", apiCtrl.DeleteTag)

	}

	v2 := r.Group("/api/v1")
	{
		//用户注册
		v2.POST("/register", apiCtrl.Register)
		//用户登录
		v2.POST("/login", apiCtrl.Login)
	}
}
