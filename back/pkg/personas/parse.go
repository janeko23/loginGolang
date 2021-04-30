package personas

import(
	"gopkg.in/guregu/null.v3"

	"igualdad.mingeneros.gob.ar/pkg/services/db/models"
//	"igualdad.mingeneros.gob.ar/pkg/utils"
	"igualdad.mingeneros.gob.ar/pkg/georef"
//	"errors"
)
func mapModelAPersona(personasModel *[]models.Persona) (*[]Persona, error) {
	var personas []Persona

	for _, personaModel := range *personasModel {

		var persona Persona

		err := persona.setResumenPersonaModel(personaModel)
		if err != nil {
			return nil, err
		}

		personas = append(personas, persona)
	}

	return &personas, nil
}

func mapExperienciasEnTematicas(idPersona int, experienciasEnTematicasIds *[]int) *[]models.ExpEnTematica {
	
	var personExperienciaEnTematica []models.ExpEnTematica

	for _, idTematica := range *experienciasEnTematicasIds {

		personExperienciaEnTematica = append(personExperienciaEnTematica, models.ExpEnTematica{IDPersona: idPersona, IDTematica: idTematica})
	}

	return &personExperienciaEnTematica
}


func mapCertificaciones(idPersona int, organismos *[]string) *[]models.CertificacionDeOrganismo {
	
	var certificaciones []models.CertificacionDeOrganismo

	for _, organismo := range *organismos {

		idOrganismo := int(getIDFromValue(organismo, models.TableOrganismos).ValueOrZero())

		certificaciones = append(certificaciones, models.CertificacionDeOrganismo{IDPersona: idPersona, IDOrganismo: idOrganismo})
	}

	return &certificaciones
}



func mapModelAPersonaCompleto(persona *Persona)  (*models.Persona, error) {
	
	var personaModel models.Persona

	
	return &personaModel, nil
}

func getPersonaModel(persona *Persona, from string)  ( *models.Persona, error) {
	var personaRes models.Persona
	// switch from {
	// case utils.FromImport:
	// 	personaExcel, err := getPersonaModelInputDenormalizado(persona)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	personaRes = *personaExcel
	// case utils.FromFormInterno:
		personaForm, err := getPersonaModelInputNormalizado(persona)
		if err != nil {
			return nil, err
		}
		personaRes = *personaForm
	// default:
	// 	return nil, errors.New("No se encontro parse para la fuente indicada")
	// }
	return &personaRes, nil
}

func getPersonaModelInputDenormalizado (persona *PersonaDenormalizada)(*models.Persona, error){

	personaModel:= models.Persona{

		IDFuente: getIDFromValue(persona.Metadata.Fuente, models.TableFuentes),
		Referente: persona.Metadata.ReferenteCarga,
		CompletedTime: persona.Metadata.CompletedTime,
		IDOrgaPertenencia: getIDFromValue(persona.Metadata.OrganizacionCarga.ValueOrZero(), models.TableOrganizaciones),
		Observaciones: persona.Metadata.Observaciones,

		Apellido: persona.DatosPersonales.Apellido,        
		Nombres:	persona.DatosPersonales.Nombres,
		IDTipoDocumento: int(getIDFromValue(persona.DatosPersonales.TipoDocumento, models.TableTiposDeDocumentos).ValueOrZero()),
		NroDocumento:  persona.DatosPersonales.NroDocumento,
		FechaNacimiento: persona.DatosPersonales.FechaNacimiento,
		IDIdentidadGenero: int(getIDFromValue(persona.DatosPersonales.IdentidadGenero, models.TableIdentidadesDeGenero).ValueOrZero()),
		
		Telefono:	persona.Contacto.Telefono,			
		Email: persona.Contacto.Email,				
		IDPaisNacimiento: getIDFromValue(persona.Contacto.PaisNacimiento.ValueOrZero(), models.TablePaises),
		IDNacionalidad: getIDFromValue(persona.Contacto.Nacionalidad.ValueOrZero(), models.TableNacionalidades),
		Edad: persona.DatosPersonales.Edad,
		IDNivelEducativo: getIDEducacion(persona.Educacion.NivelEducativo.ValueOrZero(), models.TableNivelesEducativos),
		IDCompletoEstudios: getIDEducacion(persona.Educacion.CompletoNivelEducativo.ValueOrZero(), models.TableNivelesEducativosConcluidos),
		Calle: persona.Contacto.Calle,
		Numero: persona.Contacto.Numero,
		Torre: persona.Contacto.Torre,
		Piso: persona.Contacto.Piso,
		Depto: persona.Contacto.Depto,

		FuePromotorxMinisterio: persona.Experiencia.FuePromotorxMinisterio,
		Ministerio: persona.Experiencia.Ministerio,
		OtrasTematicasFormo: persona.Experiencia.OtrasTematicasTrabajo,
		OtrasTematicasInteres: persona.Experiencia.OtrasTematicasInteres,
		OtrasTematicasCertifico: persona.Experiencia.OtrasTematicasCertifico,
		ObtuvoCertificacion: persona.Experiencia.ObtuvoCertificacion,
		ObtuvoRemuneracion: persona.Experiencia.ObtuvoRemuneracion,
		DeOrganismoDelEstadoYSociedadCivil: persona.Experiencia.DeOrganismoDelEstadoYSociedadCivil,
		DeOtroOrganismo: persona.Experiencia.DeOtroOrganismo,
		IDOrgaFormacion: getIDFromValue(persona.Experiencia.OrganizacionFormacion.ValueOrZero(), models.TableOrganizaciones),
	}

	//Parse de localidad/partido/provincia

	const provinciaBsAs = "06"

	idProvincia := getIDProvincia(persona.Contacto.Provincia)
	provinciaInicial := idProvincia //Guardo una copia para luego saber si difiere la ingresada de la finalmente guardada
	idPartido, partido := georef.GetDepartamento(persona.Contacto.Partido, idProvincia)
	idLocalidad, localidad := georef.GetLocalidad(persona.Contacto.Localidad, idPartido, idProvincia)

	if idProvincia.ValueOrZero() == provinciaBsAs && idLocalidad.IsZero() && idPartido.IsZero() {
		idLocalidad, localidad = georef.GetLocalidadEnCABA(persona.Contacto.Localidad)
	}

	salvarPartido(&idPartido, localidad)
	salvarProvincia(&idProvincia, localidad, partido)

	personaModel.IDLocalidad = idLocalidad
	personaModel.LocalidadIngresada = salvarLocalidad(localidad.ID == "", persona.Contacto.Localidad)
	personaModel.IDPartido = idPartido
	personaModel.PartidoIngresado = registrarPartidoIngresado(partido.ID == "", persona.Contacto.Partido)
	personaModel.IDProvincia = idProvincia
	personaModel.ProvinciaIngresada = registrarProvinciaIngresada(provinciaInicial != idProvincia, persona.Contacto.Provincia)

	return &personaModel, nil
}


func getPersonaModelInputNormalizado(persona *Persona)(*models.Persona, error){

	personaModel:= models.Persona{
		IDFuente: null.NewInt(int64(persona.Metadata.IDFuente), persona.Metadata.IDFuente != 0),
		Referente: persona.Metadata.ReferenteCarga,
		Apellido: persona.DatosPersonales.Apellido,        
		Nombres:	persona.DatosPersonales.Nombres,
		IDTipoDocumento: persona.DatosPersonales.IDTipoDocumento,
		NroDocumento:  persona.DatosPersonales.NroDocumento,
		//FechaNacimiento: null.TimeFrom(fechaNac),
		//IDIdentidadGenero: persona.DatosPersonales.IdentidadGenero,
		//CodArea:	persona.Contacto.CodArea,
		Telefono:	persona.Contacto.Telefono,			
		Email: persona.Contacto.Email,				
		IDProvincia: persona.Contacto.IDProvincia,	
		IDLocalidad: persona.Contacto.IDLocalidad,
		//OtrasTematicas: persona.Experiencia.OtrasTematicas,
		//IDOrgaPertenencia: null.IntFrom(persona.Experiencia.Organizacion),
		//IDPaisNacimiento: null.IntFrom(int64(persona.Contacto.PaisNacimiento)),
		//IDNacionalidad: getIDFromValue(persona.Contacto.Nacionalidad.ValueOrZero(), models.TableNacionalidades))),
		//Edad: null.IntFrom(int64(persona.DatosPersonales.Edad)),
		
		Calle: persona.Contacto.Calle,
		Numero: persona.Contacto.Numero,
		Torre: persona.Contacto.Torre,
		Piso: persona.Contacto.Piso,
		Depto: persona.Contacto.Depto,
	}

	return &personaModel, nil
}


func (persona *Persona) setResumenPersonaModel(personaModel models.Persona) error {
	
	// persona.NroSerie = personaModel.NroSerie
	 persona.DatosPersonales.Nombres = personaModel.Nombres
	 persona.DatosPersonales.Apellido = personaModel.Apellido
	 persona.DatosPersonales.NroDocumento = personaModel.NroDocumento
	 //persona.Contacto.IDLocalidad = personaModel.IDLocalidad
	 //persona.Contacto.IDProvincia = personaModel.IDProvincia
	 //persona.DatosPersonales.IdentidadGenero = personaModel.IdentidadGenero
	 //persona.DatosPersonales.IDTipoDocumento = personaModel.IDTipoDocumento

	return nil
}

func salvarLocalidad(faltaLocalidad bool, localidadPlana null.String) null.String {
	return null.NewString(localidadPlana.ValueOrZero(), faltaLocalidad)
}

func salvarPartido(idPartido *null.String, localidad *models.Localidad) {
	if !localidad.IDDepartamento.IsZero() && !localidad.IDDepartamento.Equal(*idPartido) {
		idPartido.SetValid(localidad.IDDepartamento.ValueOrZero())
	}
}

func salvarProvincia(idProvincia *null.String, localidad *models.Localidad, partido *models.Departamento){
	if idProvincia.ValueOrZero() != localidad.IDProvincia || idProvincia.ValueOrZero() != partido.IDProvincia {

		if localidad.ID != "" && partido.ID != "" {
			if localidad.IDProvincia == partido.IDProvincia {
				idProvincia.SetValid(localidad.IDProvincia)
			}

		} else if localidad.ID != "" {
			idProvincia.SetValid(localidad.IDProvincia)

		} else if partido.ID != "" {
			idProvincia.SetValid(partido.IDProvincia)
		}
	}
}

func registrarPartidoIngresado(faltaPartido bool, nombrePartido null.String) null.String {
	return null.NewString(nombrePartido.ValueOrZero(), faltaPartido)
}

func registrarProvinciaIngresada(cambioProvincia bool, nombreProvincia null.String) null.String {
	return null.NewString(nombreProvincia.ValueOrZero(), cambioProvincia)
}