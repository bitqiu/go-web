package article

import (
	"go-web/app/models"
	"go-web/app/models/tag"
	"go-web/app/models/user"
	"go-web/global"
)

type Article struct {
	models.BaseModel

	Title    string `json:"title" gorm:"not null;index:idx_title_sub"`
	SubTitle string `json:"sub_title" gorm:"index:idx_title_sub"`
	Content  string `json:"content" gorm:"text"`

	// 用户关联
	UserID string    `json:"user_id" gorm:"index"`
	User   user.User `json:"user"`

	// 标签关联
	TagID string  `json:"tag_id" gorm:"index"`
	Tag   tag.Tag `json:"tag"`

	models.CommonTimestampsField
	models.CommonSoftDeletesField
}

// Create 创建文章
func (ArticleModel *Article) Create() bool {
	result := global.DB.Create(ArticleModel)
	return result.RowsAffected > 0
}

// Get 获取文章列表
func (ArticleModel *Article) Get(pageNum, pageSize int, where interface{}) (article []Article) {
	pageNum = (pageNum - 1) * pageSize
	global.DB.Where(where).Offset(pageNum).Order("id desc").Limit(pageSize).Find(&article)
	return
}

// GetTotal 获取文章数量
func (ArticleModel *Article) GetTotal(where interface{}) (count int64) {
	global.DB.Model(ArticleModel).Where(where).Count(&count)
	return
}

// Edit 编辑文章
func (ArticleModel *Article) Edit(where interface{}) bool {
	result := global.DB.Where(where).Updates(ArticleModel)
	return result.RowsAffected > 0
}

// Delete 删除文章
func (ArticleModel *Article) Delete(where interface{}) bool {
	result := global.DB.Where(where).Delete(ArticleModel)
	return result.RowsAffected > 0
}
