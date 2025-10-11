package main

import (
	"net/http"
	"order-mock/mock"
	"os"
	"time"

	"order-mock/utils"

	"fmt"

	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

const (
	MockCount    = 18
	MockInterval = 5 * time.Second
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
	go func() {
		for {
			currentTime := time.Now()
			if currentTime.Hour() == 8 && currentTime.Minute() == 10 {
				for i := 0; i < MockCount; i++ {
					utils.Logger.Info("开始下单")
					mock.MockAllAndClose(orderUrl)
				}
				utils.Logger.Info("下单全部结束")
			}
			// * 每2秒查看一下当前时间
			time.Sleep(2 * time.Second)
		}
	}()

	utils.Logger.Info("服务启动成功, 2025/10/11, 端口: 3031")
	utils.Logger.Info("环境变量", zap.String("ORDER_URL", orderUrl))
	fmt.Println("服务启动成功，端口: 3031")
	http.ListenAndServe(":3031", nil)
}
