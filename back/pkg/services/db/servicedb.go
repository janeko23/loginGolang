package db

import (

	"igualdad.mingeneros.gob.ar/pkg/services/log"

	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/driver/mysql"

	//"gorm.io/driver/sqlite"
	//_ "github.com/jinzhu/gorm/dialects/mysql"
)

var initialized bool

var (
	db *gorm.DB
)

// Initialize obtiene los datos de configFile
//Inicializa la DB
func Initialize(host string, userName string, password string, databaseName string, debug bool) {
	
	log.Info("Inicializando Base de Datos")

	if initialized { //TODO: hacer chequeo de db por nil
		log.Panic("La Base de Datos ya fue inicializada")
	}
	
	initialized = true

	logLevel := getLoggerLevel(debug)

	dsn := userName+":"+password+"@tcp("+host+")/"+databaseName+"?parseTime=true"
	d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logLevel),
	}) // TODO: ver como recibir la conexion
	if err != nil {
		log.Panic(err.Error())
	}

		
	db = d

	log.Info("Base de Datos inicializada")
}

// GetDB obtiene la base de datos
func GetDB() *gorm.DB {
	return db
}

func getLoggerLevel(debug bool) (logLevel logger.LogLevel) {
	if debug {
		logLevel = logger.Info
	} else {
		logLevel = logger.Silent
	}
	return logLevel
}