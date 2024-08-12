package entities

import "time"

type LoginActivity struct {
	LoginId     int64 `gorm:"primaryKey;autoIncrement"`
    UserId      int64 `gorm:"column:user_id"`
    Status      LoginActivityStatus `gorm:"column:status"`
    Host        string `gorm:"column:host"`
    Browser     string `gorm:"column:browser"`
    IPAddress   string `gorm:"column:ip_address"`
    Device      string `gorm:"column:device"`
    Location    string `gorm:"column:location"`
    LastLogin   time.Time `gorm:"column:last_login"`
    LastLogout  time.Time `gorm:"column:last_logout"`
    TotalLogin  int64 `gorm:"column:total_login"`
    TotalLogout int64 `gorm:"column:total_logout"`
    UniqueId    string `gorm:"column:unique_id"`
    CreatedAt   time.Time `gorm:"column:created_at"`
    UpdatedAt   time.Time `gorm:"column:updated_at"`
}

func (LoginActivity) TableName() string {
	return "authentication.login_activity"
}

type LoginActivityStatus string
const (
	ActivityStatusOffline LoginActivityStatus = "Offline"
	ActivityStatusOnline LoginActivityStatus = "Online"
)

type LoginActivityData struct {
    LoginId     int64 `json:"login_id"`
    UniqueId    string `json:"unique_id"`
    Status      LoginActivityStatus `json:"status"`
    Host        string `json:"host"`
    Browser     string `json:"browser"`
    IPAddress   string `json:"ip_address"`
    Device      string `json:"device"`
    Location    string `json:"location"`
    LastLogin   string `json:"last_login"`
    LastLogout  string `json:"last_logout"`
    TotalLogin  int64 `json:"total_login"`
    TotalLogout int64 `json:"total_logout"`
    CreatedAt   string `json:"created_at"`
    UpdatedAt   string `json:"updated_at"`
    UserId      int64 `json:"user_id"`
    UserName    string `json:"user_name"`
    Email       string `json:"email"`
}
