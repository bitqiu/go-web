package middlewares

import "github.com/gin-gonic/gin"

func AuthJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		//1.获取 header 头 token
		//2.解析 token 值
		//3.判断解析是否成功
		//4.解析失败 返回401 错误
		//5.解析成功 写入 gin.context 上下文（数据共享）

		c.Next()
	}
}
