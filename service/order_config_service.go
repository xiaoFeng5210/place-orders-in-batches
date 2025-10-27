package service

import (
	"encoding/json"
	"io"
	"order-mock/utils"
	"os"

	"go.uber.org/zap"
)

func LoadOrderConfig() ([]OrderRequest, error) {
	jsonFile, err := os.Open("../mock_data_config.json")
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()

	fileData, err := io.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}

	var orderConfigs []OrderRequest
	err = json.Unmarshal(fileData, &orderConfigs)
	if err != nil {
		return nil, err
	}

	return orderConfigs, nil

}

func SaveOrderConfig(orderConfigs []OrderRequest) error {
	jsonData, err := json.Marshal(orderConfigs)
	if err != nil {
		utils.Logger.Error("Failed to marshal order configs: %v", zap.Error(err))
		return err
	}

	err = os.WriteFile("../mock_data_config.json", jsonData, 0644)
	if err != nil {
		utils.Logger.Error("Failed to write order configs to file: %v", zap.Error(err))
		return err
	}
	return nil
}

func DeleteOrderConfigItem(orderConfigs []OrderRequest, dealerId string) error {

}
