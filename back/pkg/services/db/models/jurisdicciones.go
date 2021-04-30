package models

// Jurisdicciones struct
type Jurisdicciones struct {
	ID     int    `gorm:"column:Id"`
	Nombre string `gorm:"column:Nombre"`
}
