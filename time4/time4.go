package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"time"
)

// The Time is now
type Time struct {
	Time string
}

func logRequest(r *http.Request) string {
	requestDump, err := httputil.DumpRequest(r, true)
	if err != nil {
		fmt.Println(err)
	}
	return string(requestDump)
}

func timeHandler(format string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		time := Time{time.Now().Format(format)}

		js, err := json.Marshal(time)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
		log.Println(logRequest(r))
	}
}

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/time", timeHandler(time.RFC1123))

	log.Println("Listening on :3000 ...")
	http.ListenAndServe(":3000", mux)

}
