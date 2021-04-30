package models

import (
	
	"gopkg.in/guregu/null.v3"
	
)

// TablePersonas Nombre de la tabla de personas
const TablePersonas  string = "personas"

//TablePendientes Nombre de la tabla de pendientes
const TablePendientes string = "pendientes"

//Persona model struct
type Persona struct {
	IDFuente			null.Int			`gorm:"column:IdFuente"`
	IDOrgaPertenencia	null.Int 	`gorm:"column:IdOrgaPertenencia"`
	Referente			null.String `gorm:"column:Referente"`
	CompletedTime 		null.Time 	`gorm:"column:CompletedTime"`
	ID 					int    		`gorm:"column:Id"`
	Nombres          	string 		`gorm:"column:Nombres"`
	Apellido        	string 		`gorm:"column:Apellido"`
	IDTipoDocumento 	int    		`gorm:"column:IdTipoDoc"`
	NroDocumento    	null.String `gorm:"column:NroDoc"`
	IDPaisNacimiento 	null.Int   	`gorm:"column:IdPaisDeNacimiento"`
	IDNacionalidad		null.Int	`gorm:"column:IdNacionalidad"`
	IDIdentidadGenero 	int			`gorm:"column:IdIdentidadDeGenero"`
	FechaNacimiento 	null.Time 	`gorm:"column:FechaDeNacimiento"`
	Edad				null.Int	`gorm:"column:Edad"`
	IDNivelEducativo 	null.Int	`gorm:"column:IdNivelEducativo"`
	NivelEducativoIngresado null.String `gorm:"column:NivelEducativoIngresado"`
	IDCompletoEstudios 	null.Int	`gorm:"column:IdCompletoEstudios"`
	Calle				null.String `gorm:"column:Calle"`
	Numero				null.String `gorm:"column:Numero"`
	Torre				null.String `gorm:"column:Torre"`
	Piso				null.String	`gorm:"column:Piso"`
	Depto				null.String `gorm:"column:Depto"`
	LocalidadIngresada  null.String `gorm:"column:LocalidadIngresada"`
	IDLocalidad 		null.String `gorm:"column:IdLocalidad"`
	PartidoIngresado	null.String `gorm:"column:PartidoIngresado"`
	IDPartido			null.String	`gorm:"column:IdPartido"`
	ProvinciaIngresada	null.String `gorm:"column:ProvinciaIngresada"`
	IDProvincia 		null.String `gorm:"column:IdProvincia"`

	/** -- Campos de base 2 -- **/
	//DatosAdicionalesDomicilio null.String `gorm:"column:"DatosAdicionalesDomicilio"`
	Telefono			null.String `gorm:"column:Telefono"`
	Email				null.String `gorm:"column:Email"`
	FuePromotorxMinisterio null.Bool `gorm:"column:FuePromotorxMinisterio"`
	Ministerio			null.String	`gorm:"column:Ministerio"`
	OtrasTematicasFormo	null.String	`gorm:"column:OtrasTematicasFormo"`
	OtrasTematicasCertifico null.String `gorm:"column:OtrasTematicasCertifico"`
	OtrasTematicasInteres	null.String `gorm:"column:OtrasTematicasInteres"`
	ObtuvoRemuneracion	null.Bool 	`gorm:"column:ObtuvoRemuneracion"`
	ObtuvoCertificacion	null.Bool	`gorm:"column:ObtuvoCertificacion"`
	DeOrganismoDelEstadoYSociedadCivil null.String	`gorm:"column:DeOrganismoDelEstadoYSociedadCivil"`
	DeOtroOrganismo		null.String	 `gorm:"column:DeOtroOrganismo"`
	IDOrgaFormacion		null.Int	`gorm:"column:IdOrgaFormacion"`

	Observaciones 		null.String	`gorm:"column:Observaciones"`



}