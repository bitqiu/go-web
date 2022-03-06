package tag

import (
	"go-web/app/models"
	"go-web/app/models/user"
	"go-web/global"
	"gorm.io/gorm/clause"
)

type Tag struct {
	models.BaseModel

	Name string `json:"name" gorm:"unique"`

	// 用户关联
	UserID string    `json:"user_id" gorm:"index"`
	User   user.User `json:"user" gorm:"migration"`

	models.CommonTimestampsField
	models.CommonSoftDeletesField
}

func GetBy(field, value string) (tag Tag) {
	global.DB.Where(field, value).First(&tag)
	return
}

func All() (tags []Tag) {
	global.DB.Find(&tags)
	return
}

func FindBy(field, value string) (tags []Tag) {
	global.DB.Where(field, value).Find(&tags)
	return
}

// Paginate 接收分页参数 返回 页码和数据
func Paginate(page, pageSize int) (tags []Tag, currentPage int, total int64) {
	// 当前页码
	if page <= 0 {
		// 默认为 1
		page = 1
	}

	offset := (page - 1) * pageSize

	global.DB.Debug().Preload(clause.Associations).
		Limit(pageSize). // 每页显示
		Offset(offset).
		Find(&tags)

	currentPage = page

	global.DB.Debug().Count(&total)

	return
}

// GetTags 获取标签列表
func GetTags(pageNum, pageSize int, where interface{}) (tags []Tag) {
	pageNum = (pageNum - 1) * pageSize
	global.DB.Where(where).Offset(pageNum).Limit(pageSize).Find(&tags)
	return
}

// GetTagsTotal 获取标签数量
func GetTagsTotal(where interface{}) (count int64) {
	global.DB.Model(&Tag{}).Where(where).Count(&count)
	return
}

func ExitTagByName(name string) bool {
	var tag Tag
	global.DB.Select("id").Where("name = ?", name).First(&tag)
	return tag.ID > 0
}

// AddTag 增加标签
func AddTag(name, userID string) bool {
	result := global.DB.Create(&Tag{Name: name, UserID: userID})
	return result.RowsAffected > 0
}

// EditTAg 修改标签
func EditTAg(name, id string) bool {
	result := global.DB.Model(&Tag{}).Where("id = ?", id).Update("name", name)
	return result.RowsAffected > 0
}

// DeleteTag 删除标签
func DeleteTag(id string) bool {
	result := global.DB.Delete(&Tag{}, id)
	return result.RowsAffected > 0
}
