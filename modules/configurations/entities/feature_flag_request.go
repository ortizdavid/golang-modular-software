package entities

// -- Feature flag
type ManageFeatureFlagRequest struct {
    FeatureId  int `json:"feature_id" form:"feature_id"`
    Status    string `json:"status" form:"status"` 
}

func (req ManageFeatureFlagRequest) Validate() error {
    return nil
}

// --- Search
type SearchFeatureFlagRequest struct {
	SearchParam		string `json:"search_param" form:"search_param"`
}
