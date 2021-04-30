package models

// TableOrganizaciones Tabla de las municipios
const TableOrganizaciones string = "organizaciones"

// Organizacion struct
type Organizacion struct {
	ID          int `gorm:"column:Id"`
	//IDTipo 		int `gorm:"column:IdTipo"`
	Nombre      string `gorm:"column:Nombre" validate:"max=150"`
	//IDReferente int `gorm:"column:IdReferente`
}