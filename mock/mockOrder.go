package mock

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"order-mock/utils"

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

// 模拟下单请求
func MockOrder(orderUrl string) (string, error) {
	requestUrl := fmt.Sprintf("%s/robotapi/order/add", orderUrl)
	request := &OrderRequest{
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
	}
	body, err := json.Marshal(request)
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
func MockCloseOrder(orderId string) error {
	requestUrl := fmt.Sprintf("%s/robotapi/order/close", Url)
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
