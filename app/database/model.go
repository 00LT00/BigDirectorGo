package database

import (
	"gorm.io/gorm"
)

type Performance struct {
	PerformanceID string `json:"performanceID" gorm:"type:varchar(40);primaryKey"`
	Name          string `json:"name" binding:"required"`
	Place         string `json:"place"`
	Sponsor       string `json:"sponsor"`
	Duration      string `json:"duration"`
	StartTime     string `json:"startTime"`
	Introduce     string `json:"introduce"`
	PosterImage   string `json:"posterImage"`
	ListImage     string `json:"listImage"`

	Processes []*Process `json:"processes,omitempty" gorm:"foreignKey:PerformanceID;references:PerformanceID;constraint:OnUpdate:CASCADE" swaggerignore:"true"` // have many
	Groups    []*Group   `json:"groups,omitempty" gorm:"foreignKey:PerformanceID;references:PerformanceID;constraint:OnUpdate:CASCADE" swaggerignore:"true"`    // have many
	Users     []*User    `json:"users,omitempty" gorm:"many2many:performance_users" swaggerignore:"true"`                                                       // many2many

	CreatedAt int            `json:"-" swaggerignore:"true"`
	UpdatedAt int            `json:"-" swaggerignore:"true"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index" swaggerignore:"true"`
}

type Process struct {
	ProcessID     string `json:"processID" gorm:"type:varchar(40);primaryKey"`
	PerformanceID string `json:"performanceID" binding:"required" gorm:"type:varchar(40);unique"`
	//OpenID        string `gorm:"type:varchar(40);unique"`
	Name   string `json:"name"`
	Props  string `json:"props"`
	Mic    string `json:"mic"`
	Remark string `json:"remark"`

	Performance *Performance `json:"performance,omitempty" gorm:"references:PerformanceID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" swaggerignore:"true"` //belongs to

	CreatedAt int            `json:"-" swaggerignore:"true"`
	UpdatedAt int            `json:"-" swaggerignore:"true"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index" swaggerignore:"true"`
}

type User struct {
	OpenID string `json:"openID" binding:"required" gorm:"type:varchar(40);unique;primaryKey;not null"`
	Name   string `json:"name" gorm:"type:varchar(40)"`
	Phone  string `json:"phone"`
	Avatar string `json:"avatar"`

	Performances []*Performance `json:"performances,omitempty" gorm:"many2many:performance_users" swaggerignore:"true"` // many2many
	Groups       []*Group       `json:"groups,omitempty" gorm:"many2many:group_users" swaggerignore:"true"`             // many2many

	CreatedAt int            `json:"-" swaggerignore:"true"`
	UpdatedAt int            `json:"-" swaggerignore:"true"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index" swaggerignore:"true"`
}

type Group struct {
	GroupID       string `json:"groupID" gorm:"type:varchar(40);primaryKey"`
	PerformanceID string `json:"performanceID" binding:"required" gorm:"type:varchar(40);unique"`
	Name          string `json:"name"`
	LeaderID      string `json:"leaderID" gorm:"type:varchar(40);unique"`

	Leader *User    `json:"leader,omitempty" binding:"-" gorm:"foreignKey:LeaderID;references:OpenID;constraint:OnUpdate:CASCADE" swaggerignore:"true"` //has one
	Users  []*User  `json:"users,omitempty" gorm:"many2many:group_users" swaggerignore:"true"`                                                          // many2many
	Roles  []string `json:"roles,omitempty" gorm:"-"`

	CreatedAt int            `json:"-" swaggerignore:"true"`
	UpdatedAt int            `json:"-" swaggerignore:"true"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index" swaggerignore:"true"`
}
