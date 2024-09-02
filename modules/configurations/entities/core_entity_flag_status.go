package entities

type CoreEntityFlagStatus struct {
	
	// Module: Authentication
	Users          		string `json:"users"`
    ActiveUsers    		string `json:"active_users"`
    InactiveUsers  		string `json:"inactive_users"`
    OnlineUsers    		string `json:"online_users"`
    OfflineUsers   		string `json:"offline_users"`
    Roles          		string `json:"roles"`
    Permissions    		string `json:"permissions"`
    LoginActivity  		string `json:"login_activity"`

	// Module: Company
	Branches    		string  `json:"branches"`
	Offices     		string  `json:"offices"`
	Departments 		string  `json:"departments"`
	Rooms       		string  `json:"rooms"`
	Projects    		string  `json:"projects"`
	Policies    		string  `json:"policies"`

	// Module: References
	Countries           string `json:"countries"`
	Currencies         	string `json:"currencies"`
	IdentificationTypes string `json:"identification_types"`
	ContactTypes        string `json:"contact_types"`
	MaritalStatuses     string `json:"marital_statuses"`
	TaskStatuses        string `json:"task_statuses"`
	ApprovalStatuses    string `json:"approval_statuses"`
	DocumentStatuses    string `json:"document_statuses"`
	WorkflowStatuses    string `json:"workflow_statuses"`
	EvaluationStatuses  string `json:"evaluation_statuses"`
	UserStatuses    	string `json:"user_statuses"`

	// Module: Employees
	Employees			string `json:"employees"`
	JobTitles			string `json:"job_title"`
}

