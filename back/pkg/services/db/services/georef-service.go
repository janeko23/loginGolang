package services

import (
	"gorm.io/gorm"

	"igualdad.mingeneros.gob.ar/pkg/services/db/models"

	"gopkg.in/guregu/null.v3"

	"errors"

)

// Provincias Servicio que obtiene las provincias
func Provincias(conn *gorm.DB) (*[]models.Provincia, error) {
	tablaProvincias := models.TableProvincias

	var provincias []models.Provincia
	result := conn.Table(tablaProvincias).Find(&provincias)

	return &provincias, result.Error
}

// MunicipiosPorProvincia Servicio que obtiene los municipios asociadas a un ID de Provincia
func MunicipiosPorProvincia(conn *gorm.DB, idProvincia string) (*[]models.Municipio, error) {
	tablaMunicipios := models.TableMunicipios

	var municipios []models.Municipio
	result := conn.Table(tablaMunicipios).Where("IdProvincia = ? ", idProvincia).Find(&municipios)

	return &municipios, result.Error
}

// Localidades Servicio que obtiene las localidades
func Localidades(conn *gorm.DB) (*[]models.Localidad, error) {
	tablaLocalidades := models.TableLocalidades

	var localidades []models.Localidad
	result := conn.Table(tablaLocalidades).Find(&localidades)

	return &localidades, result.Error
}

// LocalidadesPorProvincia Servicio que obtiene las localidades asociadas a un ID de Provincia
func LocalidadesPorProvincia(conn *gorm.DB, idProvincia string) (*[]models.Localidad, error) {
	tablaLocalidades := models.TableLocalidades

	var localidades []models.Localidad
	result := conn.Table(tablaLocalidades).Where("IdProvincia = ? ", idProvincia).Find(&localidades)

	return &localidades, result.Error
}

// Provincia Servicio que obtiene una provincia dado el Nombre de la provincia (como string). El resultado es un modelo de la provincia con su ID y su Nombre
func Provincia(conn *gorm.DB, provincia string) (*models.Provincia, error) {
	tablaProvincias := models.TableProvincias

	var provinciaModel models.Provincia

	result := conn.Table(tablaProvincias).Where("Nombre = ? ", provincia).Find(&provinciaModel)

	return &provinciaModel, result.Error
}

// Departamento Servicio que obtiene un departamento dado el Nombre de la provincia (como string). El resultado es un modelo de la provincia con su ID y su Nombre
func Departamento(conn *gorm.DB, depto string, idProvincia null.String) (*models.Departamento, error) {
	var deptoModel models.Departamento
	

	condiciones := 	map[string]interface{}{
		"Nombre": depto,
	}
	if !idProvincia.IsZero(){ 
		condiciones["IdProvincia"] = idProvincia
	}

	result := conn.Limit(2).Table(models.TableDepartamentos).Where(condiciones).Find(&deptoModel)

	if result.RowsAffected > 1 {
		return &deptoModel, errors.New("No se pudo determinar el departamento")
	}

	return &deptoModel, result.Error
}

// Localidad Servicio que obtiene una localidad dado el Nombre de la localidad (como string). El resultado es un modelo de la Localidad con su ID y su Nombre
func Localidad(conn *gorm.DB, localidad string, idDepto null.String, idProvincia null.String) (*models.Localidad, error) {
	
	tablaLocalidades := models.TableLocalidades
	var localidadModel models.Localidad

	condiciones := 	map[string]interface{}{
		"Nombre": localidad,
	}

	if !idProvincia.IsZero(){
		condiciones["IdProvincia"] = idProvincia.ValueOrZero()	

		resultPorProvincia := conn.Limit(2).Table(tablaLocalidades).Where(condiciones).Find(&localidadModel)

		if resultPorProvincia.RowsAffected == 1 {
			return &localidadModel, resultPorProvincia.Error
		}
	}
	
	//Llego acá tanto si no hay provincia como si hay +1 localidad con el mismo nombre en la provincia
	if !idDepto.IsZero() { 
		condiciones["IdDepartamento"] = idDepto.ValueOrZero()
	}

	resultPorDepto := conn.Table(tablaLocalidades).Where(condiciones).Find(&localidadModel)

	if resultPorDepto.RowsAffected != 1 {
		return &localidadModel, errors.New("No se pudo determinar la localidad " + localidad)
	}

	return &localidadModel, resultPorDepto.Error

}