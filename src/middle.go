package main

import "github.com/gin-gonic/gin"

// 校验码
func (s *Service) Check() gin.HandlerFunc {
	return func(c *gin.Context) {
		sign:=c.GetHeader("sign")
		if sign != s.Conf.Server.Key {
			c.JSON(s.makeErrJSON(403, 40300, "forbidden"))
			c.Abort()
			return
		}

		c.Next()
	}
}
