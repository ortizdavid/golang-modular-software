package entities

import (
	"time"

	shared "github.com/ortizdavid/golang-modular-software/modules/shared/entities"
)

type UserApiKey struct {
	ApiKeyId  int64 `gorm:"autoIncrement;primaryKey"`
	UserId    int64 `gorm:"column:user_id"`
	XUserId   string `gorm:"column:x_user_id"`
	XApiKey   string `gorm:"column:x_api_key"`
	IsActive  bool  `gorm:"column:is_active"`
	CreatedBy int64 `gorm:"column:created_by"`
	ExpiresAt time.Time `gorm:"column:expires_at"`
	shared.BaseEntity
}

func (UserApiKey) TableName() string {
	return "authentication.user_api_key"
}

type UserApiKeyData struct {
	ApiKeyId  	int64 `json:"api_key_id"`
	UniqueId	string `json:"unique_id"`
	XUserId     string `json:"x_user_id"`
	XApiKey     string `json:"x_api_key"`
	IsActive	string  `json:"is_active"`
	CreatedBy	int64 `json:"created_by"`
	ExpiresAt 	string `json:"expires_at"`
	CreatedAt 	string `json:"created_at"`
	UpdatedAt 	string `json:"updated_at"`
	UserId    	int64 `json:"user_id"`
	UserName    string `json:"user_name"`
	Email       string `json:"email"`
	Password    string `json:"password"` 
}
