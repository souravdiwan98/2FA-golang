package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type User struct {
	Username string
	Password string
	Secret   string
}

// Simple in-memory "database" for demonstration purposes
var users = map[string]*User{
	"john": {Username: "john", Password: "password", Secret: ""},
}

var templates = template.Must(template.ParseGlob("templates/*.html"))

func homeHandler(w http.ResponseWriter, r *http.Request) {
	err := templates.ExecuteTemplate(w, "index.html", nil)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		// Render the login page for GET requests
		err := templates.ExecuteTemplate(w, "login.html", nil)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		return
	}
	// Handle POST requests for user authentication
	// (Code for handling form submission and user authentication)
}

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/dashboard", dashboardHandler)
	http.HandleFunc("/generate-otp", generateOTPHandler)
	http.HandleFunc("/validate-otp", validateOTPHandler)

	// Start the server
	fmt.Println("Starting server at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
