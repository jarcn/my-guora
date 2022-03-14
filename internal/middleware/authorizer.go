package middleware

import (
	"log"
	"my-guora/internal/constant"
	"my-guora/internal/service/authorization"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Authorizer middleware
func Authorizer() gin.HandlerFunc {
	return func(c *gin.Context) {

		// get the ss(signed string) inside cookie
		ss, err := c.Cookie(constant.SSKEY)
		if err != nil {
			log.Print("Authorizer error: ", err)
			c.Redirect(http.StatusTemporaryRedirect, "/login")
			c.Abort()
			return
		}

		// parse a ss(signed string)
		ID, ProfileID, err := authorization.Parse(ss)
		if err != nil {
			c.Redirect(http.StatusTemporaryRedirect, "/login")
			c.Abort()
			return
		}

		c.Set("uid", ID)
		c.Set("pid", ProfileID)

		// before request
		c.Next()
		// after request

		log.Print("UID: ", ID)
		log.Print("PID: ", ProfileID)
	}
}
