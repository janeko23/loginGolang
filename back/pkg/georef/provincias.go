package georef

import ( 
    
	"igualdad.mingeneros.gob.ar/pkg/services/log"

	"igualdad.mingeneros.gob.ar/pkg/services/db/models"

	"igualdad.mingeneros.gob.ar/pkg/services/db"
	"igualdad.mingeneros.gob.ar/pkg/services/db/services"	
)

// Provincias Obtiene las provincias
func Provincias() (*[]models.Provincia, error) {
	
	conn := db.GetDB()
	provincias, err := services.Provincias(conn)
	
	if err != nil {
		log.Error("Ocurrio un error al listar provincias", log.String("error", err.Error()))
		return nil, err
	}

	return provincias, err
}

// IDProvincia Obtiene el ID una provincia dado el nombre como string
func IDProvincia(nombreProvincia string) (string, error) {

	conn := db.GetDB()
	provincia, err := services.Provincia(conn, nombreProvincia)
	
	if err != nil {
		log.Error("Ocurrio un error al obtener provincia", log.String("provincia", nombreProvincia), log.String("error", err.Error()))
		return "", err
	}

	return provincia.ID, err
}