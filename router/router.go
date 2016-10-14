package routers

import (
	"github.com/julienschmidt/httprouter"
)

// InitRoutes : Creates all of the routes for our application and returns a router
func InitRoutes() *httprouter.Router {
	// Our Main Router - Uses httprouter as it has the fastest Benchmarks
	router := httprouter.New()

	//Set the routes for our application
	//router = SetAuthenticationRoutes(router)
	return router
}
