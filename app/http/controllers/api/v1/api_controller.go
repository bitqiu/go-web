package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type ApiController struct {
	// 继承 base
	BaseController
}

func (ctl *ApiController) Index(c *gin.Context)  {
	c.JSON(http.StatusOK, gin.H{
		"mesg": "api/v1/index",
	})
}
