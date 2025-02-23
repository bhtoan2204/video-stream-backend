package model

type ESPermission struct {
	ESAbstractModel
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
}
