package entities

// -- Create 
type CreateEmployeeEmailRequest struct {
	EmployeeId		int64 `json:"employee_id" form:"employee_id"`
	ContactTypeId	int	`json:"contact_type_id" form:"contact_type_id"`
	EmailAddress   	string `json:"email_address" form:"email_address"`
}

func (req CreateEmployeeEmailRequest) Validate() error {
	return nil
}

// -- Updaye 
type UpdateEmployeeEmailRequest struct {
	EmployeeId		int64 `json:"employee_id" form:"employee_id"`
	ContactTypeId	int	`json:"contact_type_id" form:"contact_type_id"`
	EmailAddress   	string `json:"email_address" form:"email_address"`
}

func (req UpdateEmployeeEmailRequest) Validate() error {
	return nil
}

type SearchEmployeeEmailRequest struct {
	SearchParam string `json:"search_param" form:"search_param"`
}