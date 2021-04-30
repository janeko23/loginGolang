package georef

import (
	"gopkg.in/guregu/null.v3"

	"igualdad.mingeneros.gob.ar/pkg/services/log"
	"igualdad.mingeneros.gob.ar/pkg/services/db/models"
	"igualdad.mingeneros.gob.ar/pkg/services/db"
	"igualdad.mingeneros.gob.ar/pkg/services/db/services"	

)

// Localidad struct
type Localidad struct {
	ID 	   string 
	Nombre string 
}

// Localidades Obtiene las provincias
func Localidades() (*[]Localidad, error) {
	
	conn := db.GetDB()
	localidadesModel, err := services.Localidades(conn)
	
	if err != nil {
		log.Error("Ocurrio un error al listar localidades", log.String("error", err.Error()))
		return nil, err
	}

	localidades, _ := mapModelALocalidad(localidadesModel) //TODO: manejar error

	return localidades, err
}

// LocalidadesPorProvincia Obtiene las localidades por provincia
func LocalidadesPorProvincia(idProvincia string) (*[]Localidad, error) {
	
	conn := db.GetDB()
	localidadesModel, err := services.LocalidadesPorProvincia(conn, idProvincia)
	
	if err != nil {
		log.Error("Ocurrio un error al listar localidades de provincia", log.String("idProvincia", idProvincia), log.String("error", err.Error()))
		return nil, err
	}

	localidades, _ := mapModelALocalidad(localidadesModel) //TODO: manejar error

	return localidades, err
}

func mapModelALocalidad(localidadesModel *[]models.Localidad) (*[]Localidad, error) {
	var localidades []Localidad

	for _, localidadModel := range *localidadesModel {

		var localidad Localidad

		err := localidad.setLocalidadModel(&localidadModel)
		if err != nil {
			return nil, err
		}

		localidades = append(localidades, localidad)
	}

	return &localidades, nil
}

func (localidad *Localidad) setLocalidadModel(localidadModel *models.Localidad) error { //TODO: manejar error

	localidad.ID = localidadModel.ID
	localidad.Nombre = localidadModel.Nombre

	return nil
}

// GetLocalidad Obtiene el modelo Localidad de una localidad dado el nombre como string
func GetLocalidad(nombre null.String, idDepto null.String, idProvincia null.String) (null.String, *models.Localidad) {
	idLocalidad := null.NewString("",false)

	if nombre.IsZero(){
		return idLocalidad, &models.Localidad{}
	}

	conn := db.GetDB()
	loc, _ := services.Localidad(conn, nombre.ValueOrZero(), idDepto, idProvincia)

	if loc.ID != "" {
		idLocalidad.SetValid(loc.ID)
	}

	return idLocalidad, loc
}

//GetLocalidadEnCABA si existe, obtiene id de la localidad de CABA con el nombre = nombre
func GetLocalidadEnCABA(nombre null.String) (null.String, *models.Localidad) {
	provCABA := null.StringFrom("02")
	idLocalidad := null.NewString("",false)

	if nombre.IsZero(){
		return idLocalidad, &models.Localidad{}
	}

	conn := db.GetDB()
	loc, _ := services.Localidad(conn, nombre.ValueOrZero(), null.NewString("",false), provCABA)

	if loc.ID != "" {
		idLocalidad.SetValid(loc.ID)
	}

	return idLocalidad, loc

}