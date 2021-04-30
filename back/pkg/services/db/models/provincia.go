package models

// TableProvincias Tabla de las provincias
const TableProvincias string = "provincias"

// Provincia struct
type Provincia struct {
	ID     string `gorm:"column:Id"`
	Nombre string `gorm:"column:Nombre"`
}