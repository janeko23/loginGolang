package models

// TableTiposDeDocumentos Tabla de los tipos de documentos
const TableTiposDeDocumentos string = "tipos_documentos"

// TipoDocumento struct
type TipoDocumento struct {
	ID		int 	  `gorm:"column:Id"`
	Nombre  string 	  `gorm:"column:Nombre"`
}