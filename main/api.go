package main

import (
	"cgt.name/pkg/go-mwclient" // imports "mwclient"
	"github.com/antonholmquist/jason"
)

func getPage(name string) (*jason.Object, error) {

	// Initialize a *Client with New(), specifying the wiki's API URL
	// and your HTTP User-Agent. Try to use a meaningful User-Agent.
	w, err := mwclient.New("https://en.wikipedia.org/w/api.php", "wiki-races")
	if err != nil {
		return nil, err
	}

	// Specify parameters to send.
	parameters := map[string]string{
		"action": "parse",
		"format": "json",
	}
	parameters["page"] = name

	// Make the request.
	resp, err := w.Get(parameters)
	if err != nil {
		return nil, err
	}

	// Print the *jason.Object
	return resp, nil
}

func getPageText(name string) (string, error) {
	page, err := getPage(name)
	if err != nil {
		return "", err
	}
	parse, err := page.GetObject("parse")
	if err != nil {
		return "", err
	}
	text, err := parse.GetString("text")
	if err != nil {
		return "", err
	}

	return text, nil
}
