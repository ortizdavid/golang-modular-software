package entities

//- Create
type CreateAddressRequest struct {
	EmployeeId			int64 `json:"employee_id" form:"employee_id"`
	State				string `json:"state" form:"state"`
	City				string `json:"city" form:"city"`
	Neighborhood		string `json:"neighborhood" form:"neighborhood"`
	Street				string `json:"street" form:"street"`
	HouseNumber			string `json:"house_number" form:"house_number"`
	PostalCode			string `json:"portal_code" form:"postal_code"`
	CountryCode			string `json:"country_code" form:"country_code"`
	AditionalDetails	string `json:"aditional_details" form:"aditional_details"`
}

func (req CreateAddressRequest) Validate() error {
	return nil
}

//- Update
type UpdateAddressRequest struct {
	EmployeeId			int64 `json:"employee_id" form:"employee_id"`
	State				string `json:"state" form:"state"`
	City				string `json:"city" form:"city"`
	Neighborhood		string `json:"neighborhood" form:"neighborhood"`
	Street				string `json:"street" form:"street"`
	HouseNumber			string `json:"house_number" form:"house_number"`
	PostalCode			string `json:"portal_code" form:"postal_code"`
	CountryCode			string `json:"country_code" form:"country_code"`
	AditionalDetails	string `json:"aditional_details" form:"aditional_details"`
}

func (req UpdateAddressRequest) Validate() error {
	return nil
}