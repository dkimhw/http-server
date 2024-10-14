package main

import (
	"log"
	"net/http"
)

// curl -X POST http://localhost:5001/players/Pepper
// curl http://localhost:5001/players/Pepper
// curl http://localhost:5001/league
func main() {
	server := NewPlayerServer(NewInMemoryPlayerStore())

	/*
		ListenAndServe takes a port to listen on a Handler.
		If there is a problem the web server will return an error,
		an example of that might be the port already being listened to.
		For that reason we wrap the call in log.Fatal to log the error to the user.
	*/
	log.Fatal(http.ListenAndServe(":5001", server))
}
