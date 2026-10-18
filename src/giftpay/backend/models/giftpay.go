package models

import (
	"time"

	"gorm.io/gorm"
)

type DiscountType string

const (
	AmountType     DiscountType = "AMOUNT"
	PercentageType DiscountType = "PERCENTAGE"
)

type GiftCode struct {
	ID           uint           `gorm:"primaryKey" json:"id"`
	Code         string         `gorm:"uniqueIndex;not null" json:"code"`
	Type         DiscountType   `gorm:"not null" json:"type"`
	Value        float64        `gorm:"not null" json:"value"` // Either amount or percentage
	MaxUsage     int            `gorm:"default:1" json:"maxUsage"`
	CurrentUsage int            `gorm:"default:0" json:"currentUsage"`
	IsActive     bool           `gorm:"default:true" json:"isActive"`
	ExpiresAt    *time.Time     `json:"expiresAt"`
	CreatedAt    time.Time      `json:"createdAt"`
	UpdatedAt    time.Time      `json:"updatedAt"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"-"`
}

type GiftCodeUsage struct {
	ID         uint           `gorm:"primaryKey" json:"id"`
	GiftCodeID uint           `gorm:"not null;index" json:"giftCodeId"`
	UserID     string         `gorm:"not null;index" json:"userId"` // Keep generic string to align with wallet/giftcard userIds
	AppliedAt  time.Time      `gorm:"autoCreateTime" json:"appliedAt"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"-"`

	GiftCode   GiftCode       `gorm:"foreignKey:GiftCodeID" json:"-"`
}

type ApplyGiftCodeRequest struct {
	Code   string `json:"code" binding:"required"`
	UserID string `json:"userId" binding:"required"`
}

type ApplyGiftCodeResponse struct {
	Status        string  `json:"status"`
	Message       string  `json:"message"`
	AppliedValue  float64 `json:"appliedValue"`
	NewBalance    float64 `json:"newBalance,omitempty"`
}

type ErrorResponse struct {
	Error struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	} `json:"error"`
}
