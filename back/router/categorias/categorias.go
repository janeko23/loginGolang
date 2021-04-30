package categorias

import (
	"encoding/json"
	"net/http"

	"igualdad.mingeneros.gob.ar/pkg/categorias"
	"igualdad.mingeneros.gob.ar/router/common"
)

//TiposDeDocumentos método GET
//Obtiene la lista de tipos de documentos de la DB
func TiposDeDocumentos(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	response := obtenerTiposDeDocumentos()

	w.WriteHeader(http.StatusOK)

	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		common.SendErr(w, http.StatusInternalServerError, err.Error())
	}
}



//NivelesEducativos método GET
//Obtiene la lista de niveles educativos de la DB
func NivelesEducativos(w http.ResponseWriter, r *http.Request) {
/*
	w.Header().Set("Content-Type", "application/json")

	response := obtenerNivelEducativo()

	w.WriteHeader(http.StatusOK)

	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		common.SendErr(w, http.StatusInternalServerError, err.Error())
	}*/
}

//NivelEducativoConcluido método GET
//Obtiene la lista de niveles educativos de la DB
func NivelEducativoConcluido(w http.ResponseWriter, r *http.Request) {

	/*w.Header().Set("Content-Type", "application/json")

	response := obtenerNivelEducativoConcluido()

	w.WriteHeader(http.StatusOK)

	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		common.SendErr(w, http.StatusInternalServerError, err.Error())
	}*/
}



func obtenerTiposDeDocumentos() map[string]interface{} {

	categorias, err := categorias.TiposDeDocumentos()

	var response map[string]interface{}

	if err != nil {
		response = map[string]interface{}{
			"status": "ERROR",
			"error":  err.Error(),
		}
	} else {

		response = map[string]interface{}{
			"status": "OK",
			"data":   categorias,
		}
	}

	return response
}





func obtenerNivelEducativo() map[string]interface{} {

	nivelEducativo, err := categorias.NivelesEducativos()

	var response map[string]interface{}

	if err != nil {
		response = map[string]interface{}{
			"status": "ERROR",
			"error":  err.Error(),
		}
	} else {

		response = map[string]interface{}{
			"status": "OK",
			"data":   nivelEducativo,
		}
	}

	return response
}

func obtenerNivelEducativoConcluido() map[string]interface{} {

	nivelEducativoConcluido, err := categorias.NivelEducativoConcluido()

	var response map[string]interface{}

	if err != nil {
		response = map[string]interface{}{
			"status": "ERROR",
			"error":  err.Error(),
		}
	} else {

		response = map[string]interface{}{
			"status": "OK",
			"data":   nivelEducativoConcluido,
		}
	}

	return response
}

//Tematicas método GET
//Obtiene la lista de tematicas de la DB
func Tematicas(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	response := obtenerTematicas()

	w.WriteHeader(http.StatusOK)

	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		common.SendErr(w, http.StatusInternalServerError, err.Error())
	}
}

func obtenerTematicas() map[string]interface{} {

	tematicas, err := categorias.Tematicas()

	var response map[string]interface{}

	if err != nil {
		response = map[string]interface{}{
			"status": "ERROR",
			"error":  err.Error(),
		}
	} else {

		response = map[string]interface{}{
			"status": "OK",
			"data":   tematicas,
		}
	}

	return response
}
