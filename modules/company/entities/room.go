package entities

import "time"

type Room struct {
	RoomId    int `gorm:"primaryKey;autoIncrement"`
	CompanyId  int `gorm:"column:company_id"`
	BranchId  int `gorm:"column:branch_id"`
	RoomName  string `gorm:"column:room_name"`
	Number    string `gorm:"column:number"`
	Capacity  int `gorm:"column:capacity"`
	UniqueId  string `gorm:"column:unique_id"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

func (Room) TableName() string {
	return "company.rooms"
}

type RoomData struct {
	RoomId    int `json:"room_id"`
	RoomName  string `json:"room_name"`
	Number    string `json:"number"`
	Capacity  int `json:"capacity"`
	UniqueId  string `json:"unique_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	BranchId  int `json:"branch_id"`
	BranchName  string `json:"branch_name"`
}