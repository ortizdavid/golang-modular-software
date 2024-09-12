package entities

import (
	shared "github.com/ortizdavid/golang-modular-software/modules/shared/entities"
)

type Room struct {
	RoomId    int `gorm:"primaryKey;autoIncrement"`
	CompanyId  int `gorm:"column:company_id"`
	BranchId  int `gorm:"column:branch_id"`
	RoomName  string `gorm:"column:room_name"`
	Number    string `gorm:"column:number"`
	Capacity  int `gorm:"column:capacity"`
	shared.BaseEntity
}

func (Room) TableName() string {
	return "company.rooms"
}

type RoomData struct {
	RoomId    	int `json:"room_id"`
	RoomName  	string `json:"room_name"`
	Number    	string `json:"number"`
	Capacity  	int `json:"capacity"`
	UniqueId  	string `json:"unique_id"`
	CreatedAt 	string `json:"created_at"`
	UpdatedAt 	string `json:"updated_at"`
	CompanyId  		int `json:"company_id"`
	CompanyName  	string `json:"company_name"`
	BranchId  	int `json:"branch_id"`
	BranchName  string `json:"branch_name"`
}
