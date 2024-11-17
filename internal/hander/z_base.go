package hander

import (
	"blog/internal/global"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"gorm.io/gorm"
)

// 定义响应结构体
type Response[T any] struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    T      `json:"data"`
}

// 从上下文中获gorm.DB
func GetDB(c *gin.Context) *gorm.DB {
	return c.MustGet(global.CTX_DB).(*gorm.DB)
}

// 从上下文中获redis.Client
func GetRDB(c *gin.Context) *redis.Client {
	return c.MustGet(global.CTX_RDB).(*redis.Client)
}

// 预料中的错误  = 业务错误 + 系统错误  在业务层处理，返回200 http状态码
// 意外的错误  = 触发panic， 在中间件中被捕获，返回500 http状态码
// data是错误数据（可以是error和string）， error是业务错误
func ReturnError(c *gin.Context, r global.Result, data any) {
	slog.Info("[FUNC-RETURN-ERROR]] :" + r.Msg)
	var val string = r.Msg

	if data != nil {
		switch v := data.(type) {
		case error:
			val = v.Error()

		case string:
			val = v
		}
		slog.Error(val)
	}

	c.AbortWithStatusJSON(
		http.StatusOK,
		Response[any]{
			Code:    r.Code,
			Message: r.Msg,
			Data:    val,
		},
	)
}

// HTTP code + code + msg + data
func ReturnHttpResponse(c *gin.Context, httpcode int, code int, msg string, data any) {
	c.JSON(httpcode, Response[any]{
		Code:    code,
		Message: msg,
		Data:    data,
	})
}

func ReturnResponse(c *gin.Context, r global.Result, data any) {
	ReturnHttpResponse(c, http.StatusOK, r.Code, r.Msg, data)
}

func ReturnSuccess(c *gin.Context, data any) {
	ReturnResponse(c, global.OkReresult, data)
}

func ReturnFail(c *gin.Context, data any) {
	ReturnResponse(c, global.FailResult, data)
}
