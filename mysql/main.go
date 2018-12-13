package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/fcgi"
	"net/http/httputil"
	"os"
	"runtime"

	_ "github.com/go-sql-driver/mysql"
)

var appAddr string
var dbUser string
var dbName string
var dbPasswd string

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	appAddr = os.Getenv("APPADDR") // e.g. "APPADDR=0.0.0.0:3000"
	dbUser = os.Getenv("DBUSER")
	dbName = os.Getenv("DBNAME")
	dbPasswd = os.Getenv("DBPASSWD")
}

func logRequest(r *http.Request) string {
	requestDump, err := httputil.DumpRequest(r, true)
	if err != nil {
		fmt.Println(err)
	}
	return string(requestDump)
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func getTermine(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		termine, err := getTermineFromDB(db)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		respondWithJSON(w, http.StatusOK, termine)
	}
}

func main() {
	var err error

	connectionString :=
		fmt.Sprintf("%s:%s@/%s", dbUser, dbPasswd, dbName)

	db, err := sql.Open("mysql", connectionString)

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	mux := http.NewServeMux()

	if appAddr != "" {
		// Run as a local web server
		mux.HandleFunc("/termine", getTermine(db))
		log.Println("Listening on " + appAddr + "...")
		err = http.ListenAndServe(appAddr, mux)
	} else {
		// Run as FCGI via standard I/O
		mux.HandleFunc("/fcgi-bin/time.fcgi/termine", getTermine(db))
		err = fcgi.Serve(nil, mux)
	}
	if err != nil {
		log.Fatal(err)
	}

}
