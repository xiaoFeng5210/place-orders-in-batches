package mock

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func TestMockOrder(t *testing.T) {
	err := godotenv.Load("../.env")
	if err != nil {
		t.Fatal("Failed to load .env file", err)
	}
	orderUrl := os.Getenv("ORDER_URL")
	if orderUrl == "" {
		t.Fatal("ORDER_URL environment variable is required")
	}
	orderId, err := MockOrder(orderUrl, OrderRequest{
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
	})

	if err != nil {
		t.Fatal("MockOrder failed", err)
	}
	t.Log("MockOrder success!\n", orderId)
}

func TestMockCloseOrder(t *testing.T) {
	err := godotenv.Load("../.env")
	if err != nil {
		t.Fatal("Failed to load .env file", err)
	}
	orderUrl := os.Getenv("ORDER_URL")
	if orderUrl == "" {
		t.Fatal("ORDER_URL environment variable is required")
	}
	orderId := "68e9f73c96b56fcecafc190b"
	err = MockCloseOrder(orderUrl, orderId)
	if err != nil {
		t.Fatal("MockCloseOrder failed", err)
	}
	t.Log("MockCloseOrder success!", orderId)
}
