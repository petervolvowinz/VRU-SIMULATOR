/*
2020
*/

package main

import (
	"log"
	"net/http"
)
const(
	ReleaseTag = "release-0.0.1-dev"
)

func pingHandler(w http.ResponseWriter, r *http.Request){
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(" HTTP status(200) OK code returned running:  " + ReleaseTag))
}

func ServeSimulationView() error {
	http.HandleFunc("/version", pingHandler)
	http.Handle("/", http.FileServer(http.Dir("./static")))
	return http.ListenAndServe(":8081", nil)
}

func main(){
	log.Println("starting vru simulation server...")
	// Hold a server loop
	log.Fatal(ServeSimulationView())
}
