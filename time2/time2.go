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

	mux := http.NewServeMux()

	mux.HandleFunc("/time", timeHandler(time.RFC1123))

	log.Println("Listening on :3000 ...")
	http.ListenAndServe(":3000", mux)

}
