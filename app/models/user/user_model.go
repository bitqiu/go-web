package user

import (
	"go-web/app/models"
	"go-web/global"
)

type User struct {
	models.BaseModel

	LoginName string `json:"login_nam" gorm:"type:varchar(255);not null;index;comment:账号"`
	Name      string `json:"name" gorm:"type:varchar(255);not null;index;comment:用户名"`
	Password  string `json:"password" gorm:"type:varchar(255);not null;"`

	models.CommonTimestampsField
	models.CommonSoftDeletesField
}

func (userModel *User) Create() bool {
	result := global.DB.Create(&userModel)
	return result.RowsAffected > 0
}

func (userModel *User) Update() {

}

func (userModel *User) Delete() {

}

// ExitByLoginName 账号是否存在
func (userModel *User) ExitByLoginName() bool {
	global.DB.Select("id").Where("login_name = ?", userModel.LoginName).First(&userModel)
	return userModel.ID > 0
}
