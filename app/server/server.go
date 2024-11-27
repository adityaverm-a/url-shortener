package server

import (
	"log"
	"net/http"
	"time"
)

// Serve creates and starts an HTTP server instance
func Serve(port string, router http.Handler) {
	server := createServerInstance(port, router)

	startServer(server)
	log.Println("Server stopped")
}

// createServerInstance creates and returns an HTTP server instance
func createServerInstance(port string, router http.Handler) *http.Server {
	return &http.Server{
		Addr:         port,
		Handler:      router,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 100 * time.Second,
		IdleTimeout:  15 * time.Second,
	}
}

// startServer starts listening and serving HTTP requests on the provided server instance
func startServer(server *http.Server) {
	log.Println("Server is ready to handle requests at", server.Addr)

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Println("Could not listen on %s: %v\n", server.Addr, err.Error())
	}
}
