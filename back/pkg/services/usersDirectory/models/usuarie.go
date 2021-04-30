package models

// Usuarie ldap struct
type Usuarie struct {

	UID   	 string `json:"uid"`
	Nombre   string `json:"nombre"`
	Apellido string `json:"apellido"`
	Email    string `json:"email"`
	Password string `json:"password"`
}