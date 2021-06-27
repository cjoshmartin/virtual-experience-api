package main

import (
	"github.com/cjoshmartin/virtual-experience-api/webserver"
)

func main() {
	router := webserver.SetRoutes()
	router.Run()
}
