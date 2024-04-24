package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type PageData struct {
	Title string
}

func startServer() {

	// Serve static files with correct MIME types
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static", fs))

	// Register your route handlers for the server
	http.HandleFunc("/", serveMainPage)
	http.HandleFunc("/encoder", handleAction)
	http.HandleFunc("/decoder", handleAction)

	// Start the server on port 8080
	log.Print("Starting server at port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func serveMainPage(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	// Load the main page template
	tmpl, err := template.ParseFiles("static/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := PageData{
		Title: "Text Art Decoder",
	}

	// Serve the main page
	tmpl.Execute(w, data)
}

func handleAction(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Invalid request method.", http.StatusMethodNotAllowed)
		return
	}

	r.ParseForm()
	text := r.FormValue("text")
	action := r.FormValue("action")

	var result string
	var err error
	var readyToSendResult bool // For sending the result in the response

	switch action {
	case "encode":
		log.Printf("Received text to encode: %s", text)
		result = EncodeArt(text) // Replace with your actual encoding function
		readyToSendResult = true // Assuming EncodeArt is synchronous
		log.Printf("Encoded text:  %s", result)

	case "decode":
		log.Printf("Received text to decode: %s", text)
		result, err = DecodeArt(text) // Replace with your actual decoding function
		readyToSendResult = true      // Assuming DecodeArt is synchronous
		if err != nil {
			log.Printf("Error decoding text: %s\n", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		log.Printf("Decoded text:  %s", result)

	default:
		http.Error(w, "Invalid action", http.StatusBadRequest)
		return
	}

	if readyToSendResult {
		w.WriteHeader(http.StatusAccepted)
		fmt.Fprintf(w, result)
	} else {
		w.WriteHeader(http.StatusAccepted) // Send 202 without the result initially
	}
}
