package entities

// --- CREATE 
type CreateStatusRequest struct {
	StatusName  string `json:"status_name" form:"status_name"`
	Code        string `json:"code" form:"code"`
	Weight      int `json:"weight" form:"weight"`
	LblColor    string `json:"lbl_color" form:"lbl_color"`
	BgColor     string `json:"bg_color" form:"bg_color"`
	Description string `json:"description" form:"description"`
}

func (req CreateStatusRequest) Validate() error {
	return nil
}

// --- UPDATE 
type UpdateStatusRequest struct {
	StatusName  string `json:"status_name" form:"status_name"`
	Code        string `json:"code" form:"code"`
	Weight      int `json:"weight" form:"weight"`
	LblColor    string `json:"lbl_color" form:"lbl_color"`
	BgColor     string `json:"bg_color" form:"bg_color"`
	Description string `json:"description" form:"description"`
}

func (req UpdateStatusRequest) Validate() error {
	return nil
}


// search
type SearchStatusRequest struct {
	SearchParam string `json:"search_param" form:"search_param"`
}