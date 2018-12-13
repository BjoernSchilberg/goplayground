package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"time"
)

func logRequest(r *http.Request) string {
	requestDump, err := httputil.DumpRequest(r, true)
	if err != nil {
		fmt.Println(err)
	}
	return string(requestDump)
}

func timeHandler(format string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tm := time.Now().Format(format)
		w.Write([]byte("The time is: " + tm + "\n"))
		log.Println(logRequest(r))
	}
}

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/time", timeHandler(time.RFC1123))

	log.Println("Listening on :3000 ...")
	http.ListenAndServe(":3000", mux)

}
