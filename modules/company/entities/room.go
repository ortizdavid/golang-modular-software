package entities

import "time"

type Room struct {
	RoomId    int `gorm:"primaryKey;autoIncrement"`
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