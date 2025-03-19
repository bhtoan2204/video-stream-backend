package entities

type ActivityLog struct {
	AbstractModel
	UserID      string `json:"user_id"`
	Action      string `json:"action"`
	Description string `json:"description,omitempty"`
	IPAddress   string `json:"ip_address,omitempty"`
	UserAgent   string `json:"user_agent,omitempty"`
	DeviceToken string `json:"device_token"`
}
