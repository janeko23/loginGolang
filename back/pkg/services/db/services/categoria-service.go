package services

import (

	"fmt"
	"strings"
	"gorm.io/gorm"
	"igualdad.mingeneros.gob.ar/pkg/services/db/models"
)

// TiposDeDocumentos Servicio para obtener las categorias de documentos
func TiposDeDocumentos(conn *gorm.DB) (*[]models.TipoDocumento, error) {
	tablaTiposDeDocumentos := models.TableTiposDeDocumentos

	var categorias []models.TipoDocumento
	result := conn.Table(tablaTiposDeDocumentos).Find(&categorias)

	return &categorias, result.Error
}

// TipoDocumento Servicio que obtiene el tipo de documento dado el nombre (como string). El resultado es un modelo del tipo de documento con su ID y su Nombre
func TipoDocumento(conn *gorm.DB, tipoDocumento string) (*models.TipoDocumento, error) {
	tablaTiposDeDocumentos := models.TableTiposDeDocumentos

	var tipoDocumentoModel models.TipoDocumento

	result := conn.Table(tablaTiposDeDocumentos).Where("Nombre = ? ", strings.ToUpper(tipoDocumento)).Find(&tipoDocumentoModel)

	return &tipoDocumentoModel, result.Error
}


// NivelesEducativos Servicio para obtener las categorias de niveles educativos
func NivelesEducativos(conn *gorm.DB) (*[]models.NivelEducativo, error) {
	tablaNivelesEducativos := models.TableNivelesEducativos

	var nivelEducativo []models.NivelEducativo
	result := conn.Table(tablaNivelesEducativos).Find(&nivelEducativo)

	return &nivelEducativo, result.Error
}

// NivelEducativoConcluido Servicio para obtener las categorias de niveles educativos
func NivelEducativoConcluido(conn *gorm.DB) (*[]models.NivelEducativoConcluido, error) {

	var nivelEducativoConcluido []models.NivelEducativoConcluido
	result := conn.Table("nivel_educativo_concluido").Find(&nivelEducativoConcluido)

	return &nivelEducativoConcluido, result.Error
}



// NivelEducativo Servicio que obtiene el nivel educativo dado el nombre (como string). El resultado es un modelo del nivel educativo con su ID y su Nombre
func NivelEducativo(conn *gorm.DB, nivelEducativo string) (*models.NivelEducativo, error) {
	tablaNivelesEducativos := models.TableNivelesEducativos

	var nivelEducativoModel models.NivelEducativo

	result := conn.Table(tablaNivelesEducativos).Where("Nombre = ? ", strings.ToUpper(nivelEducativo)).Find(&nivelEducativoModel)

	return &nivelEducativoModel, result.Error
}

// Tematicas Servicio para obtener las tematicas
func Tematicas(conn *gorm.DB) (*[]models.Tematica, error) {

	var tematicas []models.Tematica
	result := conn.Table("tematicas").Find(&tematicas)

	return &tematicas, result.Error
}

//ObtenerIDPorNombre funcion
func ObtenerIDPorNombre(conn *gorm.DB, nombre string, nombreTabla string) int {
	nombreEnMinuscula := strings.Trim(strings.ToLower(nombre), " ")
	nombreSanitizado := strings.ReplaceAll(nombreEnMinuscula, "á", "a")
	nombreSanitizado = strings.ReplaceAll(nombreSanitizado, "é", "e")
	nombreSanitizado = strings.ReplaceAll(nombreSanitizado, "í", "i")
	nombreSanitizado = strings.ReplaceAll(nombreSanitizado, "ó", "o")
	nombreSanitizado = strings.ReplaceAll(nombreSanitizado, "ú", "u")

	var modeloNormalizado models.NormalizacionGenerica
	conn.Table(nombreTabla).Where("Nombre = ? ", nombreSanitizado).Limit(1).Find(&modeloNormalizado)
	return modeloNormalizado.ID
}

//InsertarNormailzado inserta una nuevo valor con el nombre pasado por parametro en la tabla pasada por parámetro
func InsertarNormailzado(conn *gorm.DB, nombre string, nombreTabla string) int {
	nombreEnMinuscula := strings.ToLower(nombre)
	nombreSanitizado := strings.ReplaceAll(nombreEnMinuscula, "á", "a")
	nombreSanitizado = strings.ReplaceAll(nombreSanitizado, "é", "e")
	nombreSanitizado = strings.ReplaceAll(nombreSanitizado, "í", "i")
	nombreSanitizado = strings.ReplaceAll(nombreSanitizado, "ó", "o")
	nombreSanitizado = strings.ReplaceAll(nombreSanitizado, "ú", "u")
	modeloNormalizado := models.NormalizacionGenerica{Nombre: nombreSanitizado}
	res := conn.Table(nombreTabla).Create(&modeloNormalizado)
	if (res.Error != nil){
		fmt.Println("Hubo un error al insertar un valor en",nombreTabla, ":\n", res.Error.Error())
	}
	return modeloNormalizado.ID
}