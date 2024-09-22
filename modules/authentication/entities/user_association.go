package entities

import (
	shared "github.com/ortizdavid/golang-modular-software/modules/shared/entities"
)

type UserAssociation struct {
    AssociationId int64 `gorm:"primaryKey;autoIncrement"`
    UserId     int64 `gorm:"column:user_id"`
    shared.BaseEntity
}

func (UserAssociation) TableName() string {
	return "authentication.user_associations"
}