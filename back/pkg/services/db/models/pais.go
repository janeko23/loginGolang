package models

// TablePaises Tabla de los paises
const TablePaises string = "paises"

// Paises struct
type Paises struct {
	ID     int    `gorm:"column:Id"`
	Nombre string `gorm:"column:Nombre"`
}