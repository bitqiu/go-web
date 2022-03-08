package middlewares

import (
	"github.com/gin-gonic/gin"
	"go-web/pkg/helpers"
	"net/http"
)

func AuthJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		//1.获取 header 头 token
		//2.解析 token 值
		//3.判断解析是否成功
		//4.解析失败 返回401 错误
		//5.解析成功 写入 gin.context 上下文（数据共享）
		tokenString := c.GetHeader("token")
		if helpers.Empty(tokenString) {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code":    401,
				"message": "权限不足！",
			})
			c.Abort()
			return
		}

		token, claims, err := helpers.ParseToken(tokenString)
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code":    401,
				"message": "权限不足！",
			})
			c.Abort()
			return
		}

		c.Set("UserId", claims.UserId)

		//c.Next()
	}
}
