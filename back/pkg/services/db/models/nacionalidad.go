package models

// TableNacionalidades Tabla de los paises
const TableNacionalidades string = "nacionalidades"

// Nacionalidad struct
type Nacionalidad struct {
	ID     int    `gorm:"column:Id"`
	Nombre string `gorm:"column:Nombre"`
}