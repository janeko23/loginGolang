package models

//TableCertificacionesOrganismos Nombre de la tabla que registra los certificaciones de organismos
const TableCertificacionesOrganismos string = "certificaciones_organismos"


// CertificacionDeOrganismo modela certificacion de una determinada persona en una determinada organismo
type CertificacionDeOrganismo struct {
	IDPersona int `gorm:"column:IdPersona"`
	IDOrganismo int `gorm:"column:IdOrganismo"`
}