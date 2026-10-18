package service

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"giftpay/backend/logger"
	"giftpay/backend/models"
	"giftpay/backend/repository"
)

type GiftPayService interface {
	ApplyDiscountCode(ctx context.Context, codeStr, userID string) (*models.ApplyGiftCodeResponse, error)
}

type giftPayService struct {
	repo       repository.GiftPayRepository
	walletURL  string
}

func NewGiftPayService(repo repository.GiftPayRepository, walletURL string) GiftPayService {
	return &giftPayService{
		repo:      repo,
		walletURL: walletURL,
	}
}

func (s *giftPayService) ApplyDiscountCode(ctx context.Context, codeStr, userID string) (*models.ApplyGiftCodeResponse, error) {
	log := logger.GetLogger(ctx)

	// 1. Fetch Code
	code, err := s.repo.GetCode(ctx, codeStr)
	if err != nil {
		log.Warn("Gift code not found or error fetching", "action", "APPLY_GIFT_CODE", "code", logger.Sanitize(codeStr), "userID", userID, "error", err)
		return nil, err
	}

	// 2. Validate Code Constraints
	if err := s.validateCode(code); err != nil {
		log.Warn("Gift code validation failed", "action", "APPLY_GIFT_CODE", "code", logger.Sanitize(codeStr), "userID", userID, "error", err)
		return nil, err
	}

	// 3. Check if user already used
	used, err := s.repo.HasUserUsedCode(ctx, code.ID, userID)
	if err != nil {
		log.Error("Error checking user code usage", "action", "APPLY_GIFT_CODE", "userID", userID, "error", err)
		return nil, fmt.Errorf("internal server error")
	}
	if used {
		log.Info("User already used code", "action", "APPLY_GIFT_CODE", "code", logger.Sanitize(codeStr), "userID", userID)
		return nil, repository.ErrCodeAlreadyUsed
	}

	// 4. Atomically apply in repo (creates ledger + increments use + updates wallet balance)
	// We assume ACID compliance is handled by the repository using the same DB connection.
	if err := s.repo.ApplyCode(ctx, code, userID); err != nil {
		log.Error("Failed to apply code and update balance in DB transaction", "action", "APPLY_GIFT_CODE", "userID", userID, "error", err)
		return nil, err
	}

	log.Info("Gift code successfully applied", "action", "APPLY_GIFT_CODE", "code", logger.Sanitize(codeStr), "userID", userID, "appliedValue", code.Value)

	return &models.ApplyGiftCodeResponse{
		Status:       "success",
		Message:      "Gift code applied successfully",
		AppliedValue: code.Value,
	}, nil
}

func (s *giftPayService) validateCode(code *models.GiftCode) error {
	if !code.IsActive {
		return repository.ErrCodeInactive
	}
	if code.ExpiresAt != nil && time.Now().After(*code.ExpiresAt) {
		return repository.ErrCodeExpired
	}
	if code.MaxUsage > 0 && code.CurrentUsage >= code.MaxUsage {
		return repository.ErrCodeUsageLimit
	}
	return nil
}

// Wallet Integrations
func (s *giftPayService) getWalletID(userID string) (string, error) {
	url := fmt.Sprintf("%s/wallet/%s", s.walletURL, userID)
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed status code: %d", resp.StatusCode)
	}

	var result struct {
		WalletID string `json:"walletId"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", err
	}
	return result.WalletID, nil
}

func (s *giftPayService) getWalletBalance(walletID string) (float64, error) {
	url := fmt.Sprintf("%s/wallet/%s/balance", s.walletURL, walletID)
	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("failed status code: %d", resp.StatusCode)
	}

	var result struct {
		Balance float64 `json:"balance"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return 0, err
	}
	return result.Balance, nil
}

func (s *giftPayService) addWalletBalance(walletID string, amount float64) (float64, error) {
	url := fmt.Sprintf("%s/wallet/%s/balance", s.walletURL, walletID)
	reqBody := map[string]float64{"balance": amount}
	jsonBody, _ := json.Marshal(reqBody)

	req, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return 0, err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("failed status code: %d", resp.StatusCode)
	}

	var result struct {
		NewBalance float64 `json:"newBalance"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return 0, err
	}
	return result.NewBalance, nil
}
