package main

import (
	"github.com/jinzhu/gorm"
	"time"
)

type User struct {
	UserId    string `form:"openid" binding:"required" json:"openid" gorm:"primary_key;type:varchar(30);not null;unique"`
	UserName  string `form:"username" binding:"required" json:"username" gorm:"not null"`
	PhoneNum  string `form:"phonenum" binding:"required" json:"phonenum" gorm:"not null"`
	Avatar    string `form:"avatar" binding:"required" json:"avatar" gorm:"not null"`
	QQnum     string `form:"qqnum" binding:"-" json:"qqnum" gorm:"column:qq_num"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

type Project struct {
	gorm.Model
	DirectorUserID string `form:"UserID" json:"userid" binding:"required" gorm:"not null"`
	Name           string `form:"name" json:"name" binding:"required" gorm:"not null"`
	ProjectID      string `binding:"-" gorm:"not null;unique;unique_index;type:varchar(40)"`
}

type Project_User struct {
	gorm.Model
	UserID    string `gorm:"not null"`
	ProjectID string `gorm:"not null"`
	Role      int    `gorm:"not null"`
}

var RoleTable = map[string]int{
	"director":  1,
	"manager":   2,
	"member":    3,
	"music":     4,
	"light":     5,
	"backstage": 6,
}

//更改表名
func (Project_User) TableName() string {
	return "project_user"
}
