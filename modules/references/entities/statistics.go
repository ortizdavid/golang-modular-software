package entities

type Statistics struct {
	Countries           int64 `json:"countries"`
	Currencies         int64 `json:"currencies"`
	IdentificationTypes int64 `json:"identification_types"`
	ContactTypes        int64 `json:"contact_types"`
	MaritalStatuses     int64 `json:"marital_statuses"`
	TaskStatuses        int64 `json:"task_statuses"`
	ApprovalStatuses    int64 `json:"approval_statuses"`
	DocumentStatuses    int64 `json:"document_statuses"`
	WorkflowStatuses    int64 `json:"workflow_statuses"`
	EvaluationStatuses    int64 `json:"evaluation_statuses"`
	UserStatuses    	int64 `json:"user_statuses"`
}
