package middleware

import (
	"log"
	"net/http"
	"strings"

	"github.com/260444/ginEssential/common"
	"github.com/260444/ginEssential/model"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取authorization header
		tokenstring := c.GetHeader("Authorization")
		// 验证token格式
		if tokenstring == "" || !strings.HasPrefix(tokenstring, "Bearer") {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "权限不足",
			})
			// 抛弃此次请求
			c.Abort()
			return
		}
		tokenstring = tokenstring[7:]
		token, claims, err := common.ParseToken(tokenstring)
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "权限不足",
			})
			log.Println("err", err)
			c.Abort()
			return
		}
		userId := claims.UserId
		log.Println("userId:", userId)
		DB := common.GetDB()
		var user model.User
		DB.First(&user, userId)
		if user.ID == 0 {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "权限不足",
			})
			c.Abort()
			return
		}
		c.Set("user", user)
		c.Next()
	}
}
