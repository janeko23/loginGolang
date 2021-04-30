package usuaries

import (
	"encoding/json"
	"net/http"

	"igualdad.mingeneros.gob.ar/pkg/services/usersDirectory/models"
	"igualdad.mingeneros.gob.ar/pkg/usuaries"
	"igualdad.mingeneros.gob.ar/router/common"
)

// CrearUsuarie método POST crea usuarie en el directorio de ldap
func CrearUsuarie(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var usuarie models.Usuarie

	json.NewDecoder(r.Body).Decode(&usuarie)

	response := crearUsuarie(&usuarie)

	w.WriteHeader(http.StatusOK)

	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		common.SendErr(w, http.StatusInternalServerError, err.Error())
	}
}

func crearUsuarie(usuarie *models.Usuarie) map[string]interface{} {

	err := usuaries.AltaUsuarie(usuarie)

	var response map[string]interface{}

	if err != nil {
		response = map[string]interface{}{
			"status": "ERROR",
			"error":  err.Error(),
		}
	} else {
		response = map[string]interface{}{
			"status": "OK",
			"data":   "Creación de Usuarie OK!",
		}
	}

	return response
}
