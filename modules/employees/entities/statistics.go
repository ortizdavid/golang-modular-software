package entities

type Statistics struct {
	Employees int64 `json:"employees"`
	JobTitles int64 `json:"job_titles"`
	DocumentTypes	int64 `json:"document_types"`
}