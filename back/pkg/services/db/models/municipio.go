package models

// TableMunicipios Tabla de las municipios
const TableMunicipios string = "municipios"

// Municipio struct
type Municipio struct {
	ID          string `gorm:"column:Id"`
	IDProvincia string `gorm:"column:munpro_id"`
	Nombre      string `gorm:"column:Nombre"`
}
