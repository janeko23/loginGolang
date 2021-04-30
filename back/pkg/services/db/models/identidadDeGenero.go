package models

// TableIdentidadesDeGenero Tabla de los niveles educativos
const TableIdentidadesDeGenero string = "identidades_de_genero"

// NivelEducativo struct
type identidadDeGenero struct {
	ID int 	  		`gorm:"column:Id"`
	Nombre string 	  `gorm:"column:Nombre"`
}