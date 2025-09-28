package mock

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Login() error {
	request := &LoginRequest{
		Username: "qingfeng.zhang@lebai.ltd",
		Password: "QINGFENG-lebai@saoihaiwheiq312ajak",
	}

	body, err := json.Marshal(request)
	if err != nil {
		fmt.Printf("Failed to marshal request: %v", err)
		return err
	}
	requestUrl := fmt.Sprintf("%s/auth/signin", Url)
	if resp, err := http.Post(requestUrl, "application/json", bytes.NewBuffer(body)); err != nil {
		fmt.Printf("Failed to post request: %v", err)
		return err
	} else {
		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("Failed to read response body: %v", err)
			return err
		}
		fmt.Printf("Login success: %s\n", string(body))
		return nil
	}
}
