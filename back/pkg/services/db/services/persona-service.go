package services

import (

	//"database/sql"
	
    "gorm.io/gorm"

	"igualdad.mingeneros.gob.ar/pkg/services/db/models"

	"strconv"
	//"strings"
	"gopkg.in/guregu/null.v3"

	"errors"

)

// ListarPersonasDatosResumidos Servicio que obtiene lista de personas registradas.
// Logica de filtro interesPCL
func ListarPersonasDatosResumidos(conn *gorm.DB, clasificacion string, filtroTematica string) (*[]models.ResumenPersona, error) {

	var personasResumidas []models.ResumenPersona

	filtroNumerico , _ := strconv.Atoi(filtroTematica)

	var nombreTabla string
	if clasificacion == "aprobades"{
		nombreTabla = "personas"
	} else {
		nombreTabla = "pendientes"
	}

	res := queryPersonasResumidas(conn, nombreTabla, filtroNumerico).Scan(&personasResumidas)

	return &personasResumidas, res.Error
}

func queryPersonasResumidas(conn *gorm.DB, nombreTabla string, filterVal int) *gorm.DB {
	result := conn.Table(nombreTabla).
	Select(nombreTabla+`.Id, `+nombreTabla+`.Nombres, `+nombreTabla+`.Apellido, provincias.Nombre as Provincia, organizaciones.Nombre as Organizacion`).
	Joins("LEFT JOIN localidades ON localidades.Id = ?",nombreTabla+".IdLocalidad").
	Joins("LEFT JOIN tematicas ON tematicas.Id = "+ nombreTabla+".IdTematica").
	Joins("LEFT JOIN provincias ON provincias.Id = "+nombreTabla+".IdProvincia").
	Joins("LEFT JOIN organizaciones ON organizaciones.Id = "+nombreTabla+".IdOrganizacion")

	return result

}


// FiltrarPorLocalidad metodo para listar personas por localidad
func FiltrarPorLocalidad(conn *gorm.DB, localidad string) (*[]models.Persona, error) {

	tablaPersonas := models.TablePersonas

	var personas []models.Persona
	
	result := conn.Table(tablaPersonas).Where("Localidad = ? ", localidad).Find(&personas)

	return &personas, result.Error
}


//BuscarPersonaPorID Servicio que obtiene una persona a partir de un ID con los datos completos
func BuscarPersonaPorID(conn *gorm.DB, nombreTabla string, ID int) (*models.Persona, error) {

	tabla := nombreTabla
	var persona models.Persona
	result := conn.Table(tabla).Where("Id = ?", ID).Take(&persona)

	return &persona, result.Error
}

//ActualizarDatos Servicio que actualiza datos de persona por ID 
func ActualizarDatos(conn *gorm.DB, personaActual *models.Persona) {

	conn.Where("Id = ?", personaActual.ID).Save(&personaActual)
}

//ActualizarColumnas Servicio que actualiza columnas pasadas de una persona por ID 
func ActualizarColumnas(conn *gorm.DB, personaActual *models.Persona, personaAMergear *models.Persona, columnas []string) (int, error) {
	res := conn.Model(personaActual).Select(columnas).Updates(personaAMergear)
	return personaActual.ID, res.Error
}

//BuscarPersonaPorTipoYDocumento Servicio que obtiene una persona con los datos completos
func BuscarPersonaPorTipoYDocumento(conn *gorm.DB, nombreTabla string, tipoDoc int, nroDoc null.String) (*models.Persona, error) {
	 
	tablaPersonas := nombreTabla

	if !nroDoc.Valid {
		return nil, errors.New("El número de documento no es válido")
	}

	var persona models.Persona
		
	result := conn.Table(tablaPersonas).Where("IdTipoDoc = ? AND NroDoc = ?",tipoDoc ,nroDoc.String ).Limit(1).Find(&persona)
	
	if result.Error != nil {
		return nil, result.Error
	}

	if persona.ID != 0 {
		return &persona, nil
	}

	return nil, nil
}

//BuscarTematicasDePersona busca tematicas trabajadas por la persona con id idPersona en la tabla dada por nombreTabla
func BuscarTematicasDePersona(conn *gorm.DB, nombreTabla string, idPersona int) ([]int, error){
	var exps []models.ExpEnTematica
	result := conn.Table(nombreTabla).Select("IdTematica").Where("IdPersona = ?", idPersona ).Find(&exps)

	var res []int
	for _, exp := range exps {
		res = append(res, exp.IDTematica)
	}
	return res, result.Error
}

//TieneRegistros indica si la persona tiene registrada alguna entrada en la tabla de tematicas pasada por parametro
func TieneRegistros(conn *gorm.DB, nombreTabla string, idPersona int) (bool, error){
	
	//var exps []models.ExpEnTematica
	var ids []int
	result := conn.Table(nombreTabla).Select("IdPersona").Where("IdPersona = ?", idPersona ).Find(&ids)
	tieneRegistros := result.RowsAffected > 0
	return tieneRegistros, result.Error
}


//InsertarNuevaPersona inserta persona en tabla historial
func InsertarNuevaPersona(conn *gorm.DB, persona models.Persona, tabla string) (int, error) {
	
	tablaPersonas := tabla
	persona.ID = 0

	result := conn.Table(tablaPersonas).Create(&persona)
	
	return persona.ID, result.Error
}

//InsertarNuevasPersonas inserta persona en tabla historial
func InsertarNuevasPersonas(conn *gorm.DB, personas *[]models.Persona, tabla string) ([]int, error) {
	
	tablaPersonas := tabla

	result := conn.Table(tablaPersonas).Create(personas)
	
	var ids []int = make([]int, len(*personas))
	for idx, persona := range(*personas) {
		ids[idx] = persona.ID
	}
	
	return ids, result.Error
}

// InsertarExpEnPersonasTematicas Servicio que inserta las experiencias en tematicas de una persona
func InsertarExpEnPersonasTematicas(conn *gorm.DB, nombreTabla string, ExpPersonaEnTematicas *[]models.ExpEnTematica) error {
	
	tablaPersonasTematicas := nombreTabla

	result := conn.Table(tablaPersonasTematicas).Create(ExpPersonaEnTematicas)
	
	return result.Error
}

//EliminarExperienciasEnTematicas elimina experiencias de pendientes en pendientes_tematicas
func EliminarExperienciasEnTematicas(conn *gorm.DB, nombreTabla string, idPersona int) error{

	var exps []models.ExpEnTematica
	result := conn.Table(nombreTabla).Where("IdPersona = ?", idPersona ).Delete(&exps)
	return result.Error
}

func InsertarCertificacionesPorOrganismos(conn *gorm.DB, certificaciones *[]models.CertificacionDeOrganismo) error {
	result := conn.Table(models.TableCertificacionesOrganismos).Create(certificaciones)
	return result.Error
}


//EliminarPersona elimina registro de persona de la tabla indicada por NombreTabla
func EliminarPersona(conn *gorm.DB, nombreTabla string, idPersona int) error {

	result := conn.Table(nombreTabla).Delete(models.Persona{}, idPersona)
	return result.Error
}

//getIdsTiposDoc devuelve un slice con los ids de los tipos de documento
func getIdsTiposDoc(tiposDoc *[]models.TipoDocumento) []int {
	var idsTipoDocumento []int
	for _, tipoDoc := range *tiposDoc {
		idsTipoDocumento = append(idsTipoDocumento, tipoDoc.ID)
	}

	return idsTipoDocumento
}