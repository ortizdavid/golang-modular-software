package entities

type Statistics struct {
	Users         int64 `json:"users"`
	ActiveUsers   int64 `json:"active_users"`
	InactiveUsers int64 `json:"inactive_users"`
	OnlineUsers   int64 `json:"online_users"`
	OfflineUsers  int64 `json:"offline_users"`
	Roles         int64 `json:"roles"`
	Permissions   int64 `json:"permissions"`
	LoginActivity int64 `json:"login_activity"`
}
