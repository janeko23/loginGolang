package models

//TableExperienciasEnTematicas Nombre de la tabla que registra las tematicas trabajadas por las personas
const TableExperienciasEnTematicas string = "experiencias_tematicas"

//TableInteresesEnTematicas Nombre de la tabla que registra los intereses de las persoas en tematicas
const TableInteresesEnTematicas string = "intereses_tematicas"

//TableCertificacionesEnTematicas Nombre de la tabla que registra las certificaciones en tematicas de las personas
const TableCertificacionesEnTematicas string = "certificaciones_tematicas"

//TablePendientesTematicas Nombre de la tabla que registra las tematicas trabajadas por las personas pendientes de ingresar al registro
const TablePendientesTematicas string = "pendientes_tematicas"


// ExpEnTematica modela experiencia de una determinada persona en una determinada tematica
type ExpEnTematica struct {
	IDPersona int `gorm:"column:IdPersona"`
	IDTematica int `gorm:"column:IdTematica"`
}

// InteresEnTematica modela experiencia de una determinada persona en una determinada tematica
type InteresEnTematica struct {
	IDPersona int `gorm:"column:IdPersona"`
	IDTematica int `gorm:"column:IdTematica"`
}

// CertificacionEnTematica modela experiencia de una determinada persona en una determinada tematica
type CertificacionEnTematica struct {
	IDPersona int `gorm:"column:IdPersona"`
	IDTematica int `gorm:"column:IdTematica"`
}