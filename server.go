package main

import (
	"web-server/api/handlers"
	"web-server/api/router"
)

func main() {
	handlers.HandleRequests()
	router.RunRouter()
}
