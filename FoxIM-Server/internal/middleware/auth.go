package middleware

import (
	"fmt"
	"gin_template/global"
	"gin_template/internal/libs"
	"gin_template/pkg"
	"strings"

	"github.com/gin-gonic/gin"
)

// AuthToken token校验
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 优先从请求头获取 token
		token := c.Request.Header.Get("Authorization")

		parts := strings.SplitN(token, " ", 2)

		token = parts[1]

		ip := c.ClientIP()
		// 检查 token 是否存在
		if token == "" {
			pkg.Error("token 不存在", nil)
			libs.BadRequestResponse(c, fmt.Sprintf("非法请求，请停止任何非法操作，IP %s 已记录！", ip))
			return
		}

		// 尝试解密 token
		currentUser, err := libs.ParseToken(token)
		if err != nil {
			pkg.Error("token 解密失败", err)
			libs.UnauthorizedResponse(c, fmt.Sprintf("非法请求，解密 token 失败，IP %s 已记录！", ip))
			return
		}

		// 检查 Redis 中是否存在 token
		redisKey := fmt.Sprintf("%s%d", global.RedisSessionKey, currentUser.ID)
		_, err = libs.RedisGet(c, redisKey)
		if err != nil {
			pkg.Error("Redis 中不存在 token", err)
			libs.UnauthorizedResponse(c, fmt.Sprintf("非法请求，用户未认证，IP %s 已记录！", ip))
			return
		}

		// 将用户名设置到上下文中
		c.Set("id", currentUser.ID)
		c.Next()
	}
}
