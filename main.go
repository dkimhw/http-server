package main

import (
	"log"
	"net/http"
)

func main() {
	// HandlerFunc type is an adapter to allow the use of ordinary functions as HTTP handlers
	// If f is a function with the appropriate signature, HandlerFunc(f) is a Handler that calls f.
	//  By type casting our PlayerServer function with it, we have now implemented the required Handler.
	// type HandlerFunc func(ResponseWriter, *Request)
	server := &PlayerServer{}

	/*
		ListenAndServe takes a port to listen on a Handler.
		If there is a problem the web server will return an error,
		an example of that might be the port already being listened to.
		For that reason we wrap the call in log.Fatal to log the error to the user.
	*/
	log.Fatal(http.ListenAndServe(":5001", server))
}
