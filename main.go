package main

import (
	"fmt"
	"net/http"
	"order-mock/mock"
	"time"
)

const (
	MockCount    = 20
	MockInterval = 10 * time.Second
)

func main() {
	go func() {
		for {
			currentTime := time.Now()
			if currentTime.Hour() == 8 && currentTime.Minute() == 10 {
				for i := 0; i < MockCount; i++ {
					fmt.Println("开始下单，当前时间: ", currentTime)
					handleMock()
				}
				fmt.Println("下单全部结束, 当前时间: ", currentTime)
			}
			// * 每2秒查看一下当前时间
			time.Sleep(2 * time.Second)
		}
	}()

	fmt.Println("服务启动成功，端口: 3031")
	http.ListenAndServe(":3031", nil)
}

func handleMock() {
	orderId, err := mock.MockOrder()
	if err != nil {
		return
	}
	time.Sleep(MockInterval)
	if orderId != "" {
		err := mock.MockCloseOrder(orderId)
		if err != nil {
			return
		}
	} else {
		fmt.Println("OrderId is empty")
	}
}
