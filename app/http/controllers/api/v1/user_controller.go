package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-web/app/models/user"
	"go-web/pkg/helpers"
	"net/http"
)

// Register 用户注册
func (ctl *ApiController) Register(c *gin.Context) {
	loginName := c.PostForm("login_name")
	name := c.PostForm("name")
	password := c.PostForm("password")
	//todo 参数验证

	//md5
	password = helpers.Md5(password)

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
	loginName := c.PostForm("login_name")
	password := c.PostForm("password")
	//todo 参数验证

	password = helpers.Md5(password)

	User := &user.User{LoginName: loginName, Password: password}
	User.GetByLoginNameAndPassword()

	//账号或密码错误
	if User.ID == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusBadRequest,
			"message": "账号或密码错误！",
		})
		c.Abort()
		return
	}

	//颁发token
	token, err := helpers.SignedJWT(User.ID)
	fmt.Println(err)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusBadRequest,
			"message": "系统错误！",
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": map[string]string{
			"token": token,
			"name":  User.Name,
		},
		"message": "success！",
	})
}
