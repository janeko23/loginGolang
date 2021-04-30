package models

//TableOrganismos Nombre de la tabla de organismos
const TableOrganismos string = "organismos"


//Organismo struct modela un organismo
type Organismo struct {
	ID int `gorm:"column:Id"`
	Nombre string `gorm:"column:Nombre size:150" validate:"max=150"`
}