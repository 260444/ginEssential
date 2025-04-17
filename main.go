package main

import (
	"github.com/260444/ginEssential/common"
	"github.com/gin-gonic/gin"
)

func main() {
	common.InitDB()
	// 初始化数据库连接
	router := gin.Default()
	router = CollectRouter(router) // 注册路由
	router.Run()                   // 监听并在 0.0.0.0:8080 上启动服务器
}

// curl 117.72.11.171:8080/ping
