package main

import "time"

type User struct {
	UserId   string `form:"openid" binding:"required" json:"openid" gorm:"primary_key;type:varchar(30);not null;unique"`
	UserName string `form:"username" binding:"required" json:"username" gorm:"not null"`
	PhoneNum string `form:"phonenum" binding:"required" json:"phonenum" gorm:"not null"`
	Avatar   string `form:"avatar" binding:"required" json:"avatar" gorm:"not null"`
	QQnum    string `form:"qqnum" binding:"-" json:"qqnum" gorm:"column:qq_num"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
