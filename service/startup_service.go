package service

import (
	"order-mock/utils"
	"time"

	"order-mock/mock"

	"go.uber.org/zap"
)

const (
	MockCount    = 18
	MockInterval = 5 * time.Second
)

func LoopAndMock(orderUrl string) {
	for {
		currentTime := time.Now()
		if currentTime.Hour() == 7 && currentTime.Minute() == 50 {
			utils.Logger.Info("开始执行订单模拟，当前时间: ", zap.Int("hour", currentTime.Hour()), zap.Int("minute", currentTime.Minute()))
			orderConfigResult, err := LoadOrderConfig()
			if err != nil {
				utils.Logger.Error("读取订单配置失败", zap.Error(err))
				time.Sleep(10 * time.Second)
				continue
			}
			if len(orderConfigResult) > 0 {
				for i := 0; i < MockCount; i++ {
					mock.MockAllAndCloseWithConfig(orderUrl, orderConfigResult)
				}
				utils.Logger.Info("下单全部结束")
			}
		}
		time.Sleep(5 * time.Second)
	}
}
