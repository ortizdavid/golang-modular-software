package entities

import (
	shared "github.com/ortizdavid/golang-modular-software/modules/shared/entities"
)

type ProjectAttachment struct {
	AttachmentId      	int `gorm:"primaryKey;autoIncrement"`
	ProjectId     		int `gorm:"column:project_id"`
	CompanyId     		int `gorm:"column:company_id"`
	AttachmentName    	string `gorm:"column:attachment_name"`
	FileName    		string `gorm:"column:file_name"`
	shared.BaseEntity
}

func (ProjectAttachment) TableName() string {
	return "company.project_attachments"
}