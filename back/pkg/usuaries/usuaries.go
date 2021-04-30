package usuaries

import (
	"igualdad.mingeneros.gob.ar/pkg/services/usersDirectory"
	"igualdad.mingeneros.gob.ar/pkg/services/usersDirectory/models"
)

// AltaUsuarie da de alta a un nuevo usuarie en el directorio de ldap
func AltaUsuarie(usuarie *models.Usuarie) error {

	// TODO: codificar contrasena

	err := usersdirectory.CrearUsuarie(usuarie)
	if err != nil {
		return err
	}

	return nil
}