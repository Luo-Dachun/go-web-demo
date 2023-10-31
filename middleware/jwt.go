package middleware

import (
	"errors"
	"time"

	"github.com/gin-gonic/gin"
	"web-demo/model/common/response"
	"web-demo/utils"
)

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		if token == "" {
			response.FailWithDetailed(gin.H{"reload": true}, "未登录或非法访问", c)
			c.Abort()
			return
		}

		j := utils.NewJWT()
		claims, err := j.ParseToken(token)
		if err != nil {
			if errors.Is(err, utils.TokenExpired) {
				response.FailWithDetailed(gin.H{"reload": true}, "授权已过期", c)
				c.Abort()
				return
			}
			response.FailWithDetailed(gin.H{"reload": true}, err.Error(), c)
			c.Abort()
			return
		}
		if claims.ExpiresAt.Unix()-time.Now().Unix() < claims.BufferTime {
			response.FailWithDetailed(gin.H{"reload": true}, "token已过期，请重新登录", c)
			c.Abort()
			return

			/*
				改成 token的过期时间与当前时间的差值小于预设的缓冲时间，则更新token的过期时间，并创建一个新的token。
				同时，将新的token和过期时间添加到响应头部中。
			*/
		}
		c.Next()
	}
}
