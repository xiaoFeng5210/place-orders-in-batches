package service

import (
	"encoding/json"
	"io"
	"order-mock/utils"
	"os"

	"order-mock/model"

	"errors"

	lo "github.com/samber/lo"
	lom "github.com/samber/lo/mutable"
	"go.uber.org/zap"
)

func LoadOrderConfig() ([]model.OrderRequest, error) {
	jsonFile, err := os.Open("mock_data_config.json")
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()

	fileData, err := io.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}

	var currentOrderConfigs []model.OrderRequest
	err = json.Unmarshal(fileData, &currentOrderConfigs)
	if err != nil {
		return nil, err
	}
	return currentOrderConfigs, nil
}

func SaveOrderConfig(orderConfigs []model.OrderRequest) error {
	jsonData, err := json.Marshal(orderConfigs)
	if err != nil {
		utils.Logger.Error("Failed to marshal order configs: %v", zap.Error(err))
		return err
	}

	err = os.WriteFile("mock_data_config.json", jsonData, 0644)
	if err != nil {
		utils.Logger.Error("Failed to write order configs to file: %v", zap.Error(err))
		return err
	}
	return nil
}

func AddDealerConfig(newConfig model.OrderRequest) error {
	currentOrderConfigs, err := LoadOrderConfig()
	if err != nil {
		return err
	}

	_, ok := lo.Find(currentOrderConfigs, func(item model.OrderRequest) bool {
		return item.Dealer == newConfig.Dealer
	})

	if ok {
		return errors.New("门店已存在")
	}

	currentOrderConfigs = append(currentOrderConfigs, newConfig)
	err = SaveOrderConfig(currentOrderConfigs)
	if err != nil {
		return err
	}
	return nil
}

func UpdateDealerConfig(config model.OrderRequest) error {
	localeConfig, err := LoadOrderConfig()
	if err != nil {
		return err
	}
	if len(localeConfig) == 0 {
		return errors.New("门店配置不存在")
	}
	updatedConfig := lo.Map(localeConfig, func(item model.OrderRequest, _ int) model.OrderRequest {
		if item.Dealer == config.Dealer {
			return config
		} else {
			return item
		}
	})
	err = SaveOrderConfig(updatedConfig)
	if err != nil {
		return err
	}
	return nil
}

func DeleteOrderConfigItem(dealerId string) error {
	readed, err := LoadOrderConfig()
	if err != nil {
		return err
	}
	kept := lom.Filter(readed, func(item model.OrderRequest) bool {
		return item.Dealer != dealerId
	})
	err = SaveOrderConfig(kept)
	if err != nil {
		return err
	}
	return nil
}
