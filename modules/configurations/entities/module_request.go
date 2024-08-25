package entities

//--CREATE
type CreateModuleRequest struct {
    ModuleName  string    `json:":module_name" form:"module_name"`
    Description string    `json:":description" form:"description"`
}

func (req CreateModuleRequest) Validate() error {
	return nil
}


//--CREATE
type UpdateModuleRequest struct {
    ModuleName  string    `json:":module_name" form:"module_name"`
    Description string    `json:":description" form:"description"`
}

func (req UpdateModuleRequest) Validate() error {
	return nil
}

// --- Search
type SearchModuleRequest struct {
	SearchParam		string `json:"search_param" form:"search_param"`
}
