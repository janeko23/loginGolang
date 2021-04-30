package models

// TableTiposOrgas Tabla de las municipios
const TableTiposOrgas string = "tipos_organizaciones"

// TipoOrganizacion struct
type TipoOrganizacion struct {
	ID          int `gorm:"column:Id"`
	Nombre      string `gorm:"column:Nombre"`
}