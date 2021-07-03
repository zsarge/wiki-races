package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/gorilla/mux"
)

type spaHandler struct {
	staticPath string
	indexPath  string
}

// Keep log of pages handled
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

		log.Printf("Handled Wiki Page #%d in %s", numberOfPagesHandled, duration)
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

// STOLEN FROM: https://github.com/gorilla/mux#serving-single-page-applications
// ServeHTTP inspects the URL path to locate a file within the static dir
// on the SPA handler. If a file is found, it will be served. If not, the
// file located at the index path on the SPA handler will be served. This
// is suitable behavior for serving an SPA (single page application).
func (h spaHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// get the absolute path to prevent directory traversal
	path, err := filepath.Abs(r.URL.Path)
	if err != nil {
		// if we failed to get the absolute path respond with a 400 bad request
		// and stop
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// prepend the path with the path to the static directory
	path = filepath.Join(h.staticPath, path)

	// check whether a file exists at the given path
	_, err = os.Stat(path)
	if os.IsNotExist(err) {
		// file does not exist, serve index.html
		http.ServeFile(w, r, filepath.Join(h.staticPath, h.indexPath))
		return
	} else if err != nil {
		// if we got an error (that wasn't that the file doesn't exist) stating the
		// file, return a 500 internal server error and stop
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// otherwise, use http.FileServer to serve the static dir
	http.FileServer(http.Dir(h.staticPath)).ServeHTTP(w, r)
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/wiki/{name}", wikiHandler)

	abs, err := filepath.Abs("../frontend/dist")
	if err != nil {
		log.Panic("File could not be found")
	}

	spa := spaHandler{staticPath: abs, indexPath: "index.html"}
	r.PathPrefix("/").Handler(spa)

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
