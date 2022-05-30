package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/BreCkver/st-codeChallenge/routers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

/*Handler set the port and server listining*/
func Handler() {
	router := mux.NewRouter()
	router.HandleFunc("/", routers.Index).Methods("GET")
	router.HandleFunc("/", routers.LoadFile).Methods("POST")
	router.HandleFunc("/confirmation", routers.Confirmation).Methods("GET")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8089"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))

}
