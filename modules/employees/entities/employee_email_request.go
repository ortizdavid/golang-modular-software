package entities

// -- Create 
type CreateEmployeePhoneRequest struct {
	EmployeeId		int64 `json:"employee_id" form:"employee_id"`
	ContactTypeId	int	`json:"contact_type_id" form:"contact_type_id"`
	DialingCode   	string `json:"dialing_code" form:"dialing_code"`
	PhoneNumber   	string `json:"phone_number" form:"phone_number"`
}

func (req CreateEmployeePhoneRequest) Validate() error {
	return nil
}

// -- Updaye 
type UpdateEmployeePhoneRequest struct {
	EmployeeId		int64 `json:"employee_id" form:"employee_id"`
	ContactTypeId	int	`json:"contact_type_id" form:"contact_type_id"`
	DialingCode   	string `json:"dialing_code" form:"dialing_code"`
	PhoneNumber   	string `json:"phone_number" form:"phone_number"`
}

func (req UpdateEmployeePhoneRequest) Validate() error {
	return nil
}

type SearchEmployeePhoneRequest struct {
	SearchParam string `json:"search_param" form:"search_param"`
}