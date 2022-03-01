package v1

import (
	"github.com/gin-gonic/gin"
	"go-web/app/models/tag"
	"net/http"
)

// GetTags 获取标签
func (ctl *ApiController) GetTags(c *gin.Context) {
	where := map[string]int{"id": 1}
	result := tag.GetTags(1, 10, where)
	c.JSON(http.StatusOK, result)
}

// AddTag 增加标签
func (ctl *ApiController) AddTag(c *gin.Context) {
	name := c.PostForm("name")
	userId := c.DefaultPostForm("user_id", "1")
	result := tag.AddTag(name, userId)
	c.JSON(http.StatusOK, result)
}

// EditTag 编辑标签
func (ctl *ApiController) EditTag(c *gin.Context) {

}

// DeleteTag 删除标签
func (ctl *ApiController) DeleteTag(c *gin.Context) {

}
