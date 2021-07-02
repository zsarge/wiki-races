package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

var numberOfPagesHandled = 1

func wikiHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]

	w.Header().Set("Content-Type", "text/html")
	if name != "" {
		start := time.Now()
		text, err := getPageText(name)
		if err != nil {
			log.Panic(err)
		}

		duration := time.Since(start)

		log.Printf("Handled Page #%d in %s", numberOfPagesHandled, duration)
		numberOfPagesHandled++

		_, err = io.WriteString(w, text)
		if err != nil {
			log.Panic(err)
		}
	} else {
		w.WriteHeader(http.StatusNotFound)

		fmt.Fprintf(w, "404! Not found")
	}
}

// func homeHandler(w http.ResponseWriter, r *http.Request) {

// }

func main() {
	r := mux.NewRouter()
	// r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/wiki/{name}", wikiHandler)
	// r.PathPrefix("/").Handler(catchAllHandler)
	// http.Handle("/", r)

	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:3000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	fmt.Println("Server listening on port 3000")
	log.Fatal(srv.ListenAndServe())
}
