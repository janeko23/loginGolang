package usersdirectory

import (
	"fmt"
	"github.com/go-ldap/ldap/v3"
	
	"igualdad.mingeneros.gob.ar/pkg/services/log"

	"igualdad.mingeneros.gob.ar/pkg/services/usersDirectory/models"
)

var conn *ldap.Conn

// Login function
func Login(usuarie string, password string) error {

	log.Info("Inicia binding con servidor ldap")

	l, err := bind()
	if err != nil {
		err := manejarErrorConexion(err)
		log.Error(err.Error())
		return err
	}

	log.Info("Binding con servidor ldap finalizado correctamente")

	//TODO: seguridad: ver si tenemos que configurar este certificado
	// Reconnect with TLS
	/*
	err = l.StartTLS(&tls.Config{InsecureSkipVerify: true})
	if err != nil {
		log.Fatal(err)
	}
	*/
	
	bindAsReadOnlyUser()
	if err != nil {
		log.Error(err.Error()) // TODO: ver como manejar estos errores de usuarie de solo lectura
	}
	
	sr, err := buscarUsuarie(usuarie)
	if err != nil {
		return manejarErrorBinding(err)
	}

	if len(sr.Entries) == 0 {
		return ErrInexistentUser
	} else if len(sr.Entries) > 1 {
		log.Error("User does not exist or too many entries returned", log.String("user", usuarie))
	}

	userdn := sr.Entries[0].DN

	// Bind para verificar password
	err = l.Bind(userdn, password)
	if err != nil {
		return manejarErrorBinding(err)
	}

	return nil
}

// Logout function
func Logout() {
	defer conn.Close()
	conn = nil
}

// CrearUsuarie recibe los datos para crear un registro nuevo en el directorio ldap
func CrearUsuarie(usuarie *models.Usuarie) error {

	sr, err := buscarUsuarie(usuarie.UID)
	if err != nil {
		log.Error("No se pudo crear usuarie", log.String("error", err.Error()), log.String("user", usuarie.Email))
		return manejarErrorBinding(err) // TODO: ver si esta bien manejado este error
	}

	if len(sr.Entries) != 0 {
		log.Error("No se pudo crear usuarie debido a que ya existe", log.String("error", err.Error()), log.String("user", usuarie.Email))
		return ErrUserExists
	}

	addRequest := ldap.NewAddRequest("uid=" + usuarie.UID + ",ou=people,dc=example,dc=org", nil)
	addRequest.Attribute("objectClass", []string{"extensibleObject", "uidObject", "account", "userSecurityInformation", "top"})
	addRequest.Attribute("uid", []string{usuarie.UID})
	addRequest.Attribute("email", []string{usuarie.Email})
	addRequest.Attribute("member", []string{"cn=admin,dc=example,dc=com"})
	addRequest.Attribute("name", []string{usuarie.Nombre})
	addRequest.Attribute("sn", []string{usuarie.Apellido})
	addRequest.Attribute("userPassword",[]string{usuarie.Password})

	if conn == nil {
		log.Error("Ocurrio un error debido a que no se realizo login o se finalizo")
		return ErrLoginNeeded
	}

	err = conn.Add(addRequest)	
	if err != nil {
		log.Error("Ocurrio un error de conexion", log.String("error", err.Error()))
		return manejarErrorConexion(err)	
	} 
	
	return nil
}

func bind() (*ldap.Conn, error) {

	url := "ldap://ldap:389" // TODO: recibirlo por archivo config

	l, err := ldap.DialURL(url) // TODO: es el nombre del container. Recibir por archivo config
	if err != nil {		
		log.Error("Ocurrio un error al realizar binding con servidor ldap", log.String("error", err.Error()))
		return l, manejarErrorConexion(err)	
	}

	conn = l

	return l, nil
}

func bindAsReadOnlyUser() error {

	readOnlyUserAndPass := "admin" //TODO: hardcodeada, sacarla de archivo config hasheada
	
	if conn == nil {
		return ErrLoginNeeded
	}
	
	err := conn.Bind("cn="+readOnlyUserAndPass+",dc=example,dc=org", readOnlyUserAndPass)

	return err
}

func buscarUsuarie(usuarie string) (*ldap.SearchResult, error) {

	searchRequest := ldap.NewSearchRequest(
		"dc=example,dc=org",
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
		fmt.Sprintf("(&(objectClass=*)(|(cn=%s)(uid=%s)))", usuarie, usuarie),
		[]string{"dn"},
		nil,
	)

	if conn == nil {
		return nil, ErrLoginNeeded
	}

	return conn.Search(searchRequest)
}

func manejarErrorConexion(err error) error {
	
	errMsg := err

	if (ldap.IsErrorWithCode(err, ldap.ErrorNetwork)) {
		
		errMsg = ErrConexionServidor
	
	}
	
	return errMsg
}

func manejarErrorBinding(err error) error {
	
	errMsg := err

	if (ldap.IsErrorWithCode(err, ldap.LDAPResultInvalidCredentials)) {

		errMsg = ErrInvalidPassword

	} else if (ldap.IsErrorWithCode(err, ldap.LDAPResultInvalidDNSyntax)) {

		errMsg = ErrUserSyntax
	
	} else if (ldap.IsErrorWithCode(err, ldap.ErrorEmptyPassword)) {

		errMsg = ErrUserEmptyPassword
	} 

	return errMsg
}