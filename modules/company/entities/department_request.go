package entities

// ---- Create
type CreateDepartmentRequest struct {
	CompanyId      int `json:"company_id" form:"company_id"`
	DepartmentName string `json:"department_name" form:"department_name"`
	Acronym        string `json:"acronym" form:"acronym"`
	Description    string `json:"description" form:"description"`
}

func (req CreateDepartmentRequest) Validate() error  {
	return nil
}

// ---- Update
type UpdateDepartmentRequest struct {
	CompanyId      int `json:"company_id" form:"company_id"`
	DepartmentName string `json:"department_name" form:"department_name"`
	Acronym        string `json:"acronym" form:"acronym"`
	Description    string `json:"description" form:"description"`
}

func (req UpdateDepartmentRequest) Validate() error  {
	return nil
}


// --- Search
type SearchDepartmentRequest struct {
	SearchParam		string `json:"search_param" form:"search_param"`
}
