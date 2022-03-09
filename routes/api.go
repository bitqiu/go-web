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

		//获取文章列表
		v1.GET("/articles", apiCtrl.GetArticles)
		//增加文章
		v1.POST("/articles", apiCtrl.CreateArticle)
		//编辑文章
		v1.PUT("/articles", apiCtrl.EditArticle)
		//删除文章
		v1.DELETE("/articles", apiCtrl.DeleteArticle)

	}

	v2 := r.Group("/api/v1")
	{
		//用户注册
		v2.POST("/register", apiCtrl.Register)
		//用户登录
		v2.POST("/login", apiCtrl.Login)
	}
}
