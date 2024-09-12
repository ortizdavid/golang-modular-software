package entities

import (
	shared "github.com/ortizdavid/golang-modular-software/modules/shared/entities"
)

type Currency struct {
	CurrencyId   int       `gorm:"autoIncrement;primaryKey"`
	CurrencyName string    `gorm:"column:currency_name"`
	Code         string    `gorm:"column:code"`
	Symbol       string  `gorm:"column:symbol"`
	shared.BaseEntity
}

func (Currency) TableName() string {
	return "reference.currencies"
}
