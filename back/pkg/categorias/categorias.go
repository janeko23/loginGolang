package categorias

import (
	"igualdad.mingeneros.gob.ar/pkg/services/log"

	"igualdad.mingeneros.gob.ar/pkg/services/db"
	"igualdad.mingeneros.gob.ar/pkg/services/db/models"
	"igualdad.mingeneros.gob.ar/pkg/services/db/services"
)

// TiposDeDocumentos Obtiene los tipos de documentos
func TiposDeDocumentos() (*[]models.TipoDocumento, error) {

	conn := db.GetDB()
	categorias, err := services.TiposDeDocumentos(conn)
	if err != nil {
		log.Error("Ocurrio un error al obtener tipos de documentos", log.String("error", err.Error()))
	} else {
		log.Debug("Se listaron los tipos de documentos", log.Int("cantidadDeCategorias", len(*categorias)))
	}

	return categorias, err
}

// IDTipoDocumento Obtiene el ID del tipo de documento dado el nombre como string
func IDTipoDocumento(tipoDocumento string) (int, error) {

	conn := db.GetDB()
	tipoDocumentoModel, err := services.TipoDocumento(conn, tipoDocumento)

	if err != nil {
		log.Error("Ocurrio un error al obtener el id de la Experiencia Laboral", log.String("tipoDocumento", tipoDocumento), log.String("error", err.Error()))
		return 0, err
	}

	return tipoDocumentoModel.ID, err
}
/*
// Jurisdicciones Obtiene las jurisdicciones
func Jurisdicciones() (*[]models.Jurisdicciones, error) {

	conn := db.GetDB()
	jurisdicciones, err := services.Jurisdicciones(conn)
	if err != nil {
		log.Error("Ocurrio un error al obtener jurisdicciones", log.String("error", err.Error()))
	} else {
		log.Debug("Se listaron los tipos de jurisdicciones", log.Int("cantidadDeJurisdicciones", len(*jurisdicciones)))
	}

	return jurisdicciones, err
}

// Reparticiones Obtiene las reparticiones
func Reparticiones() (*[]models.Reparticiones, error) {

	conn := db.GetDB()
	reparticiones, err := services.Reparticiones(conn)
	if err != nil {
		log.Error("Ocurrio un error al obtener reparticiones", log.String("error", err.Error()))
	} else {
		log.Debug("Se listaron los tipos de reparticiones", log.Int("cantidadDeReparticiones", len(*reparticiones)))
	}

	return reparticiones, err
}

// AreasDeNacion Obtiene los tipos de documentos
func AreasDeNacion() (*[]models.AreaNacion, error) {

	conn := db.GetDB()
	areasDeNacion, err := services.AreasDeNacion(conn)
	if err != nil {
		log.Error("Ocurrio un error al obtener areas de Nacion", log.String("error", err.Error()))
	} else {
		log.Debug("Se listaron las areas de Nacion", log.Int("cantidadDeAreas", len(*areasDeNacion)))
	}

	return areasDeNacion, err
}

// Capacidades Obtiene los tipos de documentos
func Capacidades() (*[]models.Capacidad, error) {

	conn := db.GetDB()
	capacidades, err := services.Capacidades(conn)
	if err != nil {
		log.Error("Ocurrio un error al obtener tipos de capacidades", log.String("error", err.Error()))
	} else {
		log.Debug("Se listaron las capacidades", log.Int("Capacidades", len(*capacidades)))
	}

	return capacidades, err
}
*/
// Tematicas Obtiene las tematicas registradas
func Tematicas() (*[]models.Tematica, error) {

	conn := db.GetDB()
	tematicas, err := services.Tematicas(conn)
	if err != nil {
		log.Error("Ocurrio un error al obtener tematicas", log.String("error", err.Error()))
	} else {
		log.Debug("Se listaron las tematicas", log.Int("Tematicas", len(*tematicas)))
	}

	return tematicas, err
}

// NivelesEducativos Obtiene los tipos de niveles educativos
func NivelesEducativos() (*[]models.NivelEducativo, error) {

	conn := db.GetDB()
	nivelEducativo, err := services.NivelesEducativos(conn)
	if err != nil {
		log.Error("Ocurrio un error al obtener tipos de niveles educativos", log.String("error", err.Error()))
	} else {
		log.Debug("Se listaron los niveles educativos", log.Int("NivelEducativo", len(*nivelEducativo)))
	}

	return nivelEducativo, err
}

// NivelEducativoConcluido Obtiene los tipos de niveles educativos
func NivelEducativoConcluido() (*[]models.NivelEducativoConcluido, error) {

	conn := db.GetDB()
	nivelEducativoConcluido, err := services.NivelEducativoConcluido(conn)
	if err != nil {
		log.Error("Ocurrio un error al obtener tipos de niveles educativos concluidos", log.String("error", err.Error()))
	} else {
		log.Debug("Se listaron los niveles educativos concluidos", log.Int("NivelEducativoConcluido", len(*nivelEducativoConcluido)))
	}

	return nivelEducativoConcluido, err
}
/*

// Prestaciones Obtiene los tipos de prestaciones
func Prestaciones() (*[]models.Prestaciones, error) {

	conn := db.GetDB()
	prestaciones, err := services.Prestaciones(conn)
	if err != nil {
		log.Error("Ocurrio un error al obtener tipos de prestaciones", log.String("error", err.Error()))
	} else {
		log.Debug("Se listaron las prestaciones", log.Int("Prestaciones", len(*prestaciones)))
	}

	return prestaciones, err
}*/

// IDNivelEducativo Obtiene el ID del nivel educativo dado el nombre como string
func IDNivelEducativo(nivelEducativo string) (int, error) {

	conn := db.GetDB()
	nivelEducativoModel, err := services.NivelEducativo(conn, nivelEducativo)

	if err != nil {
		log.Error("Ocurrio un error al obtener el nivel educativo", log.String("nivelEducativo", nivelEducativo), log.String("error", err.Error()))
		return 0, err
	}

	return nivelEducativoModel.ID, err
}
/*
// IDSectorPublico Obtiene el ID del sector publico dado el nombre como string
func IDSectorPublico(sectorPublico string) (int, error) {

	conn := db.GetDB()
	sectorPublicoModel, err := services.SectorPublico(conn, sectorPublico)

	if err != nil {
		log.Error("Ocurrio un error al obtener el sector publico", log.String("sectorPublico", sectorPublico), log.String("error", err.Error()))
		return 0, err
	}

	return sectorPublicoModel.ID, err
}*/

// func mapTipoDocumentoAString(tiposDocumentosModel *[]models.TipoDocumento) []string {

// 	var tiposDocumentos []string

// 	for _, tipoDocumento := range *tiposDocumentosModel {
// 		tiposDocumentos = append(tiposDocumentos, tipoDocumento.TipoDocumento)
// 	}

// 	return tiposDocumentos
// }