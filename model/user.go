package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `json:"name" gorm:"type:varchar(25);not null;comment:'用户名'"`
	Telphone string `json:"telphone" gorm:"type:varchar(11);not null;comment:'手机号'"`
	Password string `json:"password" gorm:"type:varchar(255);not null;comment:'密码'"`
}
