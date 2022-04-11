package app

import (
	"log"
	"net/http"
)

func RunApp() {

	r := routes()
	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(":8080", r))
}
