package importar

import (
	"regexp"
	"strings"
)

func parseLocalidad(texto string) string {
	texto = eliminarPuntuacion(texto)

	texto = normalizarAbreviaturas(texto)

	texto = parseSSDeJujuy(texto)

	texto = parseLocalidadCABA(texto)

	if strings.Contains(texto, "CALZADA"){
		texto = "RAFAEL CALZADA"
	}

	if strings.Contains(texto, "BARILOCHE"){
		texto = "SAN CARLOS DE BARILOCHE"
	}

	regexTresPozos := regexp.MustCompile("S\\s.*TRES\\sPOZOS")
	if regexTresPozos.MatchString(texto) {
		texto = "SANTUARIO DE TRES POZOS"
	}

	if strings.Contains(texto, "CORDOBITA"){
		texto = "CORDOBITA"
	}

	if strings.Contains(texto, "EDUVIGIS"){
		texto = "EDUVIGIS"
	}

	if strings.Contains(texto, "CASANOVA"){
		texto = "ISIDORO CASANOVA"
	}

	if strings.Contains(texto, "R CASTILLO"){
		texto = "RAFAEL CASTILLO"
	}

	if strings.Contains(texto, "V DEL PINO"){
		texto = "VIRREY DEL PINO"
	}


	if strings.Contains(texto, "GENERAL SAVIO"){
		texto = "VILLA GENERAL SAVIO"
	}

	if strings.Contains(texto, "BRANDSEN"){
		texto = "CORONEL BRANDSEN"
	}

	if strings.Contains(texto, "REMANSO"){
		texto = "EL REMANSO"
	}

	if strings.Contains(texto, "VIRREYES"){
		texto = "VIRREYES"
	}

	if strings.Contains(texto, " MORRIS"){
		texto = "WILLIAM C. MORRIS"
	}

	if strings.Contains(texto, "HURLINGHAM"){
		texto = "HURLINGHAM"
	}

	return texto
}

func parseLocalidadCABA(localidad string) string {
	regexDosLocalidades := regexp.MustCompile("\\s*\\-\\s*C\\s?A\\s?B\\s?A\\s*")
	localidad = regexDosLocalidades.ReplaceAllString(localidad, "")

	if regexCABA := regexp.MustCompile("^(C\\s?A\\s?B\\s?A\\s?|CIUDAD.*BUENOS AIRES)$"); regexCABA.MatchString(localidad){
		localidad = "CIUDAD DE BUENOS AIRES"
	}
	return localidad
}


func parseSSDeJujuy(text string) string {
	text = strings.ToUpper(text)
	if regexJujuy := regexp.MustCompile("S(AN)?\\s?S((ALVA)?DOR)?\\s{0,2}DE\\sJUJUY"); regexJujuy.MatchString(text) {
		text ="SAN SALVADOR DE JUJUY"
	}

	return text
}