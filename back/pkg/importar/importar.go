package importar

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"igualdad.mingeneros.gob.ar/pkg/personas"
	"igualdad.mingeneros.gob.ar/pkg/utils"
	"igualdad.mingeneros.gob.ar/pkg/services/db/models"
	"gopkg.in/guregu/null.v3"
	"strconv"
	"math"
	"regexp"
	"strings"
	"time"
	"reflect"
	"github.com/go-playground/validator/v10"
)

//Mapea nombres de columnas a su índice en cada registro
var idx map[string]int = map[string]int{
	"Fuente" : 0,
	"OrganizacionCarga": 1,
	"Referente": 2,
	"CompletedTime": 3,
	"Apellido": 4,
	"Nombres": 5,
	"TipoDocumento": 6,
	"NroDNI": 7,
	"NroPasaporte": 8,
	"NroOtroDoc": 9,
	"PaisNacimiento": 10,
	"OtroPaisNacimiento": 11,
	"Nacionalidad":12,
	"OtraNacionalidad": 13,
	"IdentidadDeGenero": 14,
	"OtraIdentidadDeGenero": 15,
	"FechaNacimiento": 16,
	"Edad": 17,
	"NivelEducativo": 18,
	"CompletoNivelEducativo": 19,
	"Calle": 20,
	"Numero": 21,
	"Torre": 22,
	"Piso": 23,
	"Depto": 24,
	"Localidad": 25,
	"Partido": 26,
	"Provincia": 27,
	"AdicionalesDomicilio": 28,
	"Georeferenciacion":29,
	"Telefono":30,
	"Email":31,
	"PerteneceAOrganizacion": 32,
	"TipoOrganizacionPertenencia": 33,
	"OtroTipoOrganizacionPertenencia":34,
	"OrganizacionPertenencia": 35,
	"HaSidoPromotorx": 36,
	"Ministerio": 37,
	"TrabajoViolenciasPorRazonesDeGenero": 38 ,
	"TrabajoIgualdadDeGeneros": 39,
	"TrabajoSaludSexualReproductiva": 40,
	"TrabajoDiversidad": 41,
	"OtrasTrabajo": 42,
	"InteresViolenciasPorRazonesDeGenero": 43,
	"InteresIgualdadDeGeneros": 44,
	"InteresSaludSexualReproductiva": 45,
	"InteresDiversidad": 46,
	"OtrasInteres": 47,
	"ObtuvoRemuneracion": 48,
	"ObtuvoCertificacion": 49,
	"DeOrganismoDelEstadoYSociedadCivil": 50,
	"DeOtroOrganismo": 51,
	"OrganismoCertificacion1": 52,
	"OrganismoCertificacion2": 53,
	"OrganismoCertificacion3": 54,
	"OrganismoCertificacion4": 55,
	"OrganismoCertificacion5": 56,
	"OrganizacionFormacion": 57, 
	"CertificacionViolenciasPorRazonesDeGenero": 58,
	"CertificacionIgualdadDeGeneros": 59,
	"CertificacionSaludSexualReproductiva": 60,
	"CertificacionDiversidad": 61,
	"OtrasCertifico": 62,
}

var idTematica map[string]int = map[string]int{
	"ViolenciasPorRazonesDeGenero": 1,
	"IgualdadDeGeneros": 2,
	"SaludSexualReproductiva": 3,
	"Diversidad": 4,
}

const maxFilas int = 19000 // Constante hardcodeada que indica cota superior de filas del csv a cargar
//PathToBase1 const path al archivo base1.csv en el contenedor del back
const PathToBase1 = "/app/base1.csv"
const PathToBase2 = "/app/base2.csv"

//CargarCSV lee el csv denormalizado en pathToCsv e ingresa los registros de a 100
func CargarCSV(pathToCsv string){

	csvfile, err := os.Open(pathToCsv)
	if err != nil {
		fmt.Println("No se pudo abrir el csv", err)
	}

	r := csv.NewReader(csvfile)
	personasAInsertar := make([]personas.PersonaDenormalizada, 0)
	k := 0
	for i := 0; i < maxFilas; i++ {
		// Read each row from csv
		row, err := r.Read()
		if i == 0 { //La primera fila contiene nombres de columnas
			continue
		}
		if err == io.EOF { //Terminó el archivo
			break
		}
		if err != nil {
			fmt.Println(err.Error())
		}

		if k < 5 {
			var p *personas.PersonaDenormalizada = mapRow(row)
			personasAInsertar = append(personasAInsertar, *p)
			k += 1
		}

		//Cada 100 ingreso los registros
		if math.Mod(float64(i),100) == 0.0{
			_, err := personas.AltaPersonas(&personasAInsertar, models.TablePersonas, utils.FromImport)
			if err != nil {
				fmt.Println(err.Error())
			}
			personasAInsertar = make([]personas.PersonaDenormalizada,0)
		}
	}
}

func mapRow(row []string) (*personas.PersonaDenormalizada) {
	var p personas.PersonaDenormalizada

	
	p.Metadata.Fuente = row[idx["Fuente"]]
	p.Metadata.OrganizacionCarga = mapString(row[idx["OrganizacionCarga"]])
	p.Metadata.ReferenteCarga =  mapString(row[idx["Referente"]])
	p.Metadata.CompletedTime = mapFecha(row[idx["CompletedTime"]], "01/02/2006")

	p.DatosPersonales.Apellido = row[idx["Apellido"]]
	p.DatosPersonales.Nombres = row[idx["Nombres"]]
	p.DatosPersonales.TipoDocumento = mapTipoDoc(row[idx["TipoDocumento"]], row[idx["NroDNI"]], row[idx["NroPasaporte"]], row[idx["NroOtroDoc"]])
	p.DatosPersonales.NroDocumento = null.StringFrom(mapNroDoc(row[idx["NroDNI"]], row[idx["NroPasaporte"]], row[idx["NroOtroDoc"]]))
	p.DatosPersonales.IdentidadGenero = mapIdentidadDeGenero(row[idx["IdentidadDeGenero"]], row[idx["OtraIdentidadDeGenero"]])
	p.DatosPersonales.FechaNacimiento = mapFecha(row[idx["FechaNacimiento"]], "01/02/2006")
	p.DatosPersonales.Edad = mapEntero(row[idx["Edad"]])

	p.Educacion.NivelEducativo = mapString(row[idx["NivelEducativo"]])
	p.Educacion.CompletoNivelEducativo = mapString(row[idx["CompletoNivelEducativo"]])
	
	p.Contacto.Calle = mapString(row[idx["Calle"]])
	p.Contacto.Numero = mapString(row[idx["Numero"]])
	p.Contacto.Torre = mapString(row[idx["Torre"]])
	p.Contacto.Piso = mapString(row[idx["Piso"]])
	p.Contacto.Depto = mapString(row[idx["Depto"]])
	p.Contacto.Localidad = mapString(parseLocalidad(row[idx["Localidad"]]))
	p.Contacto.Partido = mapString(parseDepartamento(p.Contacto.Localidad, row[idx["Partido"]]))
	p.Contacto.Provincia = mapString(parseProvincia(row[idx["Provincia"]]))
	p.Contacto.PaisNacimiento = null.StringFrom(mapPais(row[idx["PaisNacimiento"]], row[idx["OtroPaisNacimiento"]]))
	p.Contacto.Nacionalidad = null.StringFrom(mapNacionalidad(row[idx["Nacionalidad"]], row[idx["OtraNacionalidad"]]))
	p.Contacto.Telefono = mapString(row[idx["Telefono"]])
	p.Contacto.Email = mapString(row[idx["Email"]])
	
	p.Experiencia.FuePromotorxMinisterio = mapBooleano(row[idx["HaSidoPromotorx"]])
	p.Experiencia.Ministerio = mapString(row[idx["Ministerio"]])
	p.Experiencia.ObtuvoRemuneracion = mapBooleano(row[idx["ObtuvoRemuneracion"]])
	p.Experiencia.ObtuvoCertificacion = mapBooleano(row[idx["ObtuvoCertificacion"]])
	p.Experiencia.DeOrganismoDelEstadoYSociedadCivil = mapString(row[idx["DeOrganismoDelEstadoYSociedadCivil"]])
	p.Experiencia.DeOtroOrganismo = mapString(row[idx["DeOtroOrganismo"]])
	p.Experiencia.OrganizacionFormacion = mapString(row[idx["OrganizacionFormacion"]])
	p.Experiencia.OrganismosCertificacion = mapOrganismos(row)

	p.Experiencia.TematicasTrabajo = mapExperienciasEnTematicas(row, "Trabajo")
	p.Experiencia.TematicasInteres = mapExperienciasEnTematicas(row, "Interes")
	p.Experiencia.TematicasCertifico = mapExperienciasEnTematicas(row, "Certificacion")

	p.Experiencia.OtrasTematicasTrabajo = mapString(row[idx["OtrasTrabajo"]])
	p.Experiencia.OtrasTematicasInteres = mapString(row[idx["OtrasInteres"]])
	p.Experiencia.OtrasTematicasCertifico = mapString(row[idx["OtrasCertifico"]])

	validateFields(&p)

	return &p
}

//Funciones de mapeo especificas para cada campo

func mapNroDoc(nroDNI, nroPasaporte, nroOtroDoc string) string {
	if len(nroDNI) > 0 {
		return nroDNI
	}
	if len(nroPasaporte) > 0 {
		return nroPasaporte
	}
	return nroOtroDoc
}

func mapTipoDoc(tipoDoc, nroDNI, nroPasaporte, otroNro string) string {
	if len(tipoDoc) == 0 {
		if len(nroDNI) > 0 {
			tipoDoc = "DNI"
		} else if len(nroPasaporte) > 0 {
			tipoDoc = "Pasaporte"
		} else if len(otroNro) > 0 {
			tipoDoc = "Otro"
		} else {
			tipoDoc = "Sin DNI"
		}
	}
	return tipoDoc
}

func mapIdentidadDeGenero(identidad, otraIdentidad string) string{
	if len(identidad) > 0 {
		return identidad
	}
	return otraIdentidad
}

func mapPais(pais, otroPais string) string{
	if len(pais) > 0 {
		return pais
	}
	return otroPais
}

func mapNacionalidad(nacionalidad, otraNacionalidad string) string{
	if len(nacionalidad) > 0 {
		return nacionalidad
	}
	return otraNacionalidad
}

func mapBooleano(valor string) null.Bool {
	resBooleano := false
	hayDato := true

	valor = strings.ToLower(valor)

	if regexSi := regexp.MustCompile("s(í|i)$"); regexSi.MatchString(valor) {
		resBooleano = true
	} else if valor != "no" {
		hayDato = false
	}

	return null.NewBool(resBooleano, hayDato)

}

func mapString(valor string) null.String {
	hayDato := len(valor) > 0
	return null.NewString(valor, hayDato)
}	

func mapEntero(valor string) null.Int {
	valorInt, err := strconv.ParseInt(valor, 10, 64)
	intValido := valor != "" && err == nil
	return null.NewInt(valorInt, intValido);
}

func mapFecha(fecha string, layout string) null.Time {
	elems := strings.Split(fecha,"/")
	for idx, elem := range(elems) {
		if len(elem) == 1 {
			elems[idx] = "0"+elem
		}
	}

	fechaCompleta := strings.Join(elems, "/")
	d, err := time.Parse(layout, fechaCompleta)
	validDate := fecha != "" && err == nil
	return null.NewTime(d, validDate)
}

func mapExperienciasEnTematicas(row []string, tipo string) []int {

	tematicasARegistrar := make([]int, 0)
	tematicasPosibles := []string{ "ViolenciasPorRazonesDeGenero", "IgualdadDeGeneros", "SaludSexualReproductiva", "Diversidad"}

	for _, tematica := range(tematicasPosibles) {
		if trabajoTematica := mapBooleano(row[idx[tipo+tematica]]); !trabajoTematica.IsZero() && trabajoTematica.ValueOrZero() {
			tematicasARegistrar = append(tematicasARegistrar, idTematica[tematica])
		}
	}
	
	return tematicasARegistrar
}

func mapOrganismos(row []string) []string {
	organismosQueCertificaron := make([]string, 0)
	
	validate := validator.New()
	campoNombreOrganismo, _ := reflect.TypeOf(models.Organismo{}).FieldByName("Nombre")
	validationTag := campoNombreOrganismo.Tag.Get("validate")
	
	for i := 1; i <= 5; i++{
		org := strings.ToLower(row[idx["OrganismoCertificacion"+fmt.Sprintf("%d",i)]])
		rSpaces := regexp.MustCompile("no|\\s|\\.|\\-")
		err := validate.Var(org, validationTag)
		if rSpaces.ReplaceAllString(org, "") != "" && err == nil {
			organismosQueCertificaron = append(organismosQueCertificaron, org);
		}
	}

	return organismosQueCertificaron
}