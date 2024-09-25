package entities

import (
	"github.com/go-playground/validator/v10"
	"github.com/ortizdavid/golang-modular-software/common/helpers"
)

// --- Create
type CreateProjectRequest struct {
	CompanyId   int `json:"company_id" form:"company_id" validate:"required"`
	ProjectName string  `json:"project_name" form:"project_name" validate:"required,max=100"`
	Description string  `json:"description" form:"description" validate:"max=255"`
	StartDate  	string `json:"start_date" form:"start_date"`
	EndDate     string `json:"end_date" form:"end_date"`
	Status      string  `json:"status" form:"status" validate:"max=50"`
}

func (req CreateProjectRequest) Validate() error  {
	validate := validator.New()
	err := validate.Struct(req)
	if err != nil {
		if errs, ok := err.(validator.ValidationErrors); ok {
			return helpers.ValidatorFormatErrors(errs)
		}
		return err
	}
	return nil
}

// --- Update
type UpdateProjectRequest struct {
	CompanyId   int `json:"company_id" form:"company_id" validate:"required"`
	ProjectName string  `json:"project_name" form:"project_name" validate:"required,max=100"`
	Description string  `json:"description" form:"description" validate:"max=255"`
	StartDate  	string `json:"start_date" form:"start_date"`
	EndDate     string `json:"end_date" form:"end_date"`
	Status      string  `json:"status" form:"status" validate:"max=50"`
}

func (req UpdateProjectRequest) Validate() error  {
	validate := validator.New()
	err := validate.Struct(req)
	if err != nil {
		if errs, ok := err.(validator.ValidationErrors); ok {
			return helpers.ValidatorFormatErrors(errs)
		}
		return err
	}
	return nil
}


// --- Search
type SearchProjectRequest struct {
	SearchParam		string `json:"search_param" form:"search_param"`
}

