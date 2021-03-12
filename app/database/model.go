package database

type Performance struct {
	PerformanceID string `gorm:"type:varchar(40);primaryKey"`
	Name          string
	Place         string
	Sponsor       string
	Time          string
	Introduce     string
	PosterImage   string
	ListImage     string

	Processes []Process `gorm:"foreignKey:PerformanceID;references:PerformanceID;constraint:OnUpdate:CASCADE"` // have many
	Groups    []Group   `gorm:"foreignKey:PerformanceID;references:PerformanceID;constraint:OnUpdate:CASCADE"` // have many
	Users     []*User   `gorm:"many2many:performance_users"`
}

type Process struct {
	ProcessID     string `gorm:"type:varchar(40);primaryKey"`
	PerformanceID string `gorm:"type:varchar(40);unique"`
	//UserID        string `gorm:"type:varchar(40);unique"`
	Name   string
	Props  string
	Mic    string
	Remark string

	//User          User        `gorm:"references:UserID;constraint:OnUpdate:CASCADE"`
	Performance Performance `gorm:"references:PerformanceID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type User struct {
	UserID string `gorm:"type:varchar(40);unique;primaryKey;not null"`
	Name   string `gorm:"type:varchar(40)"`

	Performances []*Performance `gorm:"many2many:performance_users"` // many2many
	Groups       []*Group       `gorm:"many2many:group_users"`       // many2many

}

type Group struct {
	GroupID       string `gorm:"type:varchar(40);primaryKey"`
	PerformanceID string `gorm:"type:varchar(40);unique"`
	Name          string

	Users []*User `gorm:"many2many:group_users"` // many2many
}
