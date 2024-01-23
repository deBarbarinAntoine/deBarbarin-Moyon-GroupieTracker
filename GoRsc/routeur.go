package server

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"runtime"
)

var (
	_, b, _, _ = runtime.Caller(0)
	path       = filepath.Dir(filepath.Dir(b)) + "/"
)

func RUN() {

	// used same system than hangman, since it was working prety well
	http.HandleFunc("/", ErrorHandler)
	http.HandleFunc("/home", HomeHandler)
	http.HandleFunc("/select", SelectHandler)
	http.HandleFunc("/search", SearchHandler)

	// Serve static files from the "assets" directory
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(path+"assets"))))

	// Print statement indicating server is running << same
	fmt.Println("Server is running on :8080 http://localhost:8080/home")

	// Start the server << same
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
