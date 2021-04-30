package importar

import (
	"gopkg.in/guregu/null.v3"
	"regexp"
)


func parseDepartamento(localidad null.String, depto string) string {

	if localidad.ValueOrZero() == "CIUDAD DE BUENOS AIRES" {
		return ""
	}
	regexDepto := regexp.MustCompile("\\s+DEP(ARTAMEN)?TO\\s+")
	depto = regexDepto.ReplaceAllString(depto, "")
	
	depto = eliminarPuntuacion(depto)

	depto = normalizarAbreviaturas(depto)

	return depto
}