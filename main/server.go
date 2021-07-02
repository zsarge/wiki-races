package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func wikiHandler(w http.ResponseWriter, r *http.Request) {
	text, err := getPageText("Red")
	if err != nil {
		log.Panic(err)
	}
	log.Println("Page handled")

	_, err = io.WriteString(w, text)
	if err != nil {
		log.Panic(err)
	}
}

func main() {
	http.HandleFunc("/", wikiHandler)
	fmt.Println("Server listening on port 3000")
	log.Panic(
		http.ListenAndServe(":3000", nil),
	)
}
