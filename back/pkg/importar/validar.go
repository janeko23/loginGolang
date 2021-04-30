package importar

import (
	"igualdad.mingeneros.gob.ar/pkg/personas"
	
	"strings"
	"reflect"

	"gopkg.in/guregu/null.v3"
	"github.com/go-playground/validator/v10"
)

func validateFields(p *personas.PersonaDenormalizada){

	validate := validator.New()
	validate.RegisterCustomTypeFunc(nullValidator, null.String{}, null.Int{})


	 err := validate.Struct(p)
	 if err != nil {
		obs := "Variables no registradas: "
		validationErrors := err.(validator.ValidationErrors)
		
		for idx, unError := range(validationErrors){
			pType := reflect.TypeOf(*p) //No necesito referencia porque no le hago cambios
			pValue := reflect.ValueOf(p).Elem() //Necesito referencia porque le hago cambios
			
			if idx != 0 { obs += ", " }
			obs += unError.Field()
			
			fieldName := unError.Namespace()
			fieldName = strings.ReplaceAll(fieldName, "PersonaDenormalizada.", "")
			splitedNames := strings.Split(fieldName, ".")
			
			var currentField reflect.StructField
			var found bool
			var indexes []int
			for _, currentName := range(splitedNames) {
				currentField, found = pType.FieldByName(currentName)

				if !found {
					break
				}

				indexes = append(indexes, currentField.Index[0])
				pType = currentField.Type
			}

			fieldValue := pValue.FieldByIndex(indexes)
			setFieldToNull(found, &fieldValue)
		}

		p.Metadata.Observaciones = null.StringFrom(obs)
	}
}

func setFieldToNull(found bool, value *reflect.Value) {
	if found {
		nullValue := getNullValue(value)
		value.Set(reflect.ValueOf(nullValue))
	}
}

func getNullValue(value *reflect.Value) interface{} {
	var nullValue interface{}
		if _, ok := value.Interface().(null.String); ok {
			nullValue = null.NewString("",false)
		}
		if _, ok := value.Interface().(null.Int); ok {
			nullValue = null.NewInt(0,false)
		}
		return nullValue
}

func nullValidator(field reflect.Value) interface{} {
	if valuer, ok := field.Interface().(null.String); ok {
		return valuer.String
	}
	if valuer, ok := field.Interface().(null.Int); ok {
		if valuer.Valid {
			return valuer.Int64
		}
	}
	return nil
}