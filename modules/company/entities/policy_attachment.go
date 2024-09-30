package entities

import (
	shared "github.com/ortizdavid/golang-modular-software/modules/shared/entities"
)

type PolicyAttachment struct {
	AttachmentId      	int `gorm:"primaryKey;autoIncrement"`
	PolicyId     		int `gorm:"column:policy_id"`
	CompanyId     		int `gorm:"column:company_id"`
	AttachmentName    	string `gorm:"column:attachment_name"`
	FileName    		string `gorm:"column:file_name"`
	shared.BaseEntity
}

func (PolicyAttachment) TableName() string {
	return "company.policy_attachments"
}