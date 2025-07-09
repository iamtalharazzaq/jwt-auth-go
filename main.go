package main

import (
	"JWT/handlers"
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Starting server on :8080")
	http.HandleFunc("/login", logAction(handlers.Login, "/login"))
	http.HandleFunc("/logout", logAction(handlers.Logout, "/logout"))
	http.HandleFunc("/home", logAction(handlers.Home, "/home"))
	http.HandleFunc("/refresh", logAction(handlers.Refresh, "/refresh"))

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("Server error: %v\n", err)
	}
}

// logAction wraps an http.HandlerFunc to print the action to the console
func logAction(h http.HandlerFunc, route string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("%s %s\n", r.Method, route)
		h(w, r)
	}
}
