package entities

// -- Create
type CreateJobTitleRequest struct {
	TitleName   string `json:"title_name" form:"title_name"`
	Description string `json:"description" form:"description"`
}

func (req CreateJobTitleRequest) Validate() error {
	return nil
}

// -- Update
type UpdateJobTitleRequest struct {
	TitleName   string `json:"title_name" form:"title_name"`
	Description string `json:"description" form:"description"`
}

func (req UpdateJobTitleRequest) Validate() error {
	return nil
}

type SearchJobTitleRequest struct {
	SearchParam string `json:"search_param" form:"search_param"`
}