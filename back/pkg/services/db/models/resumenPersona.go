package models

// ResumenPersona struct
type ResumenPersona struct {
	ID int `json:"id"`
	Nombres  string `json:"nombre"`
	Apellido string `json:"apellido"`
	NroDocumento string `json:"tematicas"`
	Provincia string `json:"provincia"`
	Organizacion string	`json:"orga"`
	Referente	string `json:"referente"`
  }