package personas

import (
	"errors"
	"gopkg.in/guregu/null.v3"
	"fmt"
	"gorm.io/gorm"

	"igualdad.mingeneros.gob.ar/pkg/services/db/models"
	"igualdad.mingeneros.gob.ar/pkg/services/db/services"
)

func mergearPersona(personaIncompleta, personaCompleta *models.Persona) ([]string, error) {
	
	columnasMergeadas := make([]string, 0)

	huboMergeReferente, err := mergearReferente(personaIncompleta, personaCompleta.Referente)
	if err != nil {
		fmt.Println("Ocurrio un error al mergear un referente, había un ", personaIncompleta.Referente, " a mergear por ", personaCompleta.Referente)
	}
	if huboMergeReferente {
		columnasMergeadas = append(columnasMergeadas, "Referente")
	}

	//Merge Fecha de carga
	if !personaCompleta.CompletedTime.IsZero() && !personaCompleta.CompletedTime.Equal(personaIncompleta.CompletedTime){
		personaIncompleta.CompletedTime = personaCompleta.CompletedTime
		columnasMergeadas = append(columnasMergeadas, "CompletedTime")
	}

	//Merge Telefono
	if !personaCompleta.Telefono.IsZero() && personaCompleta.Telefono != personaIncompleta.Telefono {
		personaIncompleta.Telefono = personaCompleta.Telefono
		columnasMergeadas = append(columnasMergeadas, "Telefono")
	}

	//Merge correo electronico
	if !personaCompleta.Email.IsZero() && personaCompleta.Email != personaIncompleta.Email {
		personaIncompleta.Email = personaCompleta.Email
		columnasMergeadas = append(columnasMergeadas, "Email")
	}

	//Merge si fue promotore
	if !personaCompleta.FuePromotorxMinisterio.IsZero() && personaCompleta.FuePromotorxMinisterio != personaIncompleta.FuePromotorxMinisterio {
		personaIncompleta.FuePromotorxMinisterio = personaCompleta.FuePromotorxMinisterio
		columnasMergeadas = append(columnasMergeadas, "FuePromotorxMinisterio")
	}

	//Merge ministerio
	if !personaCompleta.Ministerio.IsZero() && personaCompleta.Ministerio != personaIncompleta.Ministerio {
		personaIncompleta.Ministerio = personaCompleta.Ministerio
		columnasMergeadas = append(columnasMergeadas, "Ministerio")
	}

	//Merge otras tematicas formacion
	if !personaCompleta.OtrasTematicasFormo.IsZero() && personaCompleta.OtrasTematicasFormo != personaIncompleta.OtrasTematicasFormo {
		personaIncompleta.OtrasTematicasFormo = personaCompleta.OtrasTematicasFormo
		columnasMergeadas = append(columnasMergeadas, "OtrasTematicasFormo")
	}

	//Merge otras tematicas interes
	if !personaCompleta.OtrasTematicasInteres.IsZero() && personaCompleta.OtrasTematicasInteres != personaIncompleta.OtrasTematicasInteres {
		personaIncompleta.OtrasTematicasInteres = personaCompleta.OtrasTematicasInteres
		columnasMergeadas = append(columnasMergeadas, "OtrasTematicasInteres")
	}

	//Merge otras tematicas certificacion
	if !personaCompleta.OtrasTematicasCertifico.IsZero() && personaCompleta.OtrasTematicasCertifico != personaIncompleta.OtrasTematicasCertifico {
		personaIncompleta.OtrasTematicasCertifico = personaCompleta.OtrasTematicasCertifico
		columnasMergeadas = append(columnasMergeadas, "OtrasTematicasCertifico")
	}

	//Merge obtuvo remuneracion
	if !personaCompleta.ObtuvoRemuneracion.IsZero() && personaCompleta.ObtuvoRemuneracion != personaIncompleta.ObtuvoRemuneracion {
		personaIncompleta.ObtuvoRemuneracion = personaCompleta.ObtuvoRemuneracion
		columnasMergeadas = append(columnasMergeadas, "ObtuvoRemuneracion")
	}

	//Merge obtuvo certificacion
	if !personaCompleta.ObtuvoCertificacion.IsZero() && personaCompleta.ObtuvoCertificacion != personaIncompleta.ObtuvoCertificacion {
		personaIncompleta.ObtuvoCertificacion = personaCompleta.ObtuvoCertificacion
		columnasMergeadas = append(columnasMergeadas, "ObtuvoCertificacion")
	}

	//Merge de organismo del estado o sociedad civil
	if !personaCompleta.DeOrganismoDelEstadoYSociedadCivil.IsZero() && personaCompleta.DeOrganismoDelEstadoYSociedadCivil != personaIncompleta.DeOrganismoDelEstadoYSociedadCivil {
		personaIncompleta.DeOrganismoDelEstadoYSociedadCivil = personaCompleta.DeOrganismoDelEstadoYSociedadCivil
		columnasMergeadas = append(columnasMergeadas, "DeOrganismoDelEstadoYSociedadCivil")
	}

	//Merge de otro organismo
	if !personaCompleta.DeOtroOrganismo.IsZero() && personaCompleta.DeOtroOrganismo != personaIncompleta.DeOtroOrganismo {
		personaIncompleta.DeOtroOrganismo = personaCompleta.DeOtroOrganismo
		columnasMergeadas = append(columnasMergeadas, "DeOtroOrganismo")
	}

	//Merge Organizacion formacion
	if !personaCompleta.IDOrgaFormacion.IsZero() && personaCompleta.IDOrgaFormacion != personaIncompleta.IDOrgaFormacion {
		personaIncompleta.IDOrgaFormacion = personaCompleta.IDOrgaFormacion
		columnasMergeadas = append(columnasMergeadas, "OrganiazcionFormacion")
	}

	return columnasMergeadas, nil

}

func mergearReferente(p *models.Persona, referente null.String) (bool, error) {

	huboMerge := false
	if !p.Referente.IsZero() && p.Referente != referente {
		return false, errors.New("El referente de la persona " + fmt.Sprint(p.ID) + " es otrx")
	} else if referente != p.Referente {
		p.Referente = referente
		huboMerge = true
	}

	return huboMerge ,nil

}
func mergearTematicas(tx *gorm.DB, id int, experiencia *ExperienciaD) error {

	registrosTematicas := [][]int{experiencia.TematicasTrabajo, experiencia.TematicasInteres, experiencia.TematicasCertifico}
	tablasRegistros := []string{ models.TableExperienciasEnTematicas, models.TableInteresesEnTematicas, models.TableCertificacionesEnTematicas}
	
	for idx, tabla := range(tablasRegistros){
		registraTematicas, err := services.TieneRegistros(tx, tabla, id)
		if err != nil {
			fmt.Println(err.Error())
		}
		if !registraTematicas{
			InsertarExperienciasEnTematicas(id, tabla, &registrosTematicas[idx], tx)
		}
	}

	return nil
}

func mergearCertificaciones(tx *gorm.DB, id int, experiencia *ExperienciaD) error {

	registrosOrganismos := experiencia.OrganismosCertificacion
	

	registraOrganismos, err := services.TieneRegistros(tx, models.TableCertificacionesOrganismos, id)
	if err != nil {
		fmt.Println(err.Error())
	}
	if !registraOrganismos{
		InsertarCertificacionesPorOrganismos(id, &registrosOrganismos, tx)
	}

	return nil
}
