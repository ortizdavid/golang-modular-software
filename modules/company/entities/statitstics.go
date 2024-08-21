package entities

type Statistics struct {
	Branches    int64  `json:"branches"`
	Offices     int64  `json:"offices"`
	Departments int64  `json:"departments"`
	Rooms       int64  `json:"rooms"`
	Projects    int64  `json:"projects"`
	Policies    int64  `json:"policies"`
}