package auth

import (
	"encoding/json"
	"net/http"
	"fmt"
	"os"
	"igualdad.mingeneros.gob.ar/pkg/auth"
	"igualdad.mingeneros.gob.ar/pkg/authapi"
	jwt "github.com/dgrijalva/jwt-go"
	"strings"
	"errors"
)

//AccessDetails detalles de acceso
type AccessDetails struct {
    UserID   string
}


//LoginApi loguea para la api
func LoginApi(w http.ResponseWriter, r *http.Request){
	
	var input auth.InputLogin

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		sendErr(w, http.StatusInternalServerError, err.Error())
	}
	response := loginApi(&input)

	w.WriteHeader(http.StatusOK)

	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		sendErr(w, http.StatusInternalServerError, err.Error())
	}
}


// Login se autentica en el directorio de ldap
func Login(w http.ResponseWriter, r *http.Request) {
	
	var input auth.InputLogin

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		sendErr(w, http.StatusInternalServerError, err.Error())
	}


	response := login(&input)

	w.WriteHeader(http.StatusOK)

	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		sendErr(w, http.StatusInternalServerError, err.Error())
	}
}

//VerifyToken verifica el token
func VerifyToken(r *http.Request) (*jwt.Token, error) {
	tokenString := extractToken(r)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
	   //Make sure that the token method conform to "SigningMethodHMAC"
	   if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		  return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
	   }
	   return []byte(os.Getenv("ACCESS_SECRET")), nil
	})
	if err != nil {
	   return nil, err
	}
	return token, nil
  }


//tokenValid verifica que sea valido
func tokenValid(token *jwt.Token, userToMatch string) error {

	claims, ok := token.Claims.(jwt.Claims)
	userDetails, _ := extractTokenMetadata(token)
	userID := userDetails.UserID
	if userID != userToMatch {
		return errors.New("El usuario no coincide")
	}
	fmt.Println("claims: ", claims)
	if !ok && !token.Valid {
	   return errors.New("El token no es válido")
	}
	return nil
  }

  func extractTokenMetadata(token *jwt.Token) (*AccessDetails, error) {
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
	   userId, ok := claims["user_id"].(string)
	   if !ok {
		  return nil, errors.New("Error when extracting user id")
	   }
	   return &AccessDetails{
		  UserID:   userId,
	   }, nil
	}
	return nil, nil
  }

func login(input *auth.InputLogin) map[string]interface{} {
	
	var response map[string]interface{}

	err := input.Login()
	if err != nil {
		response = map[string]interface{}{
			"status": "ERROR",
			"error":  err.Error(),
		}
	} else {
		response = map[string]interface{}{
			"status": "OK",
			"data":   "Login OK!",
		}
	}
	
	return response
}

func loginApi(input *auth.InputLogin) map[string]interface{} {
	
	var response map[string]interface{}

	tok, err := authapi.Login(input)
	if err != nil {
		response = map[string]interface{}{
			"status": "ERROR",
			"error":  err.Error(),
		}
	} else {
		response = map[string]interface{}{
			"status": "OK",
			"data":   tok,
		}
	}
	
	return response
}

func sendErr(w http.ResponseWriter, code int, message string) {
	resp, _ := json.Marshal(map[string]string{"error": message})
	http.Error(w, string(resp), code)
}

//extractToken extrae el token
func extractToken(r *http.Request) string {
	bearToken := r.Header.Get("Authorization")
	//normally Authorization the_token_xxx
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
	   return strArr[1]
	}
	return ""
  }

