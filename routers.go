package main

import (
	"github.com/260444/ginEssential/controller"
	"github.com/260444/ginEssential/middleware"
	"github.com/gin-gonic/gin"
)

func CollectRouter(router *gin.Engine) *gin.Engine {
	router.POST("/api/auth/register", controller.Register)
	router.POST("/api/auth/login", controller.Login)
	router.GET("/api/auth/info", middleware.AuthMiddleware(), controller.Info)

	return router
}
