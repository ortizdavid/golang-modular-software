package entities

// -- Create
type CreateCurrencyRequest struct {
	CurrencyName string    `json:"currency_name" form:"currency_name"`
	Code         string    `json:"code" form:"code"`
}

func (req CreateCurrencyRequest) Validate() error {
	return nil
}

// -- Update
type UpdateCurrencyRequest struct {
	CurrencyName string    `json:"currency_name" form:"currency_name"`
	Code         string    `json:"code" form:"code"`
	Symbol       string  `json:"symbol" form:"symbol"`
}

func (req UpdateCurrencyRequest) Validate() error {
	return nil
}


// search
type SearchCurrencyRequest struct {
	SearchParam string `json:"search_param" form:"search_param"`
}