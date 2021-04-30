package models

// TableMinisterios Tabla de los ministerios de la nacion
const TableMinisterios string = "ministerios"

// Ministerio struct
type Ministerio struct {
	ID     int    `gorm:"column:Id"`
	Nombre string `gorm:"column:Nombre"`
}
