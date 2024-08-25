package entities

//--CREATE
type CreateFeatureRequest struct {
    ModuleId    int    `json:":module_id" form:"module_id"`
    Code        string   `json:":code" form:"code"`
    FeatureName  string    `json:":feature_name" form:"feature_name"`
    Description string    `json:":description" form:"description"`
}

func (req CreateFeatureRequest) Validate() error {
	return nil
}

//--CREATE
type UpdateFeatureRequest struct {
    ModuleId    int    `json:":module_id" form:"module_id"`
    Code        string   `json:":code" form:"code"`
    FeatureName  string    `json:":feature_name" form:"feature_name"`
    Description string    `json:":description" form:"description"`
}

func (req UpdateFeatureRequest) Validate() error {
	return nil
}

// --- Search
type SearchFeatureRequest struct {
	SearchParam		string `json:"search_param" form:"search_param"`
}
