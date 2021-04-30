package models
import (
	"reflect"
)

//TablaAModelo map
var TablaAModelo map[string]interface{} = map[string]interface{}{
	"organismos": reflect.TypeOf(Organismo{}),
}