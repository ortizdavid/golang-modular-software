package entities

// -- Create
type CreateEmployeeRequest struct {
	IdentificationTypeId	int `json:"identification_type_id" form:"identification_type_id"`
	CountryId				int `json:"country_id" form:"country_id"`
	MaritalStatusId			int `json:"marital_status_id" form:"marital_status_id"`
	DepartmentId			int `json:"department_id" form:"department_id"`
	JobTitleId				int `json:"job_title_id" form:"job_title_id"`
	EmploymentStatusId		int `json:"employment_status_id" form:"employment_status_id"`
	FirstName				string `json:"first_name" form:"first_name"`
	LastName				string `json:"last_name" form:"last_name"`
	IdentificationNumber	string `json:"identification_number" form:"identification_number"`
	Gender					string `json:"gender" form:"gender"`
	DateOfBirth				string `json:"date_of_birth" form:"date_of_birth"`
}

func (req CreateEmployeeRequest) Validate() error {
	return nil
}

// -- Update
type UpdateEmployeeRequest struct {
	IdentificationTypeId	int `json:"identification_type_id" form:"identification_type_id"`
	CountryId				int `json:"country_id" form:"country_id"`
	MaritalStatusId			int `json:"marital_status_id" form:"marital_status_id"`
	DepartmentId			int `json:"department_id" form:"department_id"`
	JobTitleId				int `json:"job_title_id" form:"job_title_id"`
	EmploymentStatusId		int `json:"employment_status_id" form:"employment_status_id"`
	FirstName				string `json:"first_name" form:"first_name"`
	LastName				string `json:"last_name" form:"last_name"`
	IdentificationNumber	string `json:"identification_number" form:"identification_number"`
	Gender					string `json:"gender" form:"gender"`
	DateOfBirth				string `json:"date_of_birth" form:"date_of_birth"`
}

func (req UpdateEmployeeRequest) Validate() error {
	return nil
}


type SearchEmployeeRequest struct {
	SearchParam string `json:"search_param" form:"search_param"`
}