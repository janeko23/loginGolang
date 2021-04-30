package georef

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"igualdad.mingeneros.gob.ar/pkg/georef"
	"igualdad.mingeneros.gob.ar/router/common"
)

//Provincias metodo GET
func Provincias(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	response := provincias()

	w.WriteHeader(http.StatusOK)

	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		common.SendErr(w, http.StatusInternalServerError, err.Error())
	}
}

func provincias() map[string]interface{} {
	
	provincias, err := georef.Provincias()

	var response map[string]interface{}

	if err != nil {
		response = map[string]interface{}{
			"status": "ERROR",
			"error":  err.Error(),
		}
	} else {

		response = map[string]interface{}{
			"status": "OK",
			"data":   provincias,
		}
	}

	return response
}

//MunicipiosPorProvincia metodo GET
/*func MunicipiosPorProvincia(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)

	idProvincia := vars["idProvincia"]

	response := municipiosPorProvincia(idProvincia)

	w.WriteHeader(http.StatusOK)

	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		common.SendErr(w, http.StatusInternalServerError, err.Error())
	}
}*/

/*func municipiosPorProvincia(idProvincia string) map[string]interface{} {

	municipiosPorProvincia, err := georef.MunicipiosPorProvincia(idProvincia)
	var response map[string]interface{}

	if err != nil {
		response = map[string]interface{}{
			"status": "ERROR",
			"error":  err.Error(),
		}
	} else {

		response = map[string]interface{}{
			"status": "OK",
			"data":   municipiosPorProvincia,
		}
	}

	return response
}*/

//Localidades metodo GET
func Localidades(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	response := localidades()

	w.WriteHeader(http.StatusOK)

	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		common.SendErr(w, http.StatusInternalServerError, err.Error())
	}
}

func localidades() map[string]interface{} {
	
	localidades, err := georef.Localidades()

	var response map[string]interface{}

	if err != nil {
		response = map[string]interface{}{
			"status": "ERROR",
			"error":  err.Error(),
		}
	} else {

		response = map[string]interface{}{
			"status": "OK",
			"data": localidades,
		}
	}
	
	return response
}

//LocalidadesPorProvincia metodo GET
func LocalidadesPorProvincia(w http.ResponseWriter, r *http.Request){

	w.Header().Set("Content-Type", "application/json")
	
	vars := mux.Vars(r)
	
	idProvincia := vars["id"]

	response := localidadesPorProvincia(idProvincia)

	w.WriteHeader(http.StatusOK)

	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		common.SendErr(w, http.StatusInternalServerError, err.Error())
	}
}

func localidadesPorProvincia(idProvincia string) map[string]interface{} {
	
	localidades, err := georef.LocalidadesPorProvincia(idProvincia)

	var response map[string]interface{}

	if err != nil {
		response = map[string]interface{}{
			"status": "ERROR",
			"error":  err.Error(),
		}
	} else {

		response = map[string]interface{}{
			"status": "OK",
			"data": localidades,
		}
	}
	
	return response
}