package models

type User struct {
	ID        string `json:"id,omitempty"`
	Username  string `json:"username,omitempty"`
	Email     string `json:"email,omitempty"`
	FirstName string `json:"Firstname,omitempty"`
	LastName  string `json:"Lastname,omitempty"`
	Address   string `json:"Address,omitempty"`
}
