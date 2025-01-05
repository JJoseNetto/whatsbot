package main

import (
	"whatsapp-gtm/controllers"
	"whatsapp-gtm/middleware"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

func main() {
	r := chi.NewRouter()
	r.Post("/api/webhook", controllers.Webhook)
	r.With(middleware.Auth).Post("/api/send", controllers.Send)
	r.With(middleware.Auth).Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Pong")
	})
	fmt.Println("Server running in port 3000")
	log.Fatal(http.ListenAndServe("0.0.0.0:3000", r))
}
