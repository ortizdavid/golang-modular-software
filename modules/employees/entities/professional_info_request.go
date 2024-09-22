package entities

// -- Create
type CreateProfessionalInfoRequest struct {
	EmployeeId				int64 `json:"employee_id" form:"employee_id"`
	MaritalStatusId			int `json:"marital_status_id" form:"marital_status_id"`
	DepartmentId			int `json:"department_id" form:"department_id"`
	JobTitleId				int `json:"job_title_id" form:"job_title_id"`
	EmploymentStatusId		int `json:"employment_status_id" form:"employment_status_id"`
}

func (req CreateProfessionalInfoRequest) Validate() error {
	return nil
}

// -- Update
type UpdateProfessionalInfoRequest struct {
	EmployeeId				int64 `json:"employee_id" form:"employee_id"`
	MaritalStatusId			int `json:"marital_status_id" form:"marital_status_id"`
	DepartmentId			int `json:"department_id" form:"department_id"`
	JobTitleId				int `json:"job_title_id" form:"job_title_id"`
	EmploymentStatusId		int `json:"employment_status_id" form:"employment_status_id"`
}

func (req UpdateProfessionalInfoRequest) Validate() error {
	return nil
}

type SearchProfessionalInfoRequest struct {
	SearchParam string `json:"search_param" form:"search_param"`
}