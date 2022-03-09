package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"go-web/app/models/tag"
	"net/http"
)

// GetTags 获取标签
func (ctl *ApiController) GetTags(c *gin.Context) {
	pageNum := cast.ToInt(c.DefaultQuery("page_num", "1"))
	pageSize := cast.ToInt(c.DefaultQuery("page_size", "10"))
	userId, _ := c.Get("UserId")

	where := map[string]string{"user_id": cast.ToString(userId)}

	data := tag.Get(pageNum, pageSize, where)
	total := tag.GetTotal(where)

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "success",
		"data":    data,
		"total":   total,
	})
}

// CreateTag 增加标签
func (ctl *ApiController) CreateTag(c *gin.Context) {
	name := c.PostForm("name")
	userId0, _ := c.Get("UserId")
	userId := cast.ToString(userId0)

	if tag.ExitByName(name, userId) {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "标签名已存在！",
		})
		return
	}

	result := tag.Create(name, userId)

	code := http.StatusOK
	message := "success"
	if !result {
		code = http.StatusBadRequest
		message = "failed"
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    code,
		"message": message,
	})
}

// EditTag 编辑标签
func (ctl *ApiController) EditTag(c *gin.Context) {
	name := c.PostForm("name")
	id := c.PostForm("id")
	userId0, _ := c.Get("UserId")
	userId := cast.ToString(userId0)

	result := tag.Edit(name, id, userId)

	code := http.StatusOK
	message := "success"
	if !result {
		code = http.StatusBadRequest
		message = "failed"
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    code,
		"message": message,
	})
}

// DeleteTag 删除标签
func (ctl *ApiController) DeleteTag(c *gin.Context) {
	id := c.PostForm("id")
	userId0, _ := c.Get("UserId")
	userId := cast.ToString(userId0)

	result := tag.Delete(id, userId)

	code := http.StatusOK
	message := "success"
	if !result {
		code = http.StatusBadRequest
		message = "failed"
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    code,
		"message": message,
	})
}
