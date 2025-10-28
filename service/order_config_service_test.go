package service

import (
	"testing"
)

func TestDeleteOrderConfigItem(t *testing.T) {
	err := DeleteOrderConfigItem("68db351696b56fcecafbc99")
	if err != nil {
		t.Fatalf("DeleteOrderConfigItem failed: %v", err)
	}
	t.Logf("DeleteOrderConfigItem success!")
}
