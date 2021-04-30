package personas

import (
	"time"
	"gorm.io/gorm"
	"gopkg.in/guregu/null.v3"
	"fmt"
	"igualdad.mingeneros.gob.ar/pkg/services/log"

	"igualdad.mingeneros.gob.ar/pkg/services/db"
	"igualdad.mingeneros.gob.ar/pkg/services/db/models"
	"igualdad.mingeneros.gob.ar/pkg/services/db/services"

)

//ListarPersonasDatosResumidos Obtiene datos completos de las personas segun filtro.
//Si el filtro es vacio devuelve el listado completo.
func ListarPersonasDatosResumidos(filter string, clasificacion string) (*[]models.ResumenPersona, error) {
	
	conn := db.GetDB()
	ResuModel, err := services.ListarPersonasDatosResumidos(conn, clasificacion, filter)
	
	if err != nil {
		log.Error("Ocurrio un error al listar personas", log.String("error", err.Error()))
		return nil, err
	}

	log.Debug("Se listaron las personas", log.Int("cantidadDePersonas", len(*ResuModel)))
	return ResuModel, nil
}


//BuscarPorID Busca personas que fueron seleccionadas
func BuscarPorID(ID int) (*models.Persona, error) {
	
	conn := db.GetDB()
	personaID, err := services.BuscarPersonaPorID(conn, models.TablePersonas, ID)
	if (err != nil) {
		log.Error("Ocurrio un error al buscar persona", log.Int("personaID", ID), log.String("error", err.Error()))
	} else {
		log.Debug("Busqueda de persona finalizada correctamente", log.Int("personaID", ID))
	}
	
	return personaID, err
}

//ExistePromotorxConDoc indica si existe una persona con el tipo y número de documento pasado por parámetro
func ExistePromotorxConDoc(nombreTabla string, tipoDoc int, nroDoc string) (bool, error) {
	conn := db.GetDB()
	nroDocParsed := null.StringFrom(nroDoc)
	personaEncontrada, err := services.BuscarPersonaPorTipoYDocumento(conn, nombreTabla, tipoDoc, nroDocParsed)
	if err != nil {
		return false, err
	} else {
		return personaEncontrada != nil, nil
	}
}

//AltaPersona Crea nuevo registro de persona completo 
func AltaPersona(personaNueva *Persona, tabla string, fuente string) (int, error) {
	
	conn := db.GetDB()
	tx := conn.Begin()
	fmt.Println("persona entrante: ", personaNueva)
	personaNuevaModel, _ := getPersonaModel(personaNueva, fuente)

	tipoDoc := personaNuevaModel.IDTipoDocumento;
	nroDoc := personaNuevaModel.NroDocumento;

	// pendienteConMismoDoc, _:= services.BuscarPersonaPorTipoYDocumento(tx, models.TablePendientes, tipoDoc, nroDoc)
	// if pendienteConMismoDoc != nil {
	// 	return 0, errors.New("El/la promotorx ya existe en el registro de pendientes")
	// }
	personaConMismoDoc, _:= services.BuscarPersonaPorTipoYDocumento(tx, models.TablePersonas, tipoDoc, nroDoc)
	var id int

	if personaConMismoDoc != nil {

		colsActualizadas, err := mergearPersona(personaConMismoDoc, personaNuevaModel)
		if err != nil {
			return personaConMismoDoc.ID, err
		}
		if len(colsActualizadas) > 0 {
			idPersona, err := services.ActualizarColumnas(tx, personaConMismoDoc, personaNuevaModel, colsActualizadas)
			if (err != nil) {
				tx.Rollback()
				log.Error("Ocurrio un error al mergear la persona", log.String("error", err.Error()))
				return idPersona, err
			}
			id = idPersona
		} else {
			id = personaConMismoDoc.ID
		}

		//TODO: Chequear si las experiencias ya habían sido cargadas
		err = InsertarExperienciasEnTematicas(personaConMismoDoc.ID, "experiencias_tematicas", &personaNueva.Experiencia.TematicasTrabajo, tx)
		if err != nil {
			tx.Rollback()
			return id, err
		}

	} else {
	
		idPersona, err := services.InsertarNuevaPersona(tx, *personaNuevaModel, tabla)

		if (err != nil) {
			tx.Rollback()
			log.Error("Ocurrio un error al crear nueva persona", log.String("error", err.Error()))
			return idPersona, err
		}
			
		log.Debug("Se creo un nuevo registro", log.Int("personaID", idPersona))

		tablaTematicas := tabla+"_tematicas"

		err = InsertarExperienciasEnTematicas(idPersona, tablaTematicas, &personaNueva.Experiencia.TematicasTrabajo, tx)
		if err != nil {
			tx.Rollback()
			return idPersona, err
		}
		id = idPersona
	}
	
	tx.Commit()
	return id, nil
}

//AltaPersonas Crea nuevos registros de personas 
func AltaPersonas(personas *[]PersonaDenormalizada, tabla string, fuente string) ([]int, error) {
	
	conn := db.GetDB()
	tx := conn.Begin()
	var personasMapeadas []models.Persona
	for _, personaNueva := range *personas {

		personaNuevaModel, err := getPersonaModelInputDenormalizado(&personaNueva)
		if err != nil{
			fmt.Println("no se pudo registar persona:\n", err.Error())
			return []int{}, err
		}

		tipoDoc := personaNuevaModel.IDTipoDocumento;
		nroDoc := personaNuevaModel.NroDocumento;

		personaConMismoDoc, _:= services.BuscarPersonaPorTipoYDocumento(tx, models.TablePersonas, tipoDoc, nroDoc)
		if personaConMismoDoc != nil {
			fmt.Println("El/la promotorx ", personaNuevaModel.NroDocumento, " ya existe en el registro")
			colsActualizadas, err := mergearPersona(personaConMismoDoc, personaNuevaModel)
			if err != nil {
				fmt.Println("Ocurrio un error al mergear la persona", err.Error())
			}
			if len(colsActualizadas) > 0 {
				_, err := services.ActualizarColumnas(tx, personaConMismoDoc, personaNuevaModel, colsActualizadas)
				if (err != nil) {
					fmt.Println("Ocurrio un error al mergear la persona", err.Error())
				}
			}

			mergearTematicas(tx, personaConMismoDoc.ID, &personaNueva.Experiencia)

			mergearCertificaciones(tx, personaConMismoDoc.ID, &personaNueva.Experiencia)

		} else {		
			personasMapeadas = append(personasMapeadas, *personaNuevaModel)
		}
		time.Sleep(time.Second)
	}
	if len(personasMapeadas) == 0 {
		tx.Commit()
		return []int{}, nil

	}
	ids, err := services.InsertarNuevasPersonas(conn, &personasMapeadas, models.TablePersonas)
	if err != nil {
		tx.Rollback()
		return []int{}, err
	}
	tx.Commit()
	return ids, nil
}

// InsertarExperienciasEnTematicas Inserta las experiencias laborales de una persona
func InsertarExperienciasEnTematicas(idPersona int, nombreTabla string, idsTematicas *[]int, tx *gorm.DB) error {
	if len(*idsTematicas) == 0 {
		return nil
	}
	experienciasEnTematicas := mapExperienciasEnTematicas(idPersona, idsTematicas)
	
	err := services.InsertarExpEnPersonasTematicas(tx, nombreTabla, experienciasEnTematicas)

	return err
}

// InsertarCertificacionesPorOrganismos funcion
func InsertarCertificacionesPorOrganismos(idPersona int, organismos *[]string, tx *gorm.DB) error {
	if len(*organismos) == 0 {
		return nil
	}
	certificacionesEnOrganismos := mapCertificaciones(idPersona, organismos)
	
	err := services.InsertarCertificacionesPorOrganismos(tx, certificacionesEnOrganismos)

	return err
}

//AprobarPersona aprueba una persona quitandola de la tabla pendientes y agregandola a personas
func AprobarPersona(idPersona int) (int, error) {

	/*conn := db.GetDB()
	tx := conn.Begin()

	personaPendiente, errBusqueda:= services.BuscarPersonaPorID(tx, models.TablePendientes, idPersona)
	if errBusqueda != nil {
		return 0, errBusqueda
	}
	tematicasPendientes, _ := services.BuscarTematicasDePersona(tx, models.TablePendientesTematicas, idPersona)
	
	idNuevo, err := services.InsertarNuevaPersona(tx, *personaPendiente, models.TablePersonas)

	if (err != nil) {
		tx.Rollback()
		log.Error("Ocurrio un error al crear nueva persona", log.Int("personaID", idPersona), log.String("error", err.Error()))
		return idNuevo, err
	}		
	log.Debug("Se ingresó la persona a Tabla personas", log.Int("personaID", idPersona))

	if len(tematicasPendientes) != 0 { 
		err = InsertarExperienciasEnTematicas(idPersona, idNuevo, models.TableExperienciasEnTematicas, &tematicasPendientes, tx)
		if err != nil {
			tx.Rollback()
			return idNuevo, err
		}
		log.Debug("Se registraron las tematicas en personas_tematicas")

		err = services.EliminarExperienciasEnTematicas(tx, models.TablePendientesTematicas, idPersona)
		if err != nil {
			tx.Rollback()
			return idNuevo, err
		}
	}
	log.Debug("Se eiminaron las tematicas en pendientes_tematicas")

	err = services.EliminarPersona(tx, models.TablePendientes, idPersona)
	if err != nil {
		tx.Rollback()
		return idNuevo, err
	}
	log.Debug("Se eiminó la persona de pendientes")

	
	tx.Commit()
	return idNuevo, nil*/

	return 0, nil

}

//RechazarPersona aprueba una persona quitandola de la tabla pendientes y agregandola a personas
func RechazarPersona(idPersona int) error {

	conn := db.GetDB()
	tx := conn.Begin()

	_, errBusqueda:= services.BuscarPersonaPorID(tx, models.TablePendientes, idPersona)
	if errBusqueda != nil {
		return errBusqueda
	}
	tematicasPendientes, _ := services.BuscarTematicasDePersona(tx, models.TablePendientesTematicas, idPersona)

	if len(tematicasPendientes) != 0 {
		err := services.EliminarExperienciasEnTematicas(tx, models.TablePendientesTematicas, idPersona)
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	log.Debug("Se eiminaron las tematicas en pendientes_tematicas")

	err := services.EliminarPersona(tx, models.TablePendientes, idPersona)
	if err != nil {
		tx.Rollback()
		return err
	}
	log.Debug("Se eiminó la persona de pendientes")

	
	tx.Commit()
	return nil

}