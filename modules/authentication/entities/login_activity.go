package entities

import "time"

type LoginActivity struct {
	LoginId     int64              `gorm:"primaryKey;autoIncrement"`
    UserId      int64              `gorm:"column:user_id"`
    Status      LoginActivityStatus `gorm:"column:status"`
    Host        string             `gorm:"column:host"`
    Browser     string             `gorm:"column:browser"`
    IPAddress   string             `gorm:"column:ip_address"`
    Device      string             `gorm:"column:device"`
    Location    string             `gorm:"column:location"`
    LastLogin   time.Time          `gorm:"column:last_login"`
    LastLogout  time.Time          `gorm:"column:last_logout"`
    TotalLogin  int64          `gorm:"column:total_login"`
    TotalLogout int64          `gorm:"column:total_logout"`
    UniqueId    string             `gorm:"column:unique_id"`
    CreatedAt   time.Time          `gorm:"column:created_at"`
    UpdatedAt   time.Time          `gorm:"column:updated_at"`
}

func (LoginActivity) TableName() string {
	return "authentication.login_activity"
}


type LoginActivityStatus string
const (
	ActivityStatusOffline LoginActivityStatus = "Offline"
	ActivityStatusOnline LoginActivityStatus = "Online"
)