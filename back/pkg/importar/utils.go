package importar

import (
	"regexp"
	"strings"
)

//EliminarPuntuacion elimina la puntuacion
func eliminarPuntuacion(text string) string {
	text = strings.ToUpper(text)
	text = reemplazarTildes(text)
	text = strings.ReplaceAll(text, ". ", " ")
	text = strings.ReplaceAll(text, ".", " ")

	regexpEspacios := regexp.MustCompile("\\s+")

	text = regexpEspacios.ReplaceAllString(text, " ")
	text = strings.Trim(text, " ")
	return text
}

func reemplazarTildes(text string) string {
	text = strings.ToUpper(text)
	regexTildeA := regexp.MustCompile("Á")
	regexTildeE := regexp.MustCompile("É")
 	regexTildeI := regexp.MustCompile("Í")
	regexTildeO := regexp.MustCompile("Ó")
	regexTildeU := regexp.MustCompile("Ú")

	text = regexTildeA.ReplaceAllString(text,"A")
	text = regexTildeE.ReplaceAllString(text,"E")
	text = regexTildeI.ReplaceAllString(text,"I")
	text = regexTildeO.ReplaceAllString(text,"O")
	text = regexTildeU.ReplaceAllString(text,"U")

	return text
}


func normalizarAbreviaturas(texto string) string {

	regexAlmirante := regexp.MustCompile("ALTE\\s")
	texto = regexAlmirante.ReplaceAllString(texto, "ALMIRANTE ")

	regexCoronel := regexp.MustCompile("CNEL\\s")
	texto = regexCoronel.ReplaceAllString(texto, "CORONEL")

	regexIngeniero := regexp.MustCompile("ING\\s")
	texto = regexIngeniero.ReplaceAllString(texto, "INGENIERO ")

	regexGeneral := regexp.MustCompile("GRAL?\\s")
	texto = regexGeneral.ReplaceAllString(texto, "GENERAL ")

	regexGobernador := regexp.MustCompile("GDOR\\s")
	texto = regexGobernador.ReplaceAllString(texto, "GOBERNADOR ")

	regexPresidencia := regexp.MustCompile("PCIA\\s")
	texto = regexPresidencia.ReplaceAllString(texto, "PRESIDENCIA ")

	regexPresidente := regexp.MustCompile("PTE\\s")
	texto = regexPresidente.ReplaceAllString(texto, "PRESIDENTE ")

	regexSanta := regexp.MustCompile("STA\\s")
	texto = regexSanta.ReplaceAllString(texto, "SANTA ")

	regexLibertador := regexp.MustCompile("LDOR\\s")
	texto = regexLibertador.ReplaceAllString(texto, "LIBERTADOR ")

	regexComandante := regexp.MustCompile("CMTE\\s")
	texto = regexComandante.ReplaceAllString(texto, "COMANDANTE ")

	regexDoctorManuel := regexp.MustCompile("D(OCTO)?R\\sM(\\.|ANUEL)\\sBELGRANO")
	texto = regexDoctorManuel.ReplaceAllString(texto, "DR. MANUEL BELGRANO")

	regexDoctorRicardo := regexp.MustCompile("DR\\sRICARDO")
	texto = regexDoctorRicardo.ReplaceAllString(texto, "DOCTOR RICARDO")

	regexPuerto := regexp.MustCompile("PTO\\s")
	texto = regexPuerto.ReplaceAllString(texto, "PUERTO ")

	regexAlem := regexp.MustCompile("N\\sALEM")
	texto = regexAlem.ReplaceAllString(texto, "N. ALEM")

	regexCarlosPaz := regexp.MustCompile("C\\sPAZ")
	texto = regexCarlosPaz.ReplaceAllString(texto, "C. PAZ")

	regexSolano := regexp.MustCompile("F\\sSOLANO")
	texto = regexSolano.ReplaceAllString(texto, "FRANCISCO SOLANO")

	regexCatan := regexp.MustCompile("G\\sCATAN")
	texto = regexCatan.ReplaceAllString(texto, "GONZALEZ CATAN")

	regexRincon := regexp.MustCompile("S\\sJ.*RINCON")
	texto = regexRincon.ReplaceAllString(texto, "SAN JOSE DEL RINCON")

	regexDeGiles := regexp.MustCompile("SAN\\sA.*GILES")
	texto = regexDeGiles.ReplaceAllString(texto, "SAN ANDRES DE GILES")

	regexEcheverria := regexp.MustCompile("E.*ECHEVERRIA")
	texto = regexEcheverria.ReplaceAllString(texto, "ESTEBAN ECHEVERRIA")

	regexAreco := regexp.MustCompile("S\\sA\\sDE\\sARECO")
	texto = regexAreco.ReplaceAllString(texto, "SAN ANTONIO DE ARECO")

	return texto
}
