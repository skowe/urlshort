package main

import (
	"flag"
	"log"
	"net/http"

	jsu "github.com/skowe/jsonurlshort"
)

//Recieve input flag for redirect.json
//Use it's contents to create a handler
//Use the handler to serve requests

func main() {
	f := flag.String("redirect", "redirect.json", "-redirect=[path-to-file]")
	flag.Parse()

	mux := http.NewServeMux()

	handler, err := jsu.GetHandler(*f)
	if err != nil {
		log.Fatalln(err)
	}

	mux.Handle("/", handler)
	log.Fatalln(http.ListenAndServe(":8080", mux))
}
