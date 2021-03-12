package utils

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func MakeErrJSON(httpStatusCode int, ErrorCode string, msg interface{}) (int, interface{}) {
	return httpStatusCode, &gin.H{"error": ErrorCode, "msg": fmt.Sprint(msg)}
}

func MakeSuccessJSON(data interface{}) (int, interface{}) {
	return 200, &gin.H{"error": 0, "msg": "success", "data": data}
}
