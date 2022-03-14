package view

import (
	"net/http"

	"my-guora/conf"

	"github.com/gin-gonic/gin"
)

// Login func
func Login(c *gin.Context) {

	var csrdata map[string]interface{}

	template := "login.html"
	data := map[string]interface{}{
		"lang":    conf.Config().Lang,
		"csrdata": csrdata,
	}

	c.HTML(http.StatusOK, template, data)

	return
}
