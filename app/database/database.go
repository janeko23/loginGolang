package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

func AbrirConexion(driver, url string) error {
	DB, err := gorm.Open(driver, url)
	if err != nil {
		return err
	}
	fmt.Println("Inicializando la base de datos correctamente ...")

	DB.LogMode(true)
	// aca en auntomigrate faltarian los models
	DB.AutoMigrate()
	return nil
}
func CerrarConexion() error {
	return DB.Close()

}
