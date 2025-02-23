package model

type ESRole struct {
	ESAbstractModel
	Name        string         `json:"name"`
	Permissions []ESPermission `json:"permissions,omitempty"`
}
