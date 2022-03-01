package user

import (
	"go-web/app/models"
)

type User struct {
	models.BaseModel

	Name     string `json:"name" gorm:"type:varchar(255);not null;index;comment:用户名"`
	Password string `json:"password" gorm:"type:varchar(255);not null;"`

	models.CommonTimestampsField
	models.CommonSoftDeletesField
}

func (userModel *User) Create()  {

}

func (userModel *User) Update()  {

}

func (userModel *User) Delete()  {

}