package models

//TableNivelesEducativosConcluidos nombre de la tabla
const TableNivelesEducativosConcluidos = "estados_nivel_educativo"

// NivelEducativoConcluido struct
type NivelEducativoConcluido struct {
	ID int 	  `gorm:"column:Id"`
	Nombre string 	  `gorm:"column:Nombre"`
}
