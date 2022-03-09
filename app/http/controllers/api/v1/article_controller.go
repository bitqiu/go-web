package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"go-web/app/models/article"
	"net/http"
)

// GetArticles 获取文章
func (ctl *ApiController) GetArticles(c *gin.Context) {
	pageNum := cast.ToInt(c.DefaultQuery("page_num", "1"))
	pageSize := cast.ToInt(c.DefaultQuery("page_size", "10"))
	userId, _ := c.Get("UserId")

	where := map[string]string{"user_id": cast.ToString(userId)}

	ArticleModel := &article.Article{}

	data := ArticleModel.Get(pageNum, pageSize, where)
	total := ArticleModel.GetTotal(where)

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "success",
		"data":    data,
		"total":   total,
	})
}

// CreateArticle 创建文章
func (ctl *ApiController) CreateArticle(c *gin.Context) {
	title := c.PostForm("title")
	subTitle := c.PostForm("sub_title")
	content := c.PostForm("content")
	tagId := c.PostForm("tag_id")
	userId0, _ := c.Get("UserId")
	userId := cast.ToString(userId0)

	ArticleModel := &article.Article{
		Title:    title,
		SubTitle: subTitle,
		Content:  content,
		UserID:   userId,
		TagID:    tagId,
	}

	result := ArticleModel.Create()

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

// EditArticle 编辑文章
func (ctl *ApiController) EditArticle(c *gin.Context) {
	id := c.PostForm("id")
	title := c.PostForm("title")
	subTitle := c.PostForm("sub_title")
	content := c.PostForm("content")
	tagId := c.PostForm("tag_id")
	userId0, _ := c.Get("UserId")
	userId := cast.ToString(userId0)

	ArticleModel := &article.Article{
		Title:    title,
		SubTitle: subTitle,
		Content:  content,
		TagID:    tagId,
	}

	where := map[string]string{"user_id": userId, "id": id}
	result := ArticleModel.Edit(where)

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

// DeleteArticle 删除文章
func (ctl *ApiController) DeleteArticle(c *gin.Context) {
	id := c.PostForm("id")
	userId0, _ := c.Get("UserId")
	userId := cast.ToString(userId0)

	ArticleModel := &article.Article{}

	where := map[string]string{"user_id": userId, "id": id}

	result := ArticleModel.Delete(where)

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
