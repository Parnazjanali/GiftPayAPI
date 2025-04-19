package Service

import (
	"apiGateway/internal/model"
	sqliteDB "apiGateway/repository/db/sqlite"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/dustin/go-humanize"
)

func AuthenticateUser(username, password string) (*model.User, error) {

	user, err := sqliteDB.AuthenticateUser(sqliteDB.DB, username, password)
	if err != nil {
		return nil, err
	}
	return user, nil

}
func GetWalletIdByUserId(userId string) (string, error) {

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
func AddCreditByWalletId(userId, giftCardId string) (string, error) {

	url := "http://127.0.0.1:8080/api/gift-cards/link"

	log.Println("Preparing payload for request")
	payLoad := map[string]string{
		"giftCardId": giftCardId,
	}
	payLoadBytes, err := json.Marshal(payLoad)
	if err != nil {
		log.Println("Error marshalling payload:", err)
		return "", fmt.Errorf("failed to marshal payload: %v", err)
	}

	log.Println("Creating HTTP request")
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payLoadBytes))
	if err != nil {
		log.Println("Error creating request:", err)
		return "", fmt.Errorf("failed to create request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("userId", userId)

	client := &http.Client{}
	log.Println("Sending request to API")
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error sending request:", err)
		return "", fmt.Errorf("failed to send request: %v", err)
	}
	defer resp.Body.Close()

	log.Println("Received response, status code:", resp.StatusCode)
	if resp.StatusCode != http.StatusOK {
		log.Println("Error: Unexpected status code:", resp.StatusCode)
		return "", fmt.Errorf("failed to activate gift card, status code: %d", resp.StatusCode)
	}

	var response struct {
		GiftCardId string `json:"giftCardId"`
		Amount     int    `json:"amount"`
		WalletId   string `json:"walletId"`
		Message    string `json:"message"`
	}

	log.Println("Decoding response body")
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		log.Println("Error decoding response body:", err)
		return "", fmt.Errorf("failed to decode response: %v", err)
	}

	log.Println("Decoded response:", response)

	return fmt.Sprintf("%d", response.Amount), nil
}
func AddBalance(walletId string, amount int) (string, error) {
	log.Println("🔁 Starting AddBalance function")
	log.Printf("📦 walletId: %s, amount: %d", walletId, amount)

	url := fmt.Sprintf("http://127.0.0.1:8081/wallet/%s/set-balance", walletId)
	log.Printf("🔗 URL: %s", url)

	payLoad := map[string]int{
		"balance": amount,
	}

	payLoadBytes, err := json.Marshal(payLoad)
	if err != nil {
		log.Printf("❌ Failed to marshal payload: %v", err)
		return "", fmt.Errorf("failed to marshal payload: %v", err)
	}
	log.Println("✅ Payload marshaled successfully")

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payLoadBytes))
	if err != nil {
		log.Printf("❌ Failed to create request: %v", err)
		return "", fmt.Errorf("failed to create request: %v", err)
	}
	log.Println("✅ HTTP request created")

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	log.Println("📡 Sending request to Wallet service")
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("❌ Failed to send request: %v", err)
		return "", fmt.Errorf("failed to send request: %v", err)
	}
	defer resp.Body.Close()

	log.Printf("📥 Received response, status code: %d", resp.StatusCode)
	if resp.StatusCode != http.StatusOK {
		log.Printf("❌ Failed to set balance, status code: %d", resp.StatusCode)
		return "", fmt.Errorf("failed to set balance, status code: %d", resp.StatusCode)
	}

	var response struct {
		Status     string  `json:"status"`
		Message    string  `json:"message"`
		NewBalance float64 `json:"newBalance"`
	}
	log.Println("🔄 Decoding response body")
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		log.Printf("❌ Failed to decode response: %v", err)
		return "", fmt.Errorf("failed to decode response: %v", err)
	}

	log.Printf("✅ Balance updated successfully, new balance: %f", response.NewBalance)

	// حالا شما می‌توانید موجودی جدید را در پیام قرار دهید
	message := fmt.Sprintf("%s ریال به کیف پول شما اضافه شد", humanize.Comma(int64(response.NewBalance)))

	return message, nil

}
