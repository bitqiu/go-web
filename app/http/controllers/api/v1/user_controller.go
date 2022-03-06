package v1

import (
	"crypto/md5"
	"fmt"
	"github.com/gin-gonic/gin"
	"go-web/app/models/user"
	"io"
	"net/http"
)

// Register 用户注册
func (ctl *ApiController) Register(c *gin.Context) {
	loginName := c.PostForm("login_name")
	name := c.PostForm("name")
	password := c.PostForm("password")

	//md5
	h := md5.New()
	io.WriteString(h, password)
	password = fmt.Sprintf("%x", h.Sum(nil))

	userModel := &user.User{LoginName: loginName, Name: name, Password: password}

	//验证账号是否存在
	if userModel.ExitByLoginName() {
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusBadRequest,
			"message": "账号已被注册！",
		})
		c.Abort()
		return
	}

	//创建用户
	result := userModel.Create()

	code := http.StatusOK
	message := "success"
	if !result {
		code = http.StatusBadRequest
		message = "failed"
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": message,
	})
}

// Login 用户登录
func (ctl *ApiController) Login(c *gin.Context) {

}
