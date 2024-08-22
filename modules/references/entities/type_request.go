package entities

// --- CREATE 
type CreateTypeRequest struct {
	TypeName  string `json:"type_name" form:"type_name"`
	Code        string `json:"code" form:"code"`
}

func (req CreateTypeRequest) Validate() error {
	return nil
}

// --- UPDATE 
type UpdateTypeRequest struct {
	TypeName  string `json:"type_name" form:"type_name"`
	Code        string `json:"code" form:"code"`
}

func (req UpdateTypeRequest) Validate() error {
	return nil
}


// search
type SearchTypeRequest struct {
	SearchParam string `json:"search_param" form:"search_param"`
}