package main

import (
	"log"
	"net/http"

	"github.com/ks8dev/nero/api/handler"
	"github.com/ks8dev/nero/api/middleware"
)

func main() {
	// file server for public files like robots.txt, favicon.ico
	fileServer := http.NewServeMux()
	fileServer.Handle("GET /", http.FileServer(http.Dir("./public")))

	// routes
	http.Handle("GET /", middleware.Public(fileServer)(handler.Home()))

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
