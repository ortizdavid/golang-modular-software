package entities

// -- Create
type CreateCountryRequest struct {
	CountryName 	string `json:"country_name" form:"country_name"`
	IsoCode 		string `json:"iso_code" form:"iso_code"`
	DialingCode 	string `json:"dialing_code" form:"dialing_code"`
}

func (req CreateCountryRequest) Validate() error {
	return nil
}

// -- Update
type UpdateCountryRequest struct {
	CountryName 	string `json:"country_name" form:"country_name"`
	IsoCode 		string `json:"iso_code" form:"iso_code"`
	DialingCode 	string `json:"dialing_code" form:"dialing_code"`
}

func (req UpdateCountryRequest) Validate() error {
	return nil
}


// search
type SearchCountryRequest struct {
	SearchParam string `json:"search_param" form:"search_param"`
}