package models

// TableTematicas Tabla de los ministerios de la nacion
const TableTematicas string = "tematicas"

// Tematica struct
type Tematica struct {
	ID     int    `gorm:"column:Id"`
	Nombre string `gorm:"column:Nombre"`
}