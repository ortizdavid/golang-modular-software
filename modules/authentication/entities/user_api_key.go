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
	ApiKeyId  	int64 `gorm:"autoIncrement;primaryKey"`
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

/*CREATE OR REPLACE VIEW authentication.view_user_api_key_data AS 
SELECT uak.api_key_id, uak.unique_id,
    uak.x_user_id, uak.x_api_key, 
    CASE WHEN uak.is_active THEN 'Yes' ELSE 'No' END AS is_active,
    uak.created_by,
    TO_CHAR(uak.expires_at, 'YYYY-MM-DD HH24:MI:SS') AS expires_at,
    TO_CHAR(uak.created_at, 'YYYY-MM-DD HH24:MI:SS') AS created_at,
    TO_CHAR(uak.updated_at, 'YYYY-MM-DD HH24:MI:SS') AS updated_at,
    us.user_id, us.user_name,
    us.email*/