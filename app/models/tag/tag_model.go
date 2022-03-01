package tag

import (
	"go-web/app/models"
	"go-web/app/models/user"
)

type Tag struct {
	models.BaseModel

	Name string `json:"name"`

	// 用户关联
	UserID string    `json:"user_id" gorm:"index"`
	User   user.User `json:"user"`

	models.CommonTimestampsField
	models.CommonSoftDeletesField
}

// GetTags 获取标签列表
func GetTags(pageNum, pageSize int, maps interface{}) (tags []Tag) {
	models.Db.Where(maps).Offset(pageSize).Limit(pageNum).Find(&tags)
	return
}

// GetTagsTotal 获取标签数量
func GetTagsTotal(maps interface{}) (count int64) {
	models.Db.Model(&Tag{}).Where(maps).Count(&count)
	return
}

// AddTag 增加标签
func AddTag(name, userID string) bool {
	result := models.Db.Create(Tag{Name: name, UserID: userID})
	return result.RowsAffected > 0
}
