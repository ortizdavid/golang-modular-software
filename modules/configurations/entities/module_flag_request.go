package entities

// -- Module Flag
type ManageModuleFlagRequest struct {
    ModuleId  int `json:"module_id" form:"module_id"`
    Status    string `json:"status" form:"status"` 
}

func (req ManageModuleFlagRequest) Validate() error {
    return nil
}

// --- Search
type SearchModuleFlagRequest struct {
	SearchParam		string `json:"search_param" form:"search_param"`
}
