package main

import (
	"log"
	"net/http"
	"time"
)

func timeHandler(format string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tm := time.Now().Format(format)
		w.Write([]byte("The time is: " + tm))
	}
}

func main() {

	//Attention: Generally you shouldn't use the DefaultServeMux because it
	//poses a security risk. This is only for example.

	http.Handle("/time", timeHandler(time.RFC1123))

	log.Println("Listening...")

	http.ListenAndServe(":3000", nil)
}
