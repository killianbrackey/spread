package routers

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"log"
	"fmt"
	//"github.com/killianbrackey/authentication"
	//"bytes"
	//"encoding/base64"
	//"strings"

)

// InitRoutes : Creates all of the routes for our application and returns a router
func InitRoutes() *httprouter.Router {
	// Our Main Router - Uses httprouter as it has the fastest Benchmarks
	router := httprouter.New()

	//Set the routes for our application
	//router = SetAuthenticationRoutes(router)
	//router = authentication.
	return router
}

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    fmt.Fprint(w, "Welcome!\n")
}

func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "Hello, %s\n", ps.ByName("name"))
}

func main() {
	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/hello/:name", Hello)
	log.Fatal(http.ListenAndServe(":8080", router))
}
