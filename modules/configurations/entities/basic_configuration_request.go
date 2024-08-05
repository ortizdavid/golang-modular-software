package entities

type UpdateBasicConfigurationRequest struct {
	AppName      string `json:"app_name" form:"app_name"`
	AppAcronym   string `json:"app_acronym" form:"app_acronym"`
	MaxAdmninUsers		int `json:"max_admin_users" form:"max_admin_users"`
	MaxSuperAdminUsers	int `json:"max_super_admin_users" form:"max_super_admin_users"`
	MaxRecordPerPage	int `json:"max_record_per_page" form:"max_record_per_page"`
}

func (req UpdateBasicConfigurationRequest) Validate() error {
	return nil
}
