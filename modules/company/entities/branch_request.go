package entities

// ---- Create
type CreateBranchRequest struct {
	CompanyId    int `json:"company_id" form:"company_id"`
    BranchName   string `json:"branch_name" form:"branch_name"`
    Code         string `json:"code" form:"code"`
    Address      string `json:"address" form:"address"`
    Phone        string `json:"phone" form:"phone"`
    Email        string `json:"email" form:"email"`
}

func (req CreateBranchRequest) Validate() error  {
	return nil
}

// ---- Update
type UpdateBranchRequest struct {
	CompanyId    int `json:"company_id" form:"company_id"`
    Code         string `json:"code" form:"code"`
    BranchName   string `json:"branch_name" form:"branch_name"`
    Address      string `json:"address" form:"address"`
    Phone        string `json:"phone" form:"phone"`
    Email        string `json:"email" form:"email"`
}

func (req UpdateBranchRequest) Validate() error  {
	return nil
}


// --- Search
type SearchBranchRequest struct {
	SearchParam		string `json:"search_param" form:"search_param"`
}
