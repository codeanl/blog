package main

import (
	"blog/internal"
	"blog/internal/global"
	"blog/middleware"
	"log"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {
	conf := global.InitConfig("./configs/config.yml")
	db := internal.InitDatabase(conf)
	rdb := internal.InitRedis(conf)

	//初始化gin
	gin.SetMode(conf.Server.Mode)
	r := gin.New()

	// 定义信任的代理服务器
	r.SetTrustedProxies([]string{"*"})
	r.Use(gin.Logger(), gin.Recovery())

	r.Use(middleware.CORS())
	r.Use(middleware.WithGormDB(db))
	r.Use(middleware.WithRDB(rdb))

	//注册hander
	internal.RegisterAllHandler(r)

	serverAddr := conf.Server.Port
	if serverAddr[0] == ':' || strings.HasPrefix(serverAddr, "0.0.0.0:") {
		log.Printf("Serving HTTP on (http://localhost:%s/) ... \n", strings.Split(serverAddr, ":")[1])
	} else {
		log.Printf("Serving HTTP on (http//%s/)\n", serverAddr)
	}

	r.Run(serverAddr)
}
