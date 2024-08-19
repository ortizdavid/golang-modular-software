package entities

// ---- Create
type CreateRoomRequest struct {
	CompanyId      int `json:"company_id" form:"company_id"`
	BranchId    int `json:"branch_id" form:"branch_id"`
	RoomName  string `json:"room_name" form:"room_name"`
	Number    string `json:"number" form:"number"`
	Capacity  int `json:"capacity" form:"capacity"`
}

func (req CreateRoomRequest) Validate() error  {
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
	return nil
}


// --- Search
type SearchRoomRequest struct {
	SearchParam		string `json:"search_param" form:"search_param"`
}
