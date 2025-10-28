package main

import (
	"os"

	handler "order-mock/handler/gin"
	"order-mock/utils"

	"order-mock/service"

	"github.com/joho/godotenv"
	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

func main() {
	utils.InitZap()
	err := godotenv.Load()
	if err != nil {
		utils.Logger.Error("Failed to load .env file", zap.Error(err))
	}
	orderUrl := os.Getenv("ORDER_URL")
	if orderUrl == "" {
		utils.Logger.Fatal("ORDER_URL environment variable is required")
	}
	go service.LoopAndMock(orderUrl)

	engine := gin.Default()

	engine.Static("/assets", "web/dist/assets")
	engine.StaticFile("/", "web/dist/index.html")

	apiGroup := engine.Group("/api")

	apiGroup.GET("/mock_config", handler.GetMockConfigHandler)
	apiGroup.POST("/add_dealer", handler.AddDealerConfigHandler)
	apiGroup.POST("/delete_dealer_config", handler.DeleteDealerConfigHandler)

	utils.Logger.Info("开始启动服务, 端口: 3031")
	utils.Logger.Info("环境变量", zap.String("ORDER_URL", orderUrl))

	err = engine.Run("0.0.0.0:3031")

	if err != nil {
		utils.Logger.Error("服务启动失败", zap.Error(err))
		panic(err)
	}
}
