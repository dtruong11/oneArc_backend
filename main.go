package main

import (
	"fmt"
	"net/http"
	"os" // to read, create, write to file

	"github.com/gorilla/handlers" // works with cors
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/oneArc_backend/controller"
	"github.com/oneArc_backend/model"
	"github.com/oneArc_backend/routing"

	// "github.com/onearc_backend/controller"
	"github.com/rs/cors"
	"github.com/urfave/negroni" // middleware focused library, work with net/http. Bring your own router
)

func getPort() string {
	p := os.Getenv("PORT")
	if p != "" {
		return ":" + p
	}
	return ":8080"
}

func main() {
	fmt.Println("Go my backend")

	model.TestFunction()

	router := mux.NewRouter().StrictSlash(true)     // Strictslash redirects /path/ to /path
	router.HandleFunc("/", routing.HomePageHandler) // second argument: handler for home page

	// handle user registration, login
	us := router.PathPrefix("/auth").Subrouter()
	// us.HandleFunc("/signup", controller.CreatePerson).Methods("POST")
	// us.HandleFunc("/login", controller.GetPerson).Methods("GET")
	us.HandleFunc("/", controller.CreatePerson).Methods("POST")

	// apply cors middleware
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // allow all origins, using wildcard *
		AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"},
		AllowedHeaders:   []string{"X-Requested-With", "Content-Type", "Authorization", "Accept-Encoding", "Accept-Language"},
		AllowCredentials: true,
	})

	muxRouter := http.NewServeMux() // creates a new ServeMux
	muxRouter.Handle("/", router)   // registers the handler

	// Negroni binds cors & handler
	n := negroni.Classic()  // a new Negroni instance with the default middleware (Recovery, Logger, Static)
	n.Use(c)                // use a middleware with Use func. Bind the CORS middleware in the web service
	n.UseHandler(muxRouter) // adds a http.Handler onto the middleware stack. Invoke in the order added to negroni

	// listen on port & recover from panic
	http.ListenAndServe(getPort(), handlers.RecoveryHandler()(n)) // the second arg recovers from a panic, logs the panic, writes http.StatusInteranalServerError, continues to the next handler
}
