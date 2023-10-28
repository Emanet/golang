package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

var myValue = "Koyimde turat"

func main() {
	// create a server
	myServer := &http.Server{
		// set the server address
		Addr: "127.0.0.1:8080",
		// define some specific configuration
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	// launch the server
	http.HandleFunc("/home", HelloServer)
	http.HandleFunc("/set/", ByeServer)
	log.Println("listen on", myServer.Addr)
	log.Fatal(myServer.ListenAndServe())
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<title>Document</title>
	</head>
	<body>
		<h1>`+myValue+`</h1>
	</body>
	</html>`)
}

func ByeServer(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("r.URL.Path: %v\n", r.URL.Path)
	myValue = r.URL.Path[5:]
	http.Redirect(w, r, "/home", http.StatusTemporaryRedirect)
}
