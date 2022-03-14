package authorization

import (
	"errors"

	"my-guora/conf"

	"github.com/dgrijalva/jwt-go/v4"
)

// Parse Service
func Parse(tokenString string) (ID int, ProfileID int, err error) {

	SecretString := conf.Config().SecretKey
	Secret := []byte(SecretString)

	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return Secret, nil
	})

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		ID = claims.ID
		ProfileID = claims.ProfileID
	} else {
		err = errors.New("Token Not Valid")
	}
	return

}
