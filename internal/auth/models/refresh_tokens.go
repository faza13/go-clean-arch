package models

import (
	"base/utils/uuid"
	"time"
)

type RefreshToken struct {
	ID        uuid.BinaryUUID `json:"id"`
	Uuid      string          `json:"uuid"`
	UserId    uint64          `json:"user_id"`
	UserAgent string          `json:"user_agent"`
	Platform  string          `json:"platform"`
	Counter   uint8           `json:"counter"`
	ExpiredAt time.Time       `json:"expired_at"`
	CreatedAt time.Time       `json:"created_at"`
	UpdatedAt time.Time       `json:"updated_at"`
}
