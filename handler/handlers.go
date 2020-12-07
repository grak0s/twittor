package handler

import (
	"log"
	"net/http"
	"os"

	"github.com/twittor/middlew"
	"github.com/twittor/routers"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

//Manejadores seteo de puerto y pone a escuchar el servidor
func Manejadores() {
	router := mux.NewRouter()

	router.HandleFunc("/registro", middlew.ChequeoDB(routers.Registro)).Methods("POST")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"

	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))

}
