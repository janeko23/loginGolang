package models

import (
	"database/sql"
)

//User struct modela un usuario.
type User struct {
	ID       int
	Name     sql.NullString
	UserName string
	Hash     string // Pass hasheada
	Salt     string // Salt usado para hashear
}

//App struct modela una aplicación web (Diversidad, Igualdad, SICVG, etc...)
type App struct {
	ID   int
	Name string
	Host string
	port sql.NullString
}

//Resource struct modela un recurso de una determinada aplicación. Como las aplicaciones son APIs, un recurso es un endpoint de la API
type Resource struct {
	ID    int
	Path  string
	IDApp int
}

//Permission struct modela un permiso de un determinado usuario sobre un determinado recurso (endpoint)
type Permission struct {
	IDUser     int
	IDResource int
}

