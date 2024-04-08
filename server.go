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
	http.HandleFunc("/", serveMainPage)
	http.HandleFunc("/action", handleAction)
	http.HandleFunc("/decoder", handleDecoder)

	// Start the server on port 8080
	fmt.Println("Starting server at port 8080")
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
	tmpl, err := template.ParseFiles("index.html")
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

func handleDecoder(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		// Process the form data
		r.ParseForm()
		encodedText := r.FormValue("text")

		decodedText, err := DecodeArt(encodedText)
		if err != nil {
			// Handle the error
			fmt.Printf("Error decoding text: %s\n", err)
			return
		}

		// Respond with the decoded text
		fmt.Fprintf(w, decodedText)
	} else if r.Method == "GET" {
		// Serve a page with the last decoded string
		// You would need to implement storage and retrieval of the last decoded string
	} else {
		http.Error(w, "Invalid request method.", http.StatusMethodNotAllowed)
	}
}

func handleAction(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		r.ParseForm()
		text := r.FormValue("text")
		action := r.FormValue("action")

		var result string
		var err error

		if action == "encode" {
			result = EncodeArt(text)
		} else if action == "decode" {
			result, err = DecodeArt(text)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
		} else {
			http.Error(w, "Invalid action", http.StatusBadRequest)
			return
		}

		fmt.Fprintf(w, result)
	} else {
		http.Error(w, "Invalid request method.", http.StatusMethodNotAllowed)
	}
}
