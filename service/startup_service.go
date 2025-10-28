package service

import (
	"fmt"
	"order-mock/model"
	"order-mock/utils"
	"time"

	"sync"

	"order-mock/mock"

	"go.uber.org/zap"
)

var (
	orderConfigs []model.OrderRequest
	lock         sync.Mutex
)

const (
	MockCount    = 18
	MockInterval = 5 * time.Second
)

func CheckMockDataConfig() {
	for {
		lock.Lock()
		_, err := LoadOrderConfig()
		lock.Unlock()
		if err != nil {
			utils.Logger.Error("CheckMockDataConfig Error", zap.Error(err))
			time.Sleep(5 * time.Second)
			continue
		}
		utils.Logger.Info("读取当前配置数据: ", zap.Any("orderConfigs", orderConfigs))
		time.Sleep(5 * 60 * time.Second)
	}
}

func LoopAndMock(orderUrl string) {
	for {
		currentTime := time.Now()
		fmt.Printf("当前时间: %v:%v\n", currentTime.Hour(), currentTime.Minute())
		if currentTime.Hour() == 7 && currentTime.Minute() == 50 {
			orderConfigResult, err := LoadOrderConfig()
			if err != nil {
				utils.Logger.Error("LoadOrderConfig Error", zap.Error(err))
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
