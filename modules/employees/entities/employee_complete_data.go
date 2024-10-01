package entities

type EmployeeCompleteData struct {
	EmployeeId          int64            `json:"employee_id"`
	UniqueId            string           `json:"unique_id"`
	PersonalInfo        PersonalBasic     `json:"personal_info,omitempty"`
	ProfessionalInfo    ProfessionalBasic  `json:"professional_info,omitempty"`
	Documents           []DocumentBasic    `json:"documents,omitempty"`
	Phones              []PhoneBasic       `json:"phones,omitempty"`
	Emails              []EmailBasic       `json:"emails,omitempty"`
	Addresses           []AddressBasic     `json:"addresses,omitempty"`
	UserAccount         AccountBasic       `json:"user_account,omitempty"`
}

type PersonalBasic struct {
	FirstName                string `json:"first_name"`
	LastName                 string `json:"last_name"`
	IdentificationNumber     string `json:"identification_number"`
	Gender                   string `json:"gender"`
	DateOfBirth              string `json:"date_of_birth"`
	CreatedAt                string `json:"created_at"`
	UpdatedAt                string `json:"updated_at"`
	IdentificationTypeName   string `json:"identification_type_name"`
	CountryName              string `json:"country_name"`
	MaritalStatusName        string `json:"marital_status_name"`
}

type ProfessionalBasic struct {
	DepartmentName           string `json:"department_name"`
	EmploymentStatusName     string `json:"employment_status_name"`
	JobTitleName             string `json:"job_title_name"`
}

type DocumentBasic struct {
	DocumentName             string `json:"document_name"`
	DocumentNumber           string `json:"document_number"`
	ExpirationDate           string `json:"expiration_date"`
	FileName                 string `json:"file_name"`
	Status                   string `json:"status"`
	DocumentTypeName         string `json:"document_type_name"`
}

type PhoneBasic struct {
	PhoneNumber              string `json:"phone_number"`
	ContactTypeName          string `json:"contact_type_name"`
}

type EmailBasic struct {
	EmailAddress             string `json:"email_address"`
	ContactTypeName          string `json:"contact_type_name"`
}

type AddressBasic struct {
	State                   string `json:"state"`
	City                    string `json:"city"`
	Neighborhood            string `json:"neighborhood"`
	Street                  string `json:"street"`
	HouseNumber             string `json:"house_number"`
	PostalCode              string `json:"postal_code"`
	CountryCode             string `json:"country_code"`
	AdditionalDetails       string `json:"additional_details"`
	IsCurrent               bool   `json:"is_current"`
}

type AccountBasic struct {
	UserName                string `json:"user_name"`
	Email                   string `json:"email"`
}
