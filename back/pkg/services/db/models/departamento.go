package models

// TableDepartamentos Tabla de los departamentos
const TableDepartamentos string = "departamentos"

// Departamento struct
type Departamento struct {
	ID     string    `gorm:"column:Id"`
	Nombre string `gorm:"column:Nombre"`
	IDProvincia string `gorm:"column:IdProvincia"`
}