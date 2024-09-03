package handler

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

var port = "8000"

func (s *Service) Init() {
	// create a new router
	r := NewRouter()

	r.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Not Found", http.StatusNotFound)
	})

	r.HandleFunc("/", rootHandler)

	r.AddRoute(Route{
		Path:    "/get/{id}",
		Handler: s.GetUserByID,
		Method:  "GET",
	})

	fmt.Printf(`
##     ##  ##     ##  ######     ####   
 ##   ##   ##     ##  ##   ##   ##  ##  
  ## ##    ##     ##  ##    ## ##    ## 
   ###     ##     ##  ##    ## ######## 
   ###     ##     ##  ##    ## ##    ## 
   ###      ##   ##   ##   ##  ##    ## 
   ###       #####    ######   ##    ## 
`)
	fmt.Printf("Starting service at http://localhost:%s ......", port)
	err := http.ListenAndServe(":"+port, r)
	if err != nil {
		fmt.Printf("Failed to start server: %v\n", err)
	}
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Test BE Service API"))
}

type Router struct {
	*mux.Router
}

// create a new router
func NewRouter() *Router {
	router := mux.NewRouter()
	return &Router{
		Router: router,
	}
}

// define the Route type
type Route struct {
	Path    string
	Handler http.HandlerFunc
	Method  string
}

// add a new route to the router
func (r *Router) AddRoute(route Route) {
	r.HandleFunc(route.Path, route.Handler).Methods(route.Method)
}

// add a new subrouter to the router
func (r *Router) AddSubrouter(path string, subrouter *Router) {
	r.PathPrefix(path).Handler(subrouter)
}
