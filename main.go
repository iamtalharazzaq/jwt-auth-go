package main

import (
	"JWT/handlers"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/login", handlers.Login)
	http.HandleFunc("/logout", handlers.Logout)
	http.HandleFunc("/home", handlers.Home)
	http.HandleFunc("/refresh", handlers.Refresh)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
