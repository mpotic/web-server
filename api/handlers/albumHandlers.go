package handlers

import (
	//"web-server/services"
	"net/http"
	"web-server/api/router"

	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	router.Router.GET("/albums", func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, nil)
	})
}
