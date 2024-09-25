package entities

import (
	"github.com/go-playground/validator/v10"
	"github.com/ortizdavid/golang-modular-software/common/helpers"
)

// ---- Create
type CreateRoomRequest struct {
	CompanyId      int `json:"company_id" form:"company_id"`
	BranchId    int `json:"branch_id" form:"branch_id"`
	RoomName  string `json:"room_name" form:"room_name"`
	Number    string `json:"number" form:"number"`
	Capacity  int `json:"capacity" form:"capacity"`
}

func (req CreateRoomRequest) Validate() error  {
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

// ---- Update
type UpdateRoomRequest struct {
	CompanyId      int `json:"company_id" form:"company_id"`
	BranchId    int `json:"branch_id" form:"branch_id"`
	RoomName  string `json:"room_name" form:"room_name"`
	Number    string `json:"number" form:"number"`
	Capacity  int `json:"capacity" form:"capacity"`
}

func (req UpdateRoomRequest) Validate() error  {
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
type SearchRoomRequest struct {
	SearchParam		string `json:"search_param" form:"search_param"`
}
