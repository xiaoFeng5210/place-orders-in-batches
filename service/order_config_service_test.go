package service

import "testing"

func TestLoadOrderConfig(t *testing.T) {
	orderConfigs, err := LoadOrderConfig()
	if err != nil {
		t.Fatalf("LoadOrderConfig failed: %v", err)
	}
	t.Logf("LoadOrderConfig success! orderConfigs: %v", orderConfigs)
	t.Logf("第一个productId: %s", orderConfigs[0].Products[0].ProductId)
}

func TestSaveOrderConfig(t *testing.T) {
	orderConfigs, err := LoadOrderConfig()
	if err != nil {
		t.Fatalf("LoadOrderConfig failed: %v", err)
	}
	orderConfigs = append(orderConfigs, OrderRequest{
		Dealer: "68db351696b56fcecafbc99",
		Products: []Product{
			{
				ProductId: "68db479a96b56fcecafbc480",
				Name:      "测试食物",
				Quantity:  1,
			},
		},
		Amount:  0,
		PayType: "",
	})
	err = SaveOrderConfig(orderConfigs)
	if err != nil {
		t.Fatalf("SaveOrderConfig failed: %v", err)
	}
	t.Logf("SaveOrderConfig success! orderConfigs: %v", orderConfigs)
}

func TestDeleteOrderConfigItem(t *testing.T) {
	orderConfigs, _ := LoadOrderConfig()
	err := DeleteOrderConfigItem(orderConfigs, "68db351696b56fcecafbc99")
	if err != nil {
		t.Fatalf("DeleteOrderConfigItem failed: %v", err)
	}
	t.Logf("DeleteOrderConfigItem success!")
}
