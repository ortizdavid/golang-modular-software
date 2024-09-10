package entities


// CoreEntityFlagStatus: uses the below structs
type CoreEntityFlagStatus struct {
	AuthenticationFlags AuthenticationFlags `json:"authentication"`
	ConfigurationFlags  ConfigurationFlags  `json:"configurations"`
	ReferenceFlags      ReferenceFlags      `json:"references"`
	CompanyFlags        CompanyFlags        `json:"company"`
	EmployeeFlags       EmployeeFlags       `json:"employees"`
	ReportFlags         ReportFlags         `json:"reports"`
}


// Structs for each module
type AuthenticationFlags struct {
	Users           string `json:"users"` // enabled or disabled
	ActiveUsers     string `json:"active_users"`
	InactiveUsers   string `json:"inactive_users"`
	OnlineUsers     string `json:"online_users"`
	OfflineUsers    string `json:"offline_users"`
	Roles           string `json:"roles"`
	Permissions     string `json:"permissions"`
	LoginActivity   string `json:"login_activity"`
}

type ConfigurationFlags struct {
	BasicConfigurations string `json:"basic_configurations"`
	CompanyConfigurations string `json:"company_configurations"`
	EmailConfigurations string `json:"email_configurations"`
	Modules             string `json:"modules"`
	CoreEntities        string `json:"core_entities"`
	ModuleFlags         string `json:"module_flags"`
	CoreEntityFlags     string `json:"core_entity_flags"`
}

type ReferenceFlags struct {
	Countries           string `json:"countries"`
	Currencies          string `json:"currencies"`
	IdentificationTypes string `json:"identification_types"`
	ContactTypes        string `json:"contact_types"`
	MaritalStatuses     string `json:"marital_statuses"`
	TaskStatuses        string `json:"task_statuses"`
	ApprovalStatuses    string `json:"approval_statuses"`
	DocumentStatuses    string `json:"document_statuses"`
	WorkflowStatuses    string `json:"workflow_statuses"`
	EvaluationStatuses  string `json:"evaluation_statuses"`
	UserStatuses        string `json:"user_statuses"`
	EmploymentStatuses  string `json:"employment_statuses"`
}

type CompanyFlags struct {
	CompanyInfo     string `json:"company_info"`
	Branches        string `json:"branches"`
	Offices         string `json:"offices"`
	Departments     string `json:"departments"`
	Rooms           string `json:"rooms"`
	Projects        string `json:"projects"`
	Policies        string `json:"policies"`
}

type EmployeeFlags struct {
	Employees   string `json:"employees"`
	JobTitles   string `json:"job_titles"`
	DocumentTypes	string `json:"document_types"`
}

type ReportFlags struct {
	UserReports         string `json:"user_reports"`
	ConfigurationReports string `json:"configuration_reports"`
	CompanyReports      string `json:"company_reports"`
	EmployeeReports     string `json:"employee_reports"`
	ReferenceReports    string `json:"reference_reports"`
}
