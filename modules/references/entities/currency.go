package entities

import "time"

type Currency struct {
	CurrencyId   int       `gorm:"autoIncrement;primaryKey"`
	CurrencyName string    `gorm:"column:currency_name"`
	Code         string    `gorm:"column:code"`
	Symbol       string  `gorm:"column:symbol"`
	UniqueId     string    `gorm:"column:unique_id"`
	CreatedAt    time.Time `gorm:"column:created_at"`
	UpdatedAt    time.Time `gorm:"column:updated_at"`
}

func (Currency) TableName() string {
	return "reference.currencies"
}
