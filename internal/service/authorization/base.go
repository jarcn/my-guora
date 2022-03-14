package authorization

import (
	"github.com/dgrijalva/jwt-go"
)

// Claims struct
type Claims struct {
	ID        int    `json:"id"`
	Type      int    `json:"type"`
	ProfileID int    `json:"profileID"`
	Name      string `json:"profile.name"`
	Desc      string `json:"profile.desc"`
	jwt.StandardClaims
}
