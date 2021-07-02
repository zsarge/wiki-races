package main

import (
	"fmt"
	"log"
	"net/http"
)

func homePageHandler(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "Hello world!")
	if err != nil {
		log.Panic(err)
	}
}

func main() {
	http.HandleFunc("/", homePageHandler)

	fmt.Println(test())
	fmt.Println("Server listening on port 3000")

	log.Panic(
		http.ListenAndServe(":3000", nil),
	)

}

