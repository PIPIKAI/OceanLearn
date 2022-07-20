package middleware

import (
	"net/http"
	"server/common"
	"server/dto"
	"server/model"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	// 定义auth的中间件 用来检查jwt权限
	return func(ctx *gin.Context) {
		tokenString := ctx.GetHeader("Authorization")
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer") {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足"})
			// 直接放弃next
			ctx.Abort()
			return
		}
		token, clamis, err := common.ParseToken(tokenString[7:])
		if err != nil || !token.Valid {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足"})
			// 直接放弃next
			ctx.Abort()
			return
		}
		// 验证成功后 根据clamis 中的UserId 查询数据库
		userId := clamis.UserId

		DB := common.GetDB()
		var user model.User
		DB.First(&user, userId)
		// 是否查询的到用户
		if user.ID == 0 {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足"})
			// 直接放弃next
			ctx.Abort()
			return
		}
		// 用户存在 将user 在ctx中设置
		ctx.Set("user", dto.ToUserDto(user))
		ctx.Next()
	}
}
