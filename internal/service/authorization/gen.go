package authorization

import (
	"my-guora/internal/model"

	"net/url"
	"strings"

	"my-guora/conf"

	"github.com/dgrijalva/jwt-go/v4"
)

// Gen Service
func Gen(user model.User) (tokenString string, err error) {

	SecretString := conf.Config().SecretKey

	Secret := []byte(SecretString)

	claims := Claims{
		user.ID,
		user.Type,
		user.ProfileID,
		compatibleJSEncodeURIComponent(url.QueryEscape(user.Profile.Name)),
		compatibleJSEncodeURIComponent(url.QueryEscape(user.Profile.Desc)),
		jwt.StandardClaims{
			Audience:  []string{},
			ExpiresAt: &jwt.Time{},
			ID:        "",
			IssuedAt:  &jwt.Time{},
			Issuer:    "localhost",
			NotBefore: &jwt.Time{},
			Subject:   "",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString(Secret)

	return
}

func compatibleJSEncodeURIComponent(str string) (resultStr string) {
	resultStr = str
	resultStr = strings.Replace(resultStr, "+", "%20", -1)
	resultStr = strings.Replace(resultStr, "%21", "!", -1)
	resultStr = strings.Replace(resultStr, "%27", "'", -1)
	resultStr = strings.Replace(resultStr, "%28", "(", -1)
	resultStr = strings.Replace(resultStr, "%29", ")", -1)
	resultStr = strings.Replace(resultStr, "%2A", "*", -1)
	return
}
