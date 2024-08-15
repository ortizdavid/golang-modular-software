package entities

import "time"

type Room struct {
	RoomID    uint      `gorm:"primaryKey;autoIncrement"`
	BranchID  uint      `gorm:"column:branch_id;not null"`
	RoomName  string    `gorm:"column:room_name"`
	Number    string    `gorm:"column:number"`
	Capacity  int       `gorm:"column:capacity;not null"`
	UniqueID  string    `gorm:"column:unique_id"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

func (Room) TableName() string {
	return "company.rooms"
}