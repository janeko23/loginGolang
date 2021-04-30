package main

import (
	
	"os"
	"fmt"

	"igualdad.mingeneros.gob.ar/router"
	"igualdad.mingeneros.gob.ar/pkg/services/db"
	"igualdad.mingeneros.gob.ar/pkg/services/log"
	"igualdad.mingeneros.gob.ar/pkg/libs/config"
	"igualdad.mingeneros.gob.ar/pkg/importar"

)

func main() {
	configFile := os.Getenv("CONFIGFILE")

	if configFile == "" {
		panic("CONFIGFILE not defined")
	}

	// Read Config
	cfg := config.NewConfig()
	err := cfg.ReadConfigFile(configFile)
	if err != nil {
		panic(err)
	}

	cfgData := cfg.Data()

	cfgLogger := cfgData.GetParamAsData("log")
	inicializarLogger(cfgLogger)

	cfgDb := cfgData.GetParamAsData("db")
	inicializarBD(cfgDb)

	fmt.Println("---------------------------------------")
	fmt.Println("Carga del Back finalizada correctamente")
	fmt.Println("---------------------------------------")
	
	const cargar = false
	
	if cargar {
		importar.CargarCSV(importar.PathToBase1)
		importar.CargarCSV(importar.PathToBase2)
	}


	cfgServer := cfgData.GetParamAsData("server")
	inicializarRouter(cfgServer)

}

func inicializarLogger(config *config.Data) {

	file := config.GetParamAsString("file")
	devMode := config.GetParamAsBool("devMode")
	log.Initialize(devMode, file)
}

func inicializarBD(config *config.Data) {

	dbHost := config.GetParamAsString("host")
	dbUserName := config.GetParamAsString("username")
	dbPassword := config.GetParamAsString("password")
	dbDatabase := config.GetParamAsString("database")
	
	devMode := config.GetParamAsBool("devMode")

	db.Initialize(dbHost, dbUserName, dbPassword, dbDatabase, devMode)
}

func inicializarRouter(config *config.Data) {
	puerto := config.GetParamAsString("port")

	router.Initialize(puerto)
}
