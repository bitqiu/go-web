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

	where := map[string]string{}
	data := tag.GetTags(pageNum, pageSize, where)
	total := tag.GetTagsTotal(where)

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "success",
		"data":    data,
		"total":   total,
	})
}

// AddTag 增加标签
func (ctl *ApiController) AddTag(c *gin.Context) {
	name := c.PostForm("name")
	userId := c.DefaultPostForm("user_id", "1")
	result := tag.AddTag(name, userId)

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

// EditTag 编辑标签
func (ctl *ApiController) EditTag(c *gin.Context) {
	name := c.PostForm("name")
	id := c.PostForm("id")
	result := tag.EditTAg(name, id)

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
	result := tag.DeleteTag(id)

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
