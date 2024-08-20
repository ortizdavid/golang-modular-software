package entities

// ---- Create
type CreatePolicyRequest struct {
	CompanyId      	int `json:"company_id" form:"company_id"`
    PolicyName   	string `json:"policy_name" form:"policy_name"`
    Description  	string `json:"description" form:"description"`
    EffectiveDate 	string `json:"effective_date" form:"effective_date"`
}

func (req CreatePolicyRequest) Validate() error  {
	return nil
}

// ---- Update
type UpdatePolicyRequest struct {
    CompanyId      	int `json:"company_id" form:"company_id"`
    PolicyName   	string `json:"policy_name" form:"policy_name"`
    Description  	string `json:"description" form:"description"`
    EffectiveDate 	string `json:"effective_date" form:"effective_date"`
}

func (req UpdatePolicyRequest) Validate() error  {
	return nil
}

// --- Search
type SearchPolicyRequest struct {
	SearchParam		string `json:"search_param" form:"search_param"`
}
