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

}
