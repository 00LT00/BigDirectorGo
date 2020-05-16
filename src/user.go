package main

import (
	"github.com/gin-gonic/gin"
)

func (s *Service) Registered(c *gin.Context) (int, interface{}) {
	//json中绑定的参数在table/user中规定
	json := new(User)

	err := c.ShouldBindJSON(json)
	if err != nil {
		s.makeErrJSON(403, 40301, err.Error())
	}
	tx := s.DB.Begin()
	if err := tx.Create(json).Error; err != nil {
		tx.Rollback()
		return s.makeErrJSON(500, 50000, err.Error())
	}
	tx.Commit()
	return s.makeSuccessJSON(json)
}

func (s *Service) GetUser(c *gin.Context) (int, interface{}) {
	user := new(User)
	userid := c.Param("userid")
	if s.DB.Where(&User{UserId: userid}).Find(user).RowsAffected != 1 {
		return s.makeErrJSON(404, 40400, "none user")
	}
	return s.makeSuccessJSON(user)
}

func (s *Service) UpdateUser(c *gin.Context) (int, interface{}) {
	userid := c.Param("userid")
	json := new(User)
	if err := c.ShouldBindJSON(json); err != nil {
		s.makeErrJSON(403, 40302, err.Error())
	}
	if userid != json.UserId {
		return s.makeErrJSON(403, 40303, "openid error")
	}
	if s.DB.Find(&User{}, User{UserId: userid}).RowsAffected != 1 {
		return s.makeErrJSON(403, 40401, "none user")
	}
	tx := s.DB.Begin()
	if err := tx.Model(&User{}).Updates(json).Error; err != nil {
		tx.Callback()
		return s.makeErrJSON(500, 50001, err.Error())
	}
	return s.makeSuccessJSON(json)
}
