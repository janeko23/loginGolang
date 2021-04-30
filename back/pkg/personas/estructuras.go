package personas

import(
	"gopkg.in/guregu/null.v3"
)

// Persona struct
type Persona struct{
	Metadata		Metadata		`json:"formMetadata"`
	DatosPersonales DatosPersonales `json:"formDatosPersonales"`
	Contacto 		Contacto 		`json:"formContacto"`
	Educacion 		Educacion 		`json:"formEducacion"`
	Experiencia 	Experiencia 	`json:"formTrabajo"`
}

//Metadata struct
type Metadata struct {
	IDFuente				int				`json:"idFuente"`
	CompletedTime			null.Time	 	`json:"completedTime"`
	OrganizacionCarga		null.String		`json:"organizacionCarga"`
	ReferenteCarga			null.String		`json:"referente"`
}

//DatosPersonales struct
type DatosPersonales struct{
	Apellido        		string	  		`json:"apellido"`
	Nombres 				string	  		`json:"nombre"`
	IDTipoDocumento			int				`json:"idTipoDocumento"`
	NroDocumento  			null.String	  	`json:"nroDocumento"`
	FechaNacimiento 		null.Time 			`json:"fechaNacimiento"`
	IDIdentidadGenero		int				`json:"idIdentidadGenero"`
	Edad					null.Int				`json:"edad"`
}

//Contacto struct
type Contacto struct{
	Calle			null.String	`json:"calle"`
	Numero			null.String	`json:"numero"`
	Depto			null.String	`json:"depto"`
	Piso			null.String	`json:"piso"`
	Torre			null.String `json:"torre"`
	IDLocalidad		null.String `json:"idLocalidad"`	
	IDPartido		null.String `json:"idPartido"`
	IDProvincia 	null.String `json:"idProvincia"`
	IDPaisNacimiento null.Int   `json:"idPaisNacimiento"`
	PaisNacimiento 	null.String `json:"paisNacimiento"`
	IDNacionalidad	null.Int	`json:"idNacionalidad"`
	Telefono		null.String `json:"telefono"`
	Email			null.String `json:"email"`

}


//Educacion struct
type Educacion struct{
	IDNivelEducativo		null.Int	`json:"idNivelEducativo"`
	IDCompletoNivelEducativo null.Int `json:"idCompletoNivelEducativo"`
}

//Experiencia struct
type Experiencia struct{
	TematicasTrabajo			[]int		`json:"tematicasTrabajo"`
	OtrasTematicasTrabajo		null.String `json:"otrasTematicasTrabajo"`
	TematicasInteres			[]int		`json:"tematicasInteres"`
	OtrasTematicasInteres		null.String `json:"otrasTematicasInteres"`
	TematicasCertifico			[]int		`json:"tematicasCertifico"`
	OtrasTematicasCertifico		null.String	`json:"otrasTematicasCertifico"`

	OrganizacionPertenencia		null.String `json:"organizacionPertenencia"`
	FuePromotorxMinisterio		null.Bool	`json:"fuePromotorxMinisterio"`
	Ministerio					null.String `json:"ministerio"`
	ObtuvoRemuneracion			null.Bool	`json:"obtuvoRemuneracion"`
	ObtuvoCertificacion			null.Bool	`json:"obtuvoCertificacion"`
	DeOrganismoDelEstadoYSociedadCivil null.String `json:"DeOrganismoDelEstadoYSociedadCivil"`
	DeOtroOrganismo				null.String `json:"deOtroOrganismo"`
	IDsOrganismosCertificacion 	[]int 		`json:"idsOrganismosCertificacion"`
	IDOrganizacionFormacion		null.Int	`json:"idOrganizacionFormacion"`
}


/** ------------------------------------------------------------------------------------------- **/

type PersonaDenormalizada struct {
	Metadata		MetadataD		 `json:"Metadata"`
	DatosPersonales DatosPersonalesD `json:"DatosPersonales"`
	Contacto 		ContactoD 		 `json:"Contacto"`
	Educacion 		EducacionD 		 `json:"Educacion"`
	Experiencia 	ExperienciaD 	 `json:"Experiencia"`
}

//MetadataD struct
type MetadataD struct {
	Fuente					string			`json:"fuente" validate:"max=50"`
	CompletedTime			null.Time	 	`json:"completedTime"` //TODO: Agregar validacion fecha anterior a hoy?
	OrganizacionCarga		null.String		`json:"organizacionCarga" validate:"max=150"`
	ReferenteCarga			null.String		`json:"referente" validate:"max=100"`
	Observaciones			null.String		`json:"observaciones" validate:"max=100"`
}

//DatosPersonalesD struct
type DatosPersonalesD struct{
	Apellido        		string	  		`json:"apellido" validate:"max=100"`
	Nombres 				string	  		`json:"nombre" validate:"max=100"`
	TipoDocumento 			string			`json:"tipoDocumento"`
	NroDocumento  			null.String	  	`json:"nroDocumento" validate:"max=20"`
	FechaNacimiento 		null.Time 		`json:"fechaNacimiento"`
	IdentidadGenero 		string	  		`json:"identidadGenero" validate:"max=20"`
	Edad					null.Int		`json:"edad"`
}

//ContactoD struct
type ContactoD struct{
	Calle			null.String	`json:"calle" validate:"max=100"`
	Numero			null.String	`json:"numero" validate:"max=100"`
	Depto			null.String	`json:"depto" validate:"max=100"`
	Piso			null.String	`json:"piso" validate:"max=100"`
	Torre			null.String `json:"torre" validate:"max=100"`
	Localidad 		null.String `json:"localidad" validate:"max=100"`
	Partido			null.String `json:"partido" validate:"max=100"`
	Provincia 		null.String `json:"provincia" validate:"max=50"`
	PaisNacimiento 	null.String `json:"paisNacimiento" validate:"max=50"`
	Nacionalidad	null.String	`json:"nacionalidad" validate:"max=50"`
	Telefono		null.String `json:"telefono" validate:"max=20"`
	Email			null.String `json:"email" validate:"max=100"`

}


//EducacionD struct
type EducacionD struct{
	NivelEducativo 			null.String `json:"nivelEducativo" validate:"max=100"`
	CompletoNivelEducativo 	null.String `json:"completoNivelEducativo" validate:"max=100"`
}


//ExperienciaD struct
type ExperienciaD struct{
	TematicasTrabajo			[]int		`json:"tematicasTrabajo"`
	OtrasTematicasTrabajo		null.String `json:"otrasTematicasTrabajo" validate:"max=150"`
	TematicasInteres			[]int		`json:"tematicasInteres"`
	OtrasTematicasInteres		null.String `json:"otrasTematicasInteres" validate:"max=150"`
	TematicasCertifico			[]int		`json:"tematicasCertifico"`
	OtrasTematicasCertifico		null.String	`json:"otrasTematicasCertifico" validate:"max=150"`

	OrganizacionPertenencia		null.String `json:"organizacionPertenencia" validate:"max=150"`
	FuePromotorxMinisterio		null.Bool	`json:"fuePromotorxMinisterio"`
	Ministerio					null.String `json:"Ministerio" validate:"max=100"`
	ObtuvoRemuneracion			null.Bool	`json:"obtuvoRemuneracion"`
	ObtuvoCertificacion			null.Bool	`json:"obtuvoCertificacion"`
	DeOrganismoDelEstadoYSociedadCivil null.String `json:"DeOrganismoDelEstadoYSociedadCivil" validate:"max=100"`
	DeOtroOrganismo				null.String `json:"deOtroOrganismo" validate:"max=150"`
	OrganismosCertificacion		[]string		`json:"OrganismosCertificacion"`
	OrganizacionFormacion		null.String `json:"organizacionFormacion" validate:"max=150"`
}