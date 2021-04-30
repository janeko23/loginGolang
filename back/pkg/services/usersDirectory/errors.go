package usersdirectory

import "errors"

//TODO: ver descripciones de errores

// Errors
var (
	ErrUserDisabled          = errors.New("Usuarie deshabilitade")
	ErrRolDisabled           = errors.New("El rol está deshabilitado")
	ErrGrupoDisabled         = errors.New("El grupo está deshabilitado")
	ErrUserNotFound          = errors.New("Usuarie no existe")
	ErrUserSyntax            = errors.New("Sintaxis de usuarie incorrecta")
	ErrUserInvalidPassword   = errors.New("Contraseña de usuarie inválida")
	ErrUserEmptyPassword     = errors.New("Contraseña de usuarie vacía, ingrese contraseña")
	ErrInvalidPassword       = errors.New("Contraseña inválida")
	ErrInexistentUser        = errors.New("Usuarie inexistente")
	ErrConexionServidor	     = errors.New("Error de conexión con el servidor, intente nuevamente más tarde")
	ErrLoginNeeded           = errors.New("Es necesario loguearse para realizar esta acción")
	ErrUserExists            = errors.New("El nombre de usuarie ya existe. Elija otro nombre y vuelva a intentar")
)