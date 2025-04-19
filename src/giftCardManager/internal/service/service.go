package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetWalletId(userId string) (string, error) {
	
	url := fmt.Sprintf("http://127.0.0.1:8081/wallet/%s", userId)

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error:", err)
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to get wallet id, status code: %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return "", err
	}

	var response struct {
		WalletId string `json:"walletId"`
	}

	if err := json.Unmarshal(body, &response); err != nil {
		return "", fmt.Errorf("failed to unmarshal response body: %v", err)
	}

	fmt.Println("Wallet ID:", response.WalletId)

	return response.WalletId, nil

}
