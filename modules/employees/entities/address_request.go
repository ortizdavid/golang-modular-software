package entities

type CreateAddressRequest struct {
	EmployeeId			int64 `json:"employee_id" form:"employee_id"`
	State				string `json:"state" form:"state"`
	City				string `json:"city" form:"city"`
	Neighborhood		string `json:"neighborhood" form:"neighborhood"`
	Street				string `json:"street" form:"street"`
	HouseNumber			string `json:"house_number" form:"house_number"`
	PortalCode			string `json:"portal_code" form:"portal_code"`
	CountryCode			string `json:"country_code" form:"country_code"`
	AditionalDetails	string `json:"aditional_details" form:"aditional_details"`
}