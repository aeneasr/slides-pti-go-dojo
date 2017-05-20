package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"github.com/ory-am/go-dbg/http/helper"
	"encoding/json"
	"log"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", home)
	r.HandleFunc("/hi/{greeting}", greeting)

	log.Fatal(http.ListenAndServe(":8090", r))
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte{'h', 'i', '!'})
}

type response struct {
	Now string `json:"now"`
	Greeting string `json:"greeting"`
}

func greeting(w http.ResponseWriter, r *http.Request) {
	greeting := mux.Vars(r)["greeting"]

	resp := &response{
		Now: helper.PrintTime(),
		Greeting: greeting,
	}

	out, err := json.MarshalIndent(resp, "", "\t")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(out)
}
