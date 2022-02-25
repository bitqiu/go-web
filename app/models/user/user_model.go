package user

import "go-web/app/models"

type User struct {
	models.BaseModel

	Name     string `json:"name" gorm:"not null;index"`
	Password string `json:"password" gorm:"not null"`
	Mobile   string `json:"mobile" gorm:"default:null"`

	models.CommonTimestampsField
	models.CommonSoftDeletesField
}
