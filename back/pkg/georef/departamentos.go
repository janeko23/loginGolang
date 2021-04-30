package georef

import (
	"gopkg.in/guregu/null.v3"
    
	"igualdad.mingeneros.gob.ar/pkg/services/db/models"
	"igualdad.mingeneros.gob.ar/pkg/services/db"
	"igualdad.mingeneros.gob.ar/pkg/services/db/services"	

)

// GetDepartamento Obtiene el modelo Departamento de un partido/departamento dado el nombre como string
func GetDepartamento(nombre null.String, idProvincia null.String) (null.String, *models.Departamento) {

	idPartido := null.NewString("",false)

	if nombre.IsZero(){
		return idPartido, &models.Departamento{}
	}
	conn := db.GetDB()
	dep, _ := services.Departamento(conn, nombre.ValueOrZero(), idProvincia)

	if dep.ID != "" && (idProvincia.IsZero() || dep.IDProvincia == idProvincia.ValueOrZero()) {
		idPartido.SetValid(dep.ID)
	}
	return idPartido, dep
}