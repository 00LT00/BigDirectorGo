package main

import "github.com/gin-gonic/gin"

func (s *Service) OpenID(c *gin.Context) (int, interface{}) {
	code := c.Param("code")
	openid, err := s.GetOpenID(code)
	if err != nil {
		return s.makeErrJSON(500, 50000, err.Error())
	}
	return s.makeSuccessJSON(openid)
}
