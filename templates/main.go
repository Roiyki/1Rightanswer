package main

import (
	"html/template"
	"log"
	"net/http"
)

// Flight represents flight data
type Flight struct {
	ID          int
	Origin      string
	Destination string
	DepartDate  string
	ReturnDate  string
}

// Handler for the home page
func indexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./templates/index.html"))
	err := tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// Handler for the flight details page
func flightHandler(w http.ResponseWriter, r *http.Request) {
	flight := Flight{
		ID:          1,
		Origin:      "New York",
		Destination: "Los Angeles",
		DepartDate:  "15/05/2024",
		ReturnDate:  "24/05/2024",
	}
	tmpl := template.Must(template.ParseFiles("./templates/flight.html"))
	err := tmpl.Execute(w, flight)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/flight", flightHandler)
	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
