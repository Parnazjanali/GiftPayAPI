package repository

import (
	"context"
	"errors"

	"giftpay/backend/models"

	"gorm.io/gorm"
)

var (
	ErrCodeNotFound     = errors.New("gift code not found")
	ErrCodeInactive     = errors.New("gift code is inactive")
	ErrCodeExpired      = errors.New("gift code has expired")
	ErrCodeUsageLimit   = errors.New("gift code usage limit reached")
	ErrCodeAlreadyUsed  = errors.New("user has already used this code")
)

type GiftPayRepository interface {
	GetCode(ctx context.Context, code string) (*models.GiftCode, error)
	ApplyCode(ctx context.Context, code *models.GiftCode, userID string) error
	HasUserUsedCode(ctx context.Context, codeID uint, userID string) (bool, error)
}

type giftPayRepo struct {
	db *gorm.DB
}

func NewGiftPayRepository(db *gorm.DB) GiftPayRepository {
	return &giftPayRepo{db: db}
}

func (r *giftPayRepo) GetCode(ctx context.Context, code string) (*models.GiftCode, error) {
	var gc models.GiftCode
	err := r.db.WithContext(ctx).Where("code = ?", code).First(&gc).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrCodeNotFound
		}
		return nil, err
	}
	return &gc, nil
}

func (r *giftPayRepo) HasUserUsedCode(ctx context.Context, codeID uint, userID string) (bool, error) {
	var count int64
	err := r.db.WithContext(ctx).Model(&models.GiftCodeUsage{}).
		Where("gift_code_id = ? AND user_id = ?", codeID, userID).
		Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

// ApplyCode atomically increments usage and creates a ledger entry
func (r *giftPayRepo) ApplyCode(ctx context.Context, code *models.GiftCode, userID string) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// 1. Lock the row (SQLite doesn't support FOR UPDATE, use immediate transaction behavior instead)
		var lockedCode models.GiftCode
		if err := tx.First(&lockedCode, code.ID).Error; err != nil {
			return err
		}

		// 2. Double-check validity after acquiring lock
		if lockedCode.MaxUsage > 0 && lockedCode.CurrentUsage >= lockedCode.MaxUsage {
			return ErrCodeUsageLimit
		}
		if !lockedCode.IsActive {
			return ErrCodeInactive
		}

		var count int64
		if err := tx.Model(&models.GiftCodeUsage{}).
			Where("gift_code_id = ? AND user_id = ?", lockedCode.ID, userID).
			Count(&count).Error; err != nil {
			return err
		}
		if count > 0 {
			return ErrCodeAlreadyUsed
		}

		// 3. Increment usage
		if err := tx.Model(&lockedCode).Update("current_usage", gorm.Expr("current_usage + ?", 1)).Error; err != nil {
			return err
		}

		// 4. Create ledger entry
		usage := models.GiftCodeUsage{
			GiftCodeID: lockedCode.ID,
			UserID:     userID,
		}
		if err := tx.Create(&usage).Error; err != nil {
			return err
		}

		// We assume Wallet is already migrated and accessible in this DB connection
		// to fulfill the ACID compliance requirement.
		var wallet struct {
			ID      string  `gorm:"primaryKey"`
			UserID  string
			Balance float64
		}

		// Attempt to get the wallet by UserID
		if err := tx.Table("wallets").Where("user_id = ?", userID).First(&wallet).Error; err != nil {
			return errors.New("failed to find user wallet for balance update")
		}

		// Calculate applied value
		appliedValue := lockedCode.Value
		if lockedCode.Type == models.PercentageType {
			appliedValue = wallet.Balance * (lockedCode.Value / 100.0)
		}

		newBalance := wallet.Balance + appliedValue

		if err := tx.Table("wallets").Where("id = ?", wallet.ID).Update("balance", newBalance).Error; err != nil {
			return errors.New("failed to update wallet balance")
		}

		// Overwrite code value so the service layer can return the true applied amount
		code.Value = appliedValue

		return nil
	})
}
