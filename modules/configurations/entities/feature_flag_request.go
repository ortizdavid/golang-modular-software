package entities

// -- Module Flag
type ManageModuleFlagRequest struct {
    ModuleId  int `json:"module_id" form:"module_id"`
    Status    string `json:"status" form:"status"` 
}

func (req ManageModuleFlagRequest) Validate() error {
    return nil
}

// -- Feature flag
type ManageFeatureFlagRequest struct {
    FeatureId  int `json:"feature_id" form:"feature_id"`
    Status    string `json:"status" form:"status"` 
}

func (req ManageFeatureFlagRequest) Validate() error {
    return nil
}
