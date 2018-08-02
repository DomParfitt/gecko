package server

import (
	"log"
	"net/http"
)

//ShutdownListener struct
type ShutdownListener struct {
	servers []http.Server
}

//New ShutdownListener
func New(servers []http.Server) *ShutdownListener {
	return &ShutdownListener{
		servers: servers,
	}
}

//Listen for a request to shutdown the servers
func (s *ShutdownListener) Listen(port string, servers []http.Server) {
	http.HandleFunc("/", s.shutdownHandler)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func (s *ShutdownListener) shutdownHandler(w http.ResponseWriter, r *http.Request) {
	for _, server := range s.servers {
		server.Shutdown(nil)
	}
}
