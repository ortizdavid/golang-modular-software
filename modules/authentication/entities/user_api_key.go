package entities

import "time"

type UserApiKey struct {
	ApiKeyId  int64  `gorm:"autoIncrement;primaryKey"`
	UserId    int64  `gorm:"column:user_id"`
	Key       string `gorm:"column:key"`
	IsActive  bool   `gorm:"column:is_active"`
	CreatedBy int64  `gorm:"column:created_by"`
	ExpiresAt time.Time `gorm:"column:expires_at"`
	UniqueId	string `gorm:"column:unique_id"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

func (UserApiKey) TableName() string {
	return "authentication.user_api_key"
}

type UserApiKeyData struct {
	ApiKeyId  	int64  `gorm:"autoIncrement;primaryKey"`
	UniqueId	string `json:"unique_id"`
	Key       	string `json:"key"`
	IsActive	string   `json:"is_active"`
	CreatedBy	int64  `json:"created_by"`
	ExpiresAt 	string `json:"expires_at"`
	CreatedAt 	string `json:"created_at"`
	UpdatedAt 	string `json:"updated_at"`
	UserId    	int64  `json:"user_id"`
	UserName    string `json:"user_name"`
	Email       string `json:"email"`
	Password    string `json:"password"` 
}