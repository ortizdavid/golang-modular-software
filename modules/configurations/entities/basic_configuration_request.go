package entities

type UpdateBasicConfigurationRequest struct {
	MaxAdmninUsers		int `json:"max_admin_users"`
	MaxSuperAdminUsers	int `json:"max_super_admin_users"`
	MaxRecordPerPage	int `json:"max_record_per_page"`
}

func (req UpdateBasicConfigurationRequest) Validate() error {
	return nil
}
