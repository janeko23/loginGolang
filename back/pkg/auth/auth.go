package auth

import (
	
	"igualdad.mingeneros.gob.ar/pkg/services/usersDirectory"
)

// InputLogin login data
type InputLogin struct {
	User     string `json:"user"`
	Password string `json:"password"` // TODO: poner restricciones a la forma de la password
}

// Login function
func (input *InputLogin) Login() error {

	err := usersdirectory.Login(input.User, input.Password) //TODO: pasar la pass hasheada
	if err != nil {
		return err
	} 
	
	return nil
}