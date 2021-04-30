package models

// TableFuentes Tabla de las fuentes de informacion
const TableFuentes string = "fuentes"

const FuenteExcel string = "excel"
const FuenteSQL string = "sql"
const FuenteFormInterno string = "formInterno"
const FuenteFormExterno string = "formExterno"

// Fuente struct
type Fuente struct {
	ID     int    `gorm:"column:Id"`
	Nombre string `gorm:"column:Nombre"`
}
