package view

import (
	"net/http"

	"my-guora/conf"

	"github.com/gin-gonic/gin"
)

// Admin func
func Admin(c *gin.Context) {

	var csrdata map[string]interface{}

	template := "admin.html"
	data := map[string]interface{}{
		"lang":    conf.Config().Lang,
		"csrdata": csrdata,
	}

	c.HTML(http.StatusOK, template, data)

	return
}
