package middleware

import (
	. "coffee_backend/models"
	"net/http"

	"errors"
	"strings"

	"github.com/gin-gonic/gin"
)

// AuthJWT 验证 JWT 的中间件
func AuthJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader("Authorization")
		headerList := strings.Split(header, " ")
		if len(headerList) != 2 {
			err := errors.New("无法解析 Authorization 字段")
			c.JSON(http.StatusOK, gin.H{
				"result": err,
			})
			c.Abort()
			return
		}
		t := headerList[0]
		content := headerList[1]
		if t != "Bearer" {
			err := errors.New("认证类型错误, 当前只支持 Bearer")
			c.JSON(http.StatusOK, gin.H{
				"result": err,
			})
			c.Abort()
			return
		}
		// content := c.GetHeader("Authorization")
		if _, err := Verify([]byte(content)); err != nil {
			c.JSON(http.StatusOK, gin.H{
				"result": err.Error,
			})
			c.Abort()
			return
		}

		c.Next()
	}
}
