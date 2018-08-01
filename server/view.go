package server

import (
	"fmt"
	"log"
	"net/http"
)

//ServeView on the given port
func ServeView(port string) {
	fmt.Printf("Serving view on port %s...\n", port)
	log.Fatal(http.ListenAndServe(":"+port, http.FileServer(http.Dir("./frontend/build/"))))
}
