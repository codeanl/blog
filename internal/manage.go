package internal

import (
	"blog/internal/hander"

	"github.com/gin-gonic/gin"
)

var (
	userAPI      hander.User      // 用户
	universalApi hander.Universal // 通用
)

// 使用外观设计模式一口气全部完成注册

func RegisterAllHandler(r *gin.Engine) {
	RegisterBaseHandler(r)
}

// 通用的handler
func RegisterBaseHandler(r *gin.Engine) {
	base := r.Group("/api")
	base.POST("/sendCode", universalApi.SendCode)
	base.POST("/register", userAPI.Register)
	base.POST("/login", userAPI.Login)
}
