package entities

type UpdateCompanyConfigurationRequest struct {
	CompanyName      string `json:"company_name"`
	CompanyAcronym   string `json:"company_acronym"`
	CompanyPhone     string `json:"company_phone"`
	CompanyEmail     string `json:"company_email"`
	CompanyMainColor string `json:"company_main_color"`
}

func (req UpdateCompanyConfigurationRequest) Validate() error {
	return nil
}
