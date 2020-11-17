package main

import (
	"github.com/gorilla/mux"
	"github.com/stnnnghm/EventManager/handlers"
	"github.com/stnnnghm/EventManager/store"
	"log"
	"net/http"
)

// Args args used to run the server
type Args struct {
	// pg conn string
	// e.g "postgres://user:password@localhost:5433/database?sslmode=disable"
	conn string
	// port for the server eg ":8080"
	port string
}

// Run run the server based on args
func Run(args Args) error {
	// router
	router := mux.NewRouter().
		PathPrefix("/api/v1/").
		Subrouter()

	st := store.NewPostgresEventStore(args.conn)
	handler := handlers.NewEventHandler(st)
	RegisterAllRoutes(router, handler)

	// start server
	log.Println("Starting server at port: ", args.port)
	return http.ListenAndServe(args.port, router)
}

// RegisterAllRoutes registers all routes of the API
func RegisterAllRoutes(router *mux.Router, handler handlers.IEventHandler) {
	// set content type
	router.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			next.ServeHTTP(w, r)
		})
	})

	// get events
	router.HandleFunc("/event", handler.Get).Methods(http.MethodGet)
	// create events
	router.HandleFunc("/event", handler.Create).Methods(http.MethodPost)
	// delete events
	router.HandleFunc("/event", handler.Delete).Methods(http.MethodDelete)

	// cancel events
	router.HandleFunc("/event/cancel", handler.Cancel).Methods(http.MethodPatch)
	// update event details
	router.HandleFunc("/event/details", handler.UpdateDetails).Methods(http.MethodPut)
	// reschedule events
	router.HandleFunc("/event/reschedule", handler.Reschedule).Methods(http.MethodPatch)

	// list events
	router.HandleFunc("/event", handler.List).Methods(http.MethodGet)
}