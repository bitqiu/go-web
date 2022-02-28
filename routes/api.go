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
	}
}
