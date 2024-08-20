package entities

// --- Create
type CreateProjectRequest struct {
	CompanyId   int `json:"company_id" form:"company_id"`
	ProjectName string  `json:"project_name" form:"project_name"`
	Description string  `json:"description" form:"description"`
	StartDate  	string `json:"start_date" form:"start_date"`
	EndDate     string `json:"end_date" form:"end_date"`
	Status      string  `json:"status" form:"status"`
}

func (req CreateProjectRequest) Validate() error  {
	return nil
}

// --- Update
type UpdateProjectRequest struct {
	CompanyId   int `json:"company_id" form:"company_id"`
	ProjectName string  `json:"project_name" form:"project_name"`
	Description string  `json:"description" form:"description"`
	StartDate  	string `json:"start_date" form:"start_date"`
	EndDate     string `json:"end_date" form:"end_date"`
	Status      string  `json:"status" form:"status"`
}

func (req UpdateProjectRequest) Validate() error  {
	return nil
}


// --- Search
type SearchProjectRequest struct {
	SearchParam		string `json:"search_param" form:"search_param"`
}

