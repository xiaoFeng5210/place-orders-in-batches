package main

import (
	"net/http"
	"os"

	"order-mock/utils"

	"fmt"

	"order-mock/service"

	"github.com/joho/godotenv"
	"go.uber.org/zap"
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

	go service.CheckMockDataConfig()
	go service.LoopAndMock(orderUrl)

	utils.Logger.Info("服务启动成功, 2025/10/11, 端口: 3031")
	utils.Logger.Info("环境变量", zap.String("ORDER_URL", orderUrl))
	fmt.Println("服务启动成功，端口: 3031")
	http.ListenAndServe(":3031", nil)
}
