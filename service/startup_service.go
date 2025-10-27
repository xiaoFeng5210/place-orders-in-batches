package service

import (
	"fmt"
	"order-mock/mock"
	"order-mock/model"
	"time"
)

var (
	orderConfigs []model.OrderRequest
)

const (
	MockCount    = 18
	MockInterval = 5 * time.Second
)

func CheckMockDataConfig() {
	for {
		jsonData, err := LoadOrderConfig()
		if err != nil {
			fmt.Println("CheckMockDataConfig Error", err)
			time.Sleep(5 * time.Second)
			continue
		}

		orderConfigs = jsonData
		fmt.Println("CheckMockDataConfig", orderConfigs)

		time.Sleep(5 * 60 * time.Second)
	}
}

func LoopAndMock(orderUrl string) {
	for {
		currentTime := time.Now()
		if currentTime.Hour() == 15 && currentTime.Minute() == 43 {
			if len(orderConfigs) > 0 {
				for i := 0; i < MockCount; i++ {
					mock.MockAllAndCloseWithConfig(orderUrl, orderConfigs)
				}
			}
		}
		time.Sleep(5 * time.Second)
	}
}
