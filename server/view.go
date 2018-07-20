package server

import (
	"fmt"
	"log"
	"net/http"
)

func ServeView(port string) {
	fmt.Printf("Listening on port %s...\n", port)
	log.Fatal(http.ListenAndServe(":"+port, http.FileServer(http.Dir("./dist/view/"))))
}
