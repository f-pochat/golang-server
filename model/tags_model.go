package model

import "github.com/f-pochat/golang-base/model"

type Tags struct {
	model.BaseModel
	Name string `gorm:"type:varchar(255)"`
}
