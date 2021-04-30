package models

//NormalizacionGenerica struct modela una normalizacion simple que solo almacena nombres
type NormalizacionGenerica struct {
	ID	 	int 	`gorm:"column:Id"`
	Nombre  string	`gorm:"column:Nombre"`
}