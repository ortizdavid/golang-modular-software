package entities

// -- Create 
type CreateDocumentRequest struct {
	EmployeeId		int64 `json:"employee_id" form:"employee_id"`
	DocumentTypeId	int	`json:"document_type_id" form:"document_type_id"`
	DocumentName   	string `json:"document_name" form:"document_name"`
	DocumentNumber 	string `json:"document_number" form:"document_number"`
	ExpirationDate	string `json:"expiration_date" form:"expiration_date"`
	Status			string `json:"status" form:"status"`
}

func (req CreateDocumentRequest) Validate() error {
	return nil
}

// -- Updaye 
type UpdateDocumentRequest struct {
	EmployeeId		int64 `json:"employee_id" form:"employee_id"`
	DocumentTypeId	int	`json:"document_type_id" form:"document_type_id"`
	DocumentName   	string `json:"document_name" form:"document_name"`
	DocumentNumber 	string `json:"document_number" form:"document_number"`
	ExpirationDate	string `json:"expiration_date" form:"expiration_date"`
	Status			string `json:"status" form:"status"`
}

func (req UpdateDocumentRequest) Validate() error {
	return nil
}

type SearchDocumentRequest struct {
	SearchParam string `json:"search_param" form:"search_param"`
}