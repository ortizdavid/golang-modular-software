package entities


type AppConfiguration struct {
	BasicConfig BasicConfiguration `json:"basic_config"`
	CompanyConfig CompanyConfiguration `json:"company_config"`
	EmailConfig EmailConfiguration `json:"email_config"`
}
