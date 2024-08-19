package entities

// ---- Create
type CreateOfficeRequest struct {
	CompanyId    int `json:"company_id" form:"company_id"`
    OfficeName   string `json:"office_name" form:"office_name"`
    Code         string `json:"code" form:"code"`
    Address      string `json:"address" form:"address"`
    Phone        string `json:"phone" form:"phone"`
    Email        string `json:"email" form:"email"`
}

func (req CreateOfficeRequest) Validate() error  {
	return nil
}

// ---- Update
type UpdateOfficeRequest struct {
	CompanyId    int `json:"company_id" form:"company_id"`
    Code         string `json:"code" form:"code"`
    OfficeName   string `json:"office_name" form:"office_name"`
    Address      string `json:"address" form:"address"`
    Phone        string `json:"phone" form:"phone"`
    Email        string `json:"email" form:"email"`
}

func (req UpdateOfficeRequest) Validate() error  {
	return nil
}


// --- Search
type SearchOfficeRequest struct {
	SearchParam		string `json:"search_param" form:"search_param"`
}
