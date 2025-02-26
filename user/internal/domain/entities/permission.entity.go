package entities

type Permission struct {
	AbstractModel
	Name        string `json:"name"` // Unique constraint to prevent duplicates
	Description string `json:"description,omitempty"`
}

type IPermission interface {
}
