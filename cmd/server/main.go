package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"go-server/internal/api"
	"go-server/internal/auth"
	"go-server/internal/store"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	handler := &api.Handler{Store: store.NewStore()}
	mux := http.NewServeMux()

	mux.Handle("GET /items", auth.BasicAuth(http.HandlerFunc(handler.List)))
	mux.Handle("POST /items", auth.BasicAuth(http.HandlerFunc(handler.Create)))
	mux.Handle("GET /items/{id}", auth.BasicAuth(http.HandlerFunc(handler.Get)))
	mux.Handle("PUT /items/{id}", auth.BasicAuth(http.HandlerFunc(handler.Update)))
	mux.Handle("DELETE /items/{id}", auth.BasicAuth(http.HandlerFunc(handler.Delete)))

	srv := &http.Server{
		Addr:              ":" + port,
		Handler:           mux,
		ReadHeaderTimeout: 5 * time.Second,
		WriteTimeout:      10 * time.Second,
		IdleTimeout:       15 * time.Second,
	}

	fmt.Printf("Server is running on http://localhost:%s\n", port)
	log.Fatal(srv.ListenAndServe())
}
