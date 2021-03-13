package user

import (
	error2 "BigDirector/error"
	"github.com/gin-gonic/gin"
)

func OpenID(c *gin.Context) interface{} {
	code := c.Query("code")
	if code == "" {
		panic(error2.NewHttpError(404, "40401", "code null"))
	}
	return getOpenID(code)
}
