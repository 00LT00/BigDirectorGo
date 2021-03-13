package database

import (
	"gorm.io/gorm"
)

type Performance struct {
	PerformanceID string `json:"performanceId" gorm:"type:varchar(40);primaryKey"`
	Name          string `json:"name"`
	Place         string `json:"place"`
	Sponsor       string `json:"sponsor"`
	Time          string `json:"time"`
	Introduce     string `json:"introduce"`
	PosterImage   string `json:"posterImage"`
	ListImage     string `json:"listImage"`

	Processes []Process `json:"processes,omitempty" gorm:"foreignKey:PerformanceID;references:PerformanceID;constraint:OnUpdate:CASCADE"` // have many
	Groups    []Group   `json:"groups,omitempty" gorm:"foreignKey:PerformanceID;references:PerformanceID;constraint:OnUpdate:CASCADE"`    // have many
	Users     []*User   `json:"users,omitempty" gorm:"many2many:performance_users"`

	CreatedAt int
	UpdatedAt int
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type Process struct {
	ProcessID     string `json:"processId" gorm:"type:varchar(40);primaryKey"`
	PerformanceID string `json:"performanceId" gorm:"type:varchar(40);unique"`
	//UserID        string `gorm:"type:varchar(40);unique"`
	Name   string `json:"name"`
	Props  string `json:"props"`
	Mic    string `json:"mic"`
	Remark string `json:"remark"`

	//User          User        `gorm:"references:UserID;constraint:OnUpdate:CASCADE"`
	Performance Performance `json:"performance,omitempty" gorm:"references:PerformanceID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

	CreatedAt int
	UpdatedAt int
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type User struct {
	UserID string `json:"userId" gorm:"type:varchar(40);unique;primaryKey;not null"`
	Name   string `json:"name" gorm:"type:varchar(40)"`
	Phone  string `json:"phone"`
	Avatar string `json:"avatar"`

	Performances []*Performance `json:"performances,omitempty" gorm:"many2many:performance_users"` // many2many
	Groups       []*Group       `json:"groups,omitempty" gorm:"many2many:group_users"`             // many2many

	CreatedAt int
	UpdatedAt int
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type Group struct {
	GroupID       string `json:"groupId" gorm:"type:varchar(40);primaryKey"`
	PerformanceID string `json:"performanceId" gorm:"type:varchar(40);unique"`
	Name          string `json:"name"`

	Users []*User `json:"users,omitempty" gorm:"many2many:group_users"` // many2many

	CreatedAt int
	UpdatedAt int
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
