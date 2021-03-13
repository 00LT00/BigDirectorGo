package gin

import (
	error2 "BigDirector/error"
	logger "BigDirector/log"
	"BigDirector/utils"
	"github.com/gin-gonic/gin"
)

// 请求头校验
func check(c *gin.Context) {
	sign := c.GetHeader("sign")
	if sign != conf.Server.Sign {
		c.AbortWithStatusJSON(utils.MakeErrJSON(403, "40300", "forbidden"))
	}
}

// Http错误处理
func HttpRecover(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			if e, ok := r.(error2.HttpError); ok {
				c.AbortWithStatusJSON(utils.MakeErrJSON(e.HttpStatusCode, e.Err.ErrCode, e.Err.Msg))
			} else {
				c.AbortWithStatusJSON(utils.MakeErrJSON(500, "50000", "service error"))
			}
			logger.ErrLog.Println("[GINERROR]:%+V", r)
		}
	}()
	c.Next()
}
