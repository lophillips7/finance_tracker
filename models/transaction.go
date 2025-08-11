package models

import (
	"time"

	"gorm.io/gorm"
)

type Transaction struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	Amount      float64        `json:"amount" gorm:"not null"`
	Description string         `json:"description"`
	Category    string         `json:"category"`
	Type        string         `json:"type"` // "income" or "expense"
	Date        time.Time      `json:"date"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`
}
