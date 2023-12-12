package middleware

import (
	"algorithmplatform/app/common"
	"algorithmplatform/global"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenStr := c.Request.Header.Get("Authorization")
		if tokenStr == "" {
			c.AbortWithStatus(401)
			return
		}

		tokenStr = tokenStr[len("Bearer "):]

		var claims common.CustomJwt
		_, err := jwt.ParseWithClaims(tokenStr, &claims, func(t *jwt.Token) (interface{}, error) {
			return []byte(global.App.Config.Jwt.Secret), nil
		})
		if err != nil {
			c.AbortWithError(401, err)
			return
		}

		if issuer, err := claims.GetIssuer(); issuer != global.App.Config.Jwt.Issuer || err != nil {
			c.AbortWithError(401, err)
			return
		}

		if time.Now().Unix()-claims.ExpiresAt.Time.Unix() > int64(global.App.Config.Jwt.ClockSkewMinutes*60) {
			c.AbortWithError(401, err)
			return
		}
		userId, _ := strconv.ParseInt(claims.UserId, 10, 64)
		c.Set("userId", userId)
		c.Set("name", claims.Name)
	}
}
