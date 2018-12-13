package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/fcgi"
	"net/http/httputil"
	"os"
	"runtime"
	"time"
)

var appAddr string

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	appAddr = os.Getenv("APPADDR") // e.g. "APPADDR=0.0.0.0:3000"
}

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

	var err error
	if appAddr != "" {
		// Run as a local web server
		mux.HandleFunc("/time", timeHandler(time.RFC1123))
		log.Println("Listening on " + appAddr + "...")
		err = http.ListenAndServe(appAddr, mux)
	} else {
		// Run as FCGI via standard I/O
		mux.HandleFunc("/fcgi-bin/time.fcgi/time", timeHandler(time.RFC1123))
		err = fcgi.Serve(nil, mux)
	}
	if err != nil {
		log.Fatal(err)
	}

}
