package router

import (
	_ "gatesma.cn/gin-gorm-oj/docs"
	"gatesma.cn/gin-gorm-oj/service"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Router() *gin.Engine {
	r := gin.Default()

	// Swagger 配置
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	r.GET("/ping", service.Ping)
	// 问题
	r.GET("/problem-list", service.GetProblemList)

	return r
}
