package main

import (
	"github.com/jinzhu/gorm"
	"time"
)

//用户表
type User struct {
	UserId    string `form:"openid" binding:"required" json:"openid" gorm:"primary_key;type:varchar(30);not null;unique"`
	UserName  string `form:"username" binding:"required" json:"username" gorm:"not null"`
	PhoneNum  string `form:"phonenum" binding:"-" json:"phonenum" `
	Avatar    string `form:"avatar" binding:"required" json:"avatar" gorm:"not null"`
	QQnum     string `form:"qqnum" binding:"-" json:"qqnum" gorm:"column:qq_num"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

//项目表
type Project struct {
	gorm.Model
	DirectorUserID string `form:"userid" json:"userid" binding:"required" gorm:"not null"`
	Name           string `form:"name" json:"name" binding:"required" gorm:"not null"`
	ProjectID      string `binding:"-" gorm:"not null;unique;unique_index;type:varchar(40)"`
}

//权限表
type Project_User struct {
	gorm.Model
	UserID    string `gorm:"not null" json:"userid" form:"userid" binding:"required"`
	ProjectID string `gorm:"not null" json:"projectid" form:"projectid" binding:"required"`
	Role      int    `gorm:"not null" json:"-" form:"-"`
}

//更改表名
func (Project_User) TableName() string {
	return "project_user"
}

//环节表
type Process struct {
	Order       int64  `gorm:"not null" json:"order" binding:"required"`
	ProcessID   string `gorm:"not null;type:varchar(40)" json:"process_id" binding:"-"` //流程id 自己生成
	ProcessName string `gorm:"not null" json:"process_name" binding:"required"`
	ProcessType int64  `gorm:"not null" json:"process_type" binding:"required"`
	MicHand     int64  `json:"mic_hand" binding:"-"` //可选
	MicEar      int64  `json:"mic_ear" binding:"-"`  //可选
	Remark      string `json:"remark" binding:"-"`   //可选
	ProjectID   string `gorm:"not null" json:"project_id" binding:"required"`
}

//环节类型
var ProcessTypeArr = [6]string{"节目", "互动", "颁奖", "致辞", "开场", "结束"}

//数据字典
var RoleTable = map[interface{}]interface{}{
	"director":  1,
	"manager":   2,
	"member":    3,
	"music":     4,
	"light":     5,
	"backstage": 6,
	"prop":      7,
	1:           "director",
	2:           "manager",
	3:           "member",
	4:           "music",
	5:           "light",
	6:           "backstage",
	7:           "prop",
}
