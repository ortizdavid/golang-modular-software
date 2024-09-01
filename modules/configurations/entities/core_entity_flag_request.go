package entities

import "fmt"

type ManageCoreEntityFlagRequest struct {
	FlagId int    `json:"flag_id" form:"flag_id"`
	Status string `json:"statuses" form:"status"`
}

func (req ManageCoreEntityFlagRequest) Validate() error {
	if req.Status != "Enabled" && req.Status != "Disabled" {
		return fmt.Errorf("invalid status: %s", req.Status)
	}
	return nil
}
