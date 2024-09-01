package entities

//--CREATE
type CreateCoreEntityRequest struct {
    ModuleId    int    `json:"module_id" form:"module_id"`
    Code        string   `json:"code" form:"code"`
    EntityName  string    `json:"entity_name" form:"entity_name"`
    Description string    `json:"description" form:"description"`
}

func (req CreateCoreEntityRequest) Validate() error {
	return nil
}

//--CREATE
type UpdateCoreEntityRequest struct {
    ModuleId    int    `json:"module_id" form:"module_id"`
    Code        string   `json:"code" form:"code"`
    EntityName  string    `json:"entity_name" form:"entity_name"`
    Description string    `json:"description" form:"description"`
}

func (req UpdateCoreEntityRequest) Validate() error {
	return nil
}

// --- Search
type SearchCoreEntityRequest struct {
	SearchParam		string `json:"search_param" form:"search_param"`
}
