package models

// TableNivelesEducativos Tabla de los niveles educativos
const TableNivelesEducativos string = "niveles_educativos"

// NivelEducativo struct
type NivelEducativo struct {
	ID int 	  `gorm:"column:Id"`
	Nombre string 	  `gorm:"column:Nombre"`
}