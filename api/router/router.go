package router

import(
	"github.com/gin-gonic/gin"
)

var Router *gin.Engine = gin.Default()

func RunRouter() {
	Router.Run(":8080")
}