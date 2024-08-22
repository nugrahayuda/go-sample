package router

import (
	"fmt"
	"io"
	"net/http"

	"github.com/gorilla/mux"
)

func Init() {
	// create a new router
	r := NewRouter()

	r.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Not Found", http.StatusNotFound)
	})

	r.HandleFunc("/", rootHandler)

	r.AddRoute(Route{
		Path: "/post",
		Handler: func(w http.ResponseWriter, r *http.Request) {
			// Read the request body
			body, err := io.ReadAll(r.Body)
			if err != nil {
				http.Error(w, "Failed to read request body", http.StatusBadRequest)
				return
			}

			// Process the request body
			// TODO: Add your logic here

			// Send a response
			fmt.Fprintln(w, "This is a POST request!")
			fmt.Fprintln(w, string(body))
		},
		Method: "POST",
	})

	//add get route
	r.AddRoute(Route{
		Path: "/get",
		Handler: func(w http.ResponseWriter, r *http.Request) {
			// Process the request body
			// TODO: Add your logic here

			// Send a response
			fmt.Fprintln(w, "This is a GET request!")
		},
		Method: "GET",
	})

	err := http.ListenAndServe(":8000", r)
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
