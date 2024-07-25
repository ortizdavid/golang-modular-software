package helpers

import "errors"

type PaginationParam struct {
	CurrentPage int `json:"current_page"`
	Limit       int `json:"limit"`
}

func (p *PaginationParam) Validate() error {
	if p.CurrentPage <= 0 {
		return errors.New("current_page must be greater than 0")
	}
	if p.Limit <= 0 {
		return errors.New("limit must be greater than 0")
	}
	return nil
}

func DefaultPaginationParams() PaginationParam {
	return PaginationParam{
		CurrentPage: 1,
		Limit:       10,
	}
}