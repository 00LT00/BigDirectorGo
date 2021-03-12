package gin

import (
	"BigDirector/utils"
	"github.com/gin-gonic/gin"
)

func initRouter(r *gin.Engine) {
	r.Group("/")

}

func f(m func(c *gin.Context) interface{}) gin.HandlerFunc {
	return func(c *gin.Context) {
		data := m(c)
		c.JSON(utils.MakeSuccessJSON(data))
	}
}
