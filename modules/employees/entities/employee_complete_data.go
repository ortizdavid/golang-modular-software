package entities

type EmployeeCompleteData struct {
	EmployeeId       int64             `json:"employee_id"`
	UniqueId         string            `json:"unique_id"`
	CreatedAt        string            `json:"created_at"`
	UpdatedAt        string            `json:"updated_at"`
	PersonalInfo     PersonalBasic     `json:"personal_info"`
	ProfessionalInfo ProfessionalBasic `json:"professional_info"`
	Documents        []DocumentBasic   `json:"documents"`
	Phones           []PhoneBasic      `json:"phones"`
	Emails           []EmailBasic      `json:"emails"`
	Addresses        []AddressBasic    `json:"addresses"`
	UserAccount      AccountBasic      `json:"user_account"`
}

type PersonalBasic struct {
	FirstName              string `json:"first_name"`
	LastName               string `json:"last_name"`
	IdentificationNumber   string `json:"identification_number"`
	Gender                 string `json:"gender"`
	DateOfBirth            string `json:"date_of_birth"`
	IdentificationTypeName string `json:"identification_type_name"`
	CountryName            string `json:"country_name"`
	MaritalStatusName      string `json:"marital_status_name"`
}

type ProfessionalBasic struct {
	DepartmentName       string `json:"department_name,omitempty"`
	EmploymentStatusName string `json:"employment_status_name,omitempty"`
	JobTitleName         string `json:"job_title_name,omitempty"`
}

type DocumentBasic struct {
	DocumentName     string `json:"document_name,omitempty"`
	DocumentNumber   string `json:"document_number,omitempty"`
	ExpirationDate   string `json:"expiration_date,omitempty"`
	FileName         string `json:"file_name,omitempty"`
	UploadPath       string `json:"upload_path,omitempty"`
	Status           string `json:"status,omitempty"`
	DocumentTypeName string `json:"document_type,omitempty"`
}

type PhoneBasic struct {
	PhoneNumber     string `json:"phone_number,omitempty"`
	ContactTypeName string `json:"contact_type,omitempty"`
}

type EmailBasic struct {
	EmailAddress    string `json:"email_address,omitempty"`
	ContactTypeName string `json:"contact_type,omitempty"`
}

type AddressBasic struct {
	State             string `json:"state,omitempty"`
	City              string `json:"city,omitempty"`
	Neighborhood      string `json:"neighborhood,omitempty"`
	Street            string `json:"street,omitempty"`
	HouseNumber       string `json:"house_number,omitempty"`
	PostalCode        string `json:"postal_code,omitempty"`
	CountryCode       string `json:"country_code,omitempty"`
	AdditionalDetails string `json:"additional_details,omitempty"`
	IsCurrent         bool   `json:"is_current,omitempty"`
}

type AccountBasic struct {
	UserName  string `json:"user_name,omitempty"`
	UserEmail string `json:"user_email,omitempty"`
}

type EmployeeCompleteDataRaw struct {

	// Personal info
	EmployeeId             int64  `json:"employee_id"`
	UniqueId               string `json:"unique_id"`
	CreatedAt              string `json:"created_at"`
	UpdatedAt              string `json:"updated_at"`
	FirstName              string `json:"first_name"`
	LastName               string `json:"last_name"`
	IdentificationNumber   string `json:"identification_number"`
	Gender                 string `json:"gender"`
	DateOfBirth            string `json:"date_of_birth"`
	IdentificationTypeName string `json:"identification_type"`
	CountryName            string `json:"country_name"`
	MaritalStatusName      string `json:"marital_status"`

	// Professional Info
	DepartmentName       string `json:"department_name"`
	EmploymentStatusName string `json:"employment_status"`
	JobTitleName         string `json:"title_name"`

	// Documents
	DocumentName     string `json:"document_name"`
	DocumentNumber   string `json:"document_number"`
	ExpirationDate   string `json:"expiration_date"`
	FileName         string `json:"file_name"`
	UploadPath       string `json:"upload_path"`
	DocumentStatus   string `json:"document_status"`
	DocumentTypeName string `json:"document_type"`

	// Phones and Emails
	PhoneNumber          string `json:"phone_number"`
	PhoneContactTypeName string `json:"phone_contact_type"`
	EmailAddress         string `json:"email_address"`
	EmailContactTypeName string `json:"email_contact_type"`

	// Addresses
	State             string `json:"state"`
	City              string `json:"city"`
	Neighborhood      string `json:"neighborhood"`
	Street            string `json:"street"`
	HouseNumber       string `json:"house_number"`
	PostalCode        string `json:"postal_code"`
	CountryCode       string `json:"country_code"`
	AdditionalDetails string `json:"additional_details"`
	IsCurrent         bool   `json:"is_current"`

	// User Account
	UserName  string `json:"user_name"`
	UserEmail string `json:"user_email"`
}
