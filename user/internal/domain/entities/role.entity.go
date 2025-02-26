package entities

type Role struct {
	AbstractModel
	Name        string        `json:"name"`
	Permissions []*Permission `json:"permissions"`
}
