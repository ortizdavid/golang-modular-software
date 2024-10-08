package entities

type EmployeeAccountData struct {
	UserId				int16 `json:"user_id"`
	UserName			string `json:"user_name"`
	Email				string `json:"email"`
	UniqueId			string `json:"unique_id"`
	CreatedAt			string `json:"created_at"`
	UpdatedAt			string `json:"updated_at"`
	EmployeeId			int64 `json:"employee_id"`
	EmployeeUniqueId	string `json:"employee_unique_id"`
	FirstName			string `json:"first_name"`
	LastName			string `json:"last_name"`
	IdentificationNumber	string `json:"identification_number"`
}
