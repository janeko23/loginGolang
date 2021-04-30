package importar

import (
	"regexp"
	"strings"
)

func parseProvincia(provincia string) string{
	provincia = eliminarPuntuacion(provincia)

	provincia = parsePifies(provincia)

	provincia = parseCABA(provincia)

	provincia = parseBuenosAires(provincia)

	return provincia
}


func parseBuenosAires(text string) string {
	const provinciaBsAs ="BUENOS AIRES"
	text = strings.ToUpper(text)

	if text == "CIUDAD AUTONOMA DE BUENOS AIRES"{
		return text
	}

	if text =="QUILLMES" {
		text = provinciaBsAs
	}
	if regexPifio := regexp.MustCompile("BUENS?OA?S? I?AI?RES"); regexPifio.MatchString(text){
		text = provinciaBsAs
	}
	regexBsAs := regexp.MustCompile("B(UENO)?S\\s?A(IRE)?S");
	if regexBsAs.MatchString(text){
		text = provinciaBsAs
	}
	return text
}

func parseCABA(provincia string) string {
	const provinciaCABA ="CIUDAD AUTONOMA DE BUENOS AIRES"
	if provincia =="CAB" {
		provincia = provinciaCABA
	}
	if regexCABA := regexp.MustCompile("(CABA|(C A B A|CIUDAD.*BUENOS AIRES))"); regexCABA.MatchString(provincia){
		provincia = provinciaCABA
	}
	return provincia
}


func parsePifies(provincia string) string {

	if provincia =="MEDNOZA" {
		provincia ="MENDOZA"
	}

	if provincia =="CORRIENTES CAPITAL"{
		provincia ="CORRIENTES"
	}

	if provincia =="CATAMATCA" {
		provincia ="CATAMARCA"
	}

	if regNeu:= regexp.MustCompile("NEUQU(É|E)N"); regNeu.MatchString(provincia){
		provincia ="NEUQUEN"
	}

	if provincia =="CORODBA" || strings.Contains(provincia, "CBA") {
		provincia ="CORDOBA"
	}

	if strings.Contains(provincia,"JUJUY") {
		provincia ="JUJUY"
	}

	if provincia =="SANTAFE" || strings.Contains(provincia,"ALMAFUERTE") {
		provincia ="SANTA FE"
	}

	if strings.Contains(provincia, "DEL ESTERO") {
		provincia ="SANTIAGO DEL ESTERO"
	}

	if provincia == "COMODORO RIVADAVIA"{
		provincia = "CHUBUT"
	}

	if strings.Contains(provincia, "SAN LUIS") {
		provincia = "SAN LUIS"
	}

	if strings.Contains(provincia, "SAN CAYETANO") {
		provincia = "MISIONES"
	}




	return provincia
}