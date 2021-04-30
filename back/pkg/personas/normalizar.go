package personas

import (
	"gopkg.in/guregu/null.v3"

	"fmt"
	"igualdad.mingeneros.gob.ar/pkg/services/db"
	"igualdad.mingeneros.gob.ar/pkg/services/db/services"
)

//obtenerIDElem devuelve el id del valor. Si no existe, lo registra
// Puede crear un registro en la DB
func obtenerIDElem(nombreElem string, nombreTabla string) (idRes int) {
	conn := db.GetDB()
	if idObtenido := services.ObtenerIDPorNombre(conn, nombreElem, nombreTabla); idObtenido != 0 {
		idRes = idObtenido
	} else {
		idNuevo := services.InsertarNormailzado(conn, nombreElem, nombreTabla);
		idRes = idNuevo
	}
	return idRes
}

func obtenerIDProvincia(nombreElem string) (string, error) {
	conn := db.GetDB()
	prov, err := services.Provincia(conn, nombreElem)
	if err != nil {
		return "", err
	}
	return prov.ID, nil
}

func getIDFromValue(nombre string, nombreTabla string) null.Int {
	var idRes int
	if len(nombre) > 0 {
		idRes = obtenerIDElem(nombre, nombreTabla)
	}

	validId := idRes != 0
		return null.NewInt(int64(idRes), validId)
}

func getIDProvincia(nombre null.String) (null.String) {
	idProvincia := null.NewString("",false)
	if nombre.IsZero() {
		return idProvincia
	}
	idRes, err := obtenerIDProvincia(nombre.ValueOrZero())
	if err != nil {
		fmt.Println("No existe la provincia ", nombre.ValueOrZero())
		return idProvincia
	}
	if len(idRes) > 0 {
		idProvincia.SetValid(idRes)
	}

	return idProvincia
}


func getIDEducacion(nombre string, nombreTabla string) null.Int {
	if nombre == ""{
		return null.NewInt(0, false)
	}
	return getIDFromValue(nombre, nombreTabla)
}