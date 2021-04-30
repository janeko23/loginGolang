package authapi

import(
	jwt "github.com/dgrijalva/jwt-go"
	"time"
	"os"
	"errors"
	"igualdad.mingeneros.gob.ar/pkg/auth"
	"fmt"
)

//Login Loguea un usuario de la api
func Login(u *auth.InputLogin) (string, error) {
	//compare the user from the request, with the one we defined:
	userid := u.User
	pass := u.Password
	fmt.Println("pass: ", pass)
	if pass != "unapass" {
	   return "", errors.New("User or pass is incorrect")
	}
	token, err := createToken(userid)
	if err != nil {
	   return "", err
	}
	return token, nil
}

func createToken(userid string) (string, error) {
	var err error

	//Creating Access Token
	os.Setenv("ACCESS_SECRET", "jdnfksdmfksd") //this should be in an env file
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_id"] = userid
	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return "", err
	}
	return token, nil
}
