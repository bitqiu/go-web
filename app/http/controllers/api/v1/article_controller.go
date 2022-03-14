package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
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

// ArticleId 文章编号
type ArticleId struct {
	Id string `binding:"required" form:"id"`
}

// ArticleInfo 文章信息
type ArticleInfo struct {
	Title    string `binding:"required" form:"title"`
	SubTitle string `binding:"required" form:"sub_title"`
	Content  string `binding:"required" form:"content"`
	TagId    string `binding:"required" form:"tag_id"`
	//userId   string
}

// EditArticle 编辑文章
func (ctl *ApiController) EditArticle(c *gin.Context) {
	var articleInfo ArticleInfo
	var articleId ArticleId

	//绑定参数
	err := c.ShouldBindWith(&articleId, binding.Form)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	err = c.ShouldBindWith(&articleInfo, binding.Form)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	userId0, _ := c.Get("UserId")
	userId := cast.ToString(userId0)

	fmt.Println(articleInfo)

	ArticleModel := &article.Article{
		Title:    articleInfo.Title,
		SubTitle: articleInfo.SubTitle,
		Content:  articleInfo.Content,
		TagID:    articleInfo.TagId,
	}

	where := map[string]string{"user_id": userId, "id": articleId.Id}
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
