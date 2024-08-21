package entities

import "time"

type Currency struct {
	CurrencyId   int       `gorm:"autoIncrement;primaryKey;column:currency_id"`
	CurrencyName string    `gorm:"column:currency_name;unique"`
	Code         string    `gorm:"column:code;size:3"`
	UniqueId     string    `gorm:"column:unique_id;unique;default:uuid_generate_v4()::text"`
	CreatedAt    time.Time `gorm:"column:created_at;default:now()"`
	UpdatedAt    time.Time `gorm:"column:updated_at;default:now()"`
}

func (Currency) TableName() string {
	return "reference.currencies"
}
