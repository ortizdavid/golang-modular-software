package entities

// -- Create
type CreateDocumentTypeRequest struct {
	TypeName   string `json:"type_name" form:"type_name"`
	Description string `json:"description" form:"description"`
}

func (req CreateDocumentTypeRequest) Validate() error {
	return nil
}

// -- Update
type UpdateDocumentTypeRequest struct {
	TypeName   string `json:"type_name" form:"type_name"`
	Description string `json:"description" form:"description"`
}

func (req UpdateDocumentTypeRequest) Validate() error {
	return nil
}

type SearchDocumentTypeRequest struct {
	SearchParam string `json:"search_param" form:"search_param"`
}