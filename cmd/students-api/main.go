package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/UtkarshSahu9906/students-api/inernal/config"
)

func main() {
	cfg := config.MustLoad()

	 router :=http.NewServeMux()

	 router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		//w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello, Students API!"))
	})
	server := &http.Server{
		Addr: cfg.HTTPServer.Addr,
		Handler: router,
	}
		fmt.Printf("Server is running on %s\n", cfg.HTTPServer.Addr)

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal("Failed to start server")
	}


}
