package middleware

import (
	"algorithmplatform/app/common/response"

	"github.com/gin-gonic/gin"
)

func Recover() gin.HandlerFunc {
	return func(c *gin.Context) {

		defer func() {
			if err := recover(); err != nil {
				response.FailError(c, err.(error))
			}
		}()
		c.Next()
	}
}
