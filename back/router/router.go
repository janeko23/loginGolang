package router

import (
	"net/http"

	"igualdad.mingeneros.gob.ar/router/auth"
	"igualdad.mingeneros.gob.ar/router/georef"
	"igualdad.mingeneros.gob.ar/router/personas"
	"igualdad.mingeneros.gob.ar/router/categorias"
	"igualdad.mingeneros.gob.ar/router/usuaries"
	"fmt"
	"igualdad.mingeneros.gob.ar/pkg/services/log"

	"github.com/gorilla/mux"
)

// App router
type App struct {
	r    *mux.Router
	port string
}


// Initialize router
func Initialize(puerto string) {
	app := App{
		r:    mux.NewRouter(),
		port: "3000",
	}

	app.start()
}


func (a *App) start() {

	// Usuaries
	a.r.HandleFunc("/api/login", auth.Login).Methods("POST")
	a.r.HandleFunc("/api/altausuarie", usuaries.CrearUsuarie).Methods("POST")
	a.r.HandleFunc("/api/loginapi", auth.LoginApi).Methods("POST")
	
	// Personas
	a.r.HandleFunc("/api/personas", personas.ListarPersonasDatosResumidos).Methods("GET")
	a.r.HandleFunc("/api/pendientes", personas.ListarPendientesDatosResumidos).Methods("GET")
	a.r.HandleFunc("/api/existe-promotorx", personas.ExistePersona).Methods("GET")
	a.r.HandleFunc("/api/aprobar-persona/{id}", personas.AprobarPersona).Methods("GET")
	a.r.HandleFunc("/api/rechazar-persona/{id}", personas.RechazarPersona).Methods("GET")
	a.r.HandleFunc("/api/registro-personas", personas.RegistroPersonas).Methods("POST")

	// Categorias
	a.r.HandleFunc("/api/tipos-documentos", categorias.TiposDeDocumentos).Methods("GET")
	a.r.HandleFunc("/api/tematicas", categorias.Tematicas).Methods("GET")

	// Georef
	a.r.HandleFunc("/api/provincias", georef.Provincias).Methods("GET")
	a.r.HandleFunc("/api/localidades", georef.Localidades).Methods("GET")
	a.r.HandleFunc("/api/localidades-por-provincia/{id}", georef.LocalidadesPorProvincia).Methods("GET")


	err := http.ListenAndServe(":"+a.port, a.r)
	if err != nil {
		fmt.Println("error: ", err.Error())
		log.Fatal("Error al inicializar el router", log.String("error", err.Error()))
	} else {
		fmt.Println("El router se ha inicializado en ", a.port)
	}

}
