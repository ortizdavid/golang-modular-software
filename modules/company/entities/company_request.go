package entities

import "time"

// ------ Create --
// CreateCompanyRequest represents the request structure for creating a new company.
type CreateCompanyRequest struct {
	CompanyName    string    `json:"company_name" form:"company_name"`
    CompanyAcronym string    `json:"company_acronym" form:"company_acronym"`
    CompanyType    string    `json:"company_type" form:"company_type"`
    Industry       string    `json:"industry" form:"industry"`
    FoundedDate    time.Time `json:"founded_date" form:"founded_date"`
    Address        string    `json:"address" form:"address"`
    Phone          string    `json:"phone" form:"phone"`
    Email          string    `json:"email" form:"email"`
    WebsiteURL     string    `json:"website_url" form:"website_url"`
}

func (CreateCompanyRequest) Validate() error {
	return nil
}

// ------ Update --
// UpdateCompanyRequest represents the request structure for updating an existing company.
type UpdateCompanyRequest struct {
	CompanyName    string    `json:"company_name" form:"company_name"`
    CompanyAcronym string    `json:"company_acronym" form:"company_acronym"`
    CompanyType    string    `json:"company_type" form:"company_type"`
    Industry       string    `json:"industry" form:"industry"`
    FoundedDate    time.Time `json:"founded_date" form:"founded_date"`
    Address        string    `json:"address" form:"address"`
    Phone          string    `json:"phone" form:"phone"`
    Email          string    `json:"email" form:"email"`
    WebsiteURL     string    `json:"website_url" form:"website_url"`
}

func (UpdateCompanyRequest) Validate() error {
	return nil
}
