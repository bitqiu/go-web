package routes

import (
	"github.com/gin-gonic/gin"
	controllers "go-web/app/http/controllers/api/v1"
)

func RegisterAPIRoutes(r *gin.Engine) {
	v1 := r.Group("/api/v1")
	{
		apiCtrl := new(controllers.ApiController)

		v1.GET("/index", apiCtrl.Index)

		//获取标签列表
		v1.GET("/tags", apiCtrl.GetTags)
		//增加标签
		v1.POST("/tags", apiCtrl.AddTag)
		//编辑标签
		v1.PUT("/tags", apiCtrl.EditTag)
		//删除标签
		v1.DELETE("/tags", apiCtrl.DeleteTag)
	}
}
