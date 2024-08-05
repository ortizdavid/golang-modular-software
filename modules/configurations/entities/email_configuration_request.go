package entities

type UpdateEmailConfigurationRequest struct {
	SMTPServer     string `json:"smtp_server" form:"smtp_server"`
	SMTPPort       string `json:"smtp_port" form:"smtp_port"`
	SenderEmail    string `json:"sender_email" form:"sender_email"`
	SenderPassword string `json:"sender_password" form:"sender_password"`
}

func (req UpdateEmailConfigurationRequest) Validate() error {
	return nil
}
