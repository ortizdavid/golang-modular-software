package entities

import (
	shared "github.com/ortizdavid/golang-modular-software/modules/shared/entities"
)

type Currency struct {
	CurrencyId   int    `gorm:"autoIncrement;primaryKey" json:"currency_id"`
	CurrencyName string `gorm:"column:currency_name" json:"currency_name"`
	Code         string `gorm:"column:code" json:"code"`
	Symbol       string `gorm:"column:symbol" json:"symbol"`
	shared.BaseEntity
}

func (Currency) TableName() string {
	return "reference.currencies"
}
