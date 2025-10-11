package mock

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"order-mock/utils"

	"time"

	"go.uber.org/zap"
)

type Product struct {
	ProductId string `json:"product_id"`
	Name      string `json:"name"`
	Quantity  int    `json:"quantity"`
}

type OrderRequest struct {
	Dealer   string    `json:"dealer"`
	Products []Product `json:"products"`
	Amount   float64   `json:"amount"`
	PayType  string    `json:"pay_type"`
}

type OrderRespData struct {
	Id string `json:"_id"`
}

type OrderResponse struct {
	Code int           `json:"code"`
	Data OrderRespData `json:"data"`
}

var allNeedMockData = []OrderRequest{
	{
		// 如家酒店
		Dealer: "68c39717ceb6fa9a057abd00",
		Products: []Product{
			{
				ProductId: "68c3e86dceb6fa9a057abfdc",
				Name:      "面条1",
				Quantity:  1,
			},
		},
		Amount:  0,
		PayType: "",
	},

	{
		Dealer: "68db351696b56fcecafbc433",
		Products: []Product{
			{
				ProductId: "68db479a96b56fcecafbc4a4",
				Name:      "面条2",
				Quantity:  1,
			},
		},
		Amount:  0,
		PayType: "",
	},
}

func MockAllAndClose(orderUrl string) {
	for _, data := range allNeedMockData {
		orderId, err := MockOrder(orderUrl, data)
		if err != nil {
			utils.Logger.Error("模拟下单失败！", zap.Error(err))
			continue
		}
		time.Sleep(10 * time.Second)
		err = MockCloseOrder(orderUrl, orderId)
		if err != nil {
			utils.Logger.Error("模拟关闭订单失败！", zap.Error(err))
		}
	}
}

// 模拟下单请求
func MockOrder(orderUrl string, data OrderRequest) (string, error) {
	requestUrl := fmt.Sprintf("%s/robotapi/order/add", orderUrl)
	body, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	resp, err := http.Post(requestUrl, "application/json", bytes.NewBuffer(body))
	if err != nil {
		utils.Logger.Error("请求直接失败", zap.Error(err))
		return "", err
	}
	defer resp.Body.Close()
	body, err = io.ReadAll(resp.Body)
	if err != nil {
		utils.Logger.Error("读取响应体失败", zap.Error(err))
		return "", err
	}
	var orderResp OrderResponse
	err = json.Unmarshal(body, &orderResp)
	if err != nil {
		utils.Logger.Error("解析响应体失败", zap.Error(err))
		return "", err
	}
	if orderResp.Code != 0 {
		utils.Logger.Error("mock order Code != 0", zap.String("response", string(body)))
		return "", errors.New("mock order Code != 0")
	}
	return orderResp.Data.Id, nil
}

type CloseOrderRequest struct {
	Id string `json:"id"`
}

type CloseOrderResponse struct {
	Code int `json:"code"`
}

// 模拟订单制作完成
func MockCloseOrder(orderUrl, orderId string) error {
	requestUrl := fmt.Sprintf("%s/robotapi/order/close", orderUrl)
	request := &CloseOrderRequest{
		Id: orderId,
	}
	body, err := json.Marshal(request)
	if err != nil {
		utils.Logger.Error("Failed to marshal mock close order request: %v", zap.Error(err))
		return err
	}

	resp, err := http.Post(requestUrl, "application/json", bytes.NewBuffer(body))
	if err != nil {
		utils.Logger.Error("Failed to post mock close order request: %v", zap.Error(err))
		return err
	}
	defer resp.Body.Close()
	body, err = io.ReadAll(resp.Body)
	if err != nil {
		utils.Logger.Error("Failed to read response body: %v", zap.Error(err))
		return err
	}

	var closeOrderResp CloseOrderResponse
	err = json.Unmarshal(body, &closeOrderResp)
	if err != nil {
		utils.Logger.Error("Failed to unmarshal response body: %v", zap.Error(err))
		return err
	}
	if closeOrderResp.Code != 0 {
		utils.Logger.Error("Failed to mock close order: %s", zap.String("response", string(body)))
		return errors.New("mock close order Code != 0")
	}
	return nil
}
