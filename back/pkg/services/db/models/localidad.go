package models

import (
	"gopkg.in/guregu/null.v3"
)

// TableLocalidades Tabla de las localidades
const TableLocalidades string = "localidades"

// Localidad struct
type Localidad struct {
	ID     		   string 		  `gorm:"column:Id"`
	IDProvincia    string 		  `gorm:"column:IdProvincia"`
	IDDepartamento null.String 	  `gorm:"column:IdDepartamento"`
	IDMunicipio    null.String    `gorm:"column:IdMunicipio"`
	Nombre 		   string 		  `gorm:"column:Nombre"`
}