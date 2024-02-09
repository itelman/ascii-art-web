package main

import (
	"fmt"
	"functions/functions"
	"net/http"
)

func main() {
	// Run the server: http://localhost:8080/

	http.HandleFunc("/", functions.IndexHandler)
	http.HandleFunc("/ascii-art", functions.AsciiArtHandler)
	http.HandleFunc("/style.css", functions.ServeCss)
	fmt.Println("Server is running on http://localhost:8080/...")
	http.ListenAndServe(":8080", nil)
}
