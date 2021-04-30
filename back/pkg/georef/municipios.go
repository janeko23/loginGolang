package georef

import (
	"igualdad.mingeneros.gob.ar/pkg/services/log"

	"igualdad.mingeneros.gob.ar/pkg/services/db/models"

	"igualdad.mingeneros.gob.ar/pkg/services/db"
	"igualdad.mingeneros.gob.ar/pkg/services/db/services"
)

// Municipios struct
type Municipio struct {
	ID     string
	Nombre string
}

// LocalidadesPorProvincia Obtiene las localidades por provincia
func MunicipiosPorProvincia(idProvincia string) (*[]Municipio, error) {

	conn := db.GetDB()
	municipiosModel, err := services.MunicipiosPorProvincia(conn, idProvincia)

	if err != nil {
		log.Error("Ocurrio un error al listar municipios por provincia", log.String("idProvincia", idProvincia), log.String("error", err.Error()))
		return nil, err
	}

	municipios, _ := mapModelAMunicipio(municipiosModel) //TODO: manejar error

	return municipios, err
}

func mapModelAMunicipio(municipiosModel *[]models.Municipio) (*[]Municipio, error) {
	var municipios []Municipio

	for _, municipioModel := range *municipiosModel {

		var municipio Municipio

		err := municipio.setMunicipioModel(&municipioModel)
		if err != nil {
			return nil, err
		}

		municipios = append(municipios, municipio)
	}

	return &municipios, nil
}

func (municipio *Municipio) setMunicipioModel(municipioModel *models.Municipio) error { //TODO: manejar error

	municipio.ID = municipioModel.ID
	municipio.Nombre = municipioModel.Nombre

	return nil
}
