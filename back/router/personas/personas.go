package personas

import (
	
	"encoding/json"

	"net/http"
	"strconv"
	"github.com/gorilla/mux"
	//"igualdad.mingeneros.gob.ar/pkg/services/log"

	"igualdad.mingeneros.gob.ar/pkg/personas"
	"igualdad.mingeneros.gob.ar/pkg/services/db/models"
	"igualdad.mingeneros.gob.ar/router/common"
)

//DatosCompletos método GET
//Devuelve una persona con los datos completos. Utiliza el ID para buscarla
func DatosCompletos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	
	personaID := vars["id"]
	
	ID, err := strconv.Atoi(personaID)

	if err != nil {
		common.SendErr(w, http.StatusInternalServerError, err.Error())
	}

	response := datosCompletos(ID)
	w.WriteHeader(http.StatusOK)

	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		common.SendErr(w, http.StatusInternalServerError, err.Error())
	}
}

//ListarPersonasDatosResumidos método GET 
//Devuelve un listado con todos los registros de la DB
func ListarPersonasDatosResumidos(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	//TODO: deberia recibir un filtro con todos los campos requeridos
	filter := r.URL.Query().Get("requestData")

	response := listaDePersonas(filter, "aprobades")

	w.WriteHeader(http.StatusOK)

	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		common.SendErr(w, http.StatusInternalServerError, err.Error())
	}
}

//ListarPendientesDatosResumidos método GET 
//Devuelve un listado con todos los registros de la DB
func ListarPendientesDatosResumidos(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	//TODO: deberia recibir un filtro con todos los campos requeridos
	filter := r.URL.Query().Get("requestData")

	response := listaDePersonas(filter, "pendientes")

	w.WriteHeader(http.StatusOK)

	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		common.SendErr(w, http.StatusInternalServerError, err.Error())
	}
}

/*
//EditarDatos método PUT
//Buscar una persona por ID. Devuelve el registro completo.
//Actualiza los datos en la DB
func EditarDatos (w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	
	var personaUpdate *models.Persona
	json.NewDecoder(r.Body).Decode(&personaUpdate)
	
	vars := mux.Vars(r)
	
	personaID := vars["ID"]
	
	ID, err := strconv.Atoi(personaID)

	if err != nil {
		common.SendErr(w, http.StatusInternalServerError, err.Error())
	}
 
		
	personas.EditarDatos(ID, personaUpdate)
	w.WriteHeader(http.StatusOK)

	err = json.NewEncoder(w).Encode(personaUpdate)
	if err != nil {
		common.SendErr(w, http.StatusInternalServerError, err.Error())
	}

	
}
*/
//RegistroPersonas método POST
//Recibe el modelo completo desde el front.
//Crea un registro en la DB
func RegistroPersonas(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var personaNueva *personas.Persona
	json.NewDecoder(r.Body).Decode(&personaNueva)
		
	response := crearPersona(personaNueva)

	w.WriteHeader(http.StatusOK)

	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		common.SendErr(w, http.StatusInternalServerError, err.Error())
	}

}

//AprobarPersona metodo GET
//Aprueba una persona pasandola de la tabla de pendientes a la de personas
//Elimina y crea un registro en la DB
func AprobarPersona(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	personaID, _ := strconv.Atoi(vars["id"])
	response := aprobarPersona(personaID)

	w.WriteHeader(http.StatusOK)

	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		common.SendErr(w, http.StatusInternalServerError, err.Error())
	}

}

//RechazarPersona metodo GET
//Rechaza una persona eliminandola de la tabla de pendientes
//Elimina un registro en la DB
func RechazarPersona(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	personaID, _ := strconv.Atoi(vars["id"])
	response := rechazarPersona(personaID)

	w.WriteHeader(http.StatusOK)

	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		common.SendErr(w, http.StatusInternalServerError, err.Error())
	}

}

//ExistePersona metodo GET
//Indica si existe la persona con el tipo y número de documento pasados por parametro
//No realiza cambios en la DB
func ExistePersona(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	query := r.URL.Query()
	tipoDoc := query.Get("tipoDoc")
	nroDoc := query.Get("nroDoc")

	response := esPromotorx(tipoDoc, nroDoc)

	w.WriteHeader(http.StatusOK)

	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		common.SendErr(w, http.StatusInternalServerError, err.Error())
	}

}

func crearPersona(personaNueva *personas.Persona) map[string]interface{} {
	
	id, err := personas.AltaPersona(personaNueva, "personas", "formInterno")

	response := getResponse(err)
	if response["status"] == "OK" {
		response["data"] = id
	}
	
	return response
}

func aprobarPersona(idPersona int) map[string]interface{} {
	
	idNuevo, err := personas.AprobarPersona(idPersona)

	response := getResponse(err)
	if response["status"] == "OK" {
		response["data"] = idNuevo
	}
	
	return response
}

func rechazarPersona(idPersona int) map[string]interface{} {
	
	err := personas.RechazarPersona(idPersona)

	response := getResponse(err)
	
	return response
}

func datosCompletos(ID int)map[string]interface{}{

	personaDetails, err := personas.BuscarPorID(ID)

	response := getResponse(err)
	if response["status"] == "OK" {
		response["data"] = personaDetails
	}

	return response

}

func listaDePersonas(filter string, clasificacion string) map[string]interface{} {

	listaPersonas, err := personas.ListarPersonasDatosResumidos(filter, clasificacion)

	response := getResponse(err)
	if response["status"] == "OK" {
		response["data"] = listaPersonas
	}

	return response
}

func esPromotorx(tipoDoc string, nroDoc string) (response map[string]interface{}) {

	numericTipoDoc, err := strconv.Atoi(tipoDoc)

	if err != nil {		
		response = getResponse(err)

	} else {
		hayPromotorx, err := personas.ExistePromotorxConDoc(models.TablePersonas, numericTipoDoc, nroDoc)

		response = getResponse(err)
		if response["status"] == "OK" {
			response["data"] = hayPromotorx
		}
	}

	return response
}



func getResponse(err error) (response map[string]interface{}) {
	if err != nil {
		response = map[string]interface{}{
			"status": "ERROR",
			"error":  err.Error(),
		}
	} else {
		response = map[string]interface{}{
			"status": "OK",
		}
	}
	return response

}

