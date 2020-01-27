package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Get(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, item := range database {
		if item.Id == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	var resp Resp
	resp.Id = params["id"]
	resp.Progress = 0
	database = append(database, resp)
	json.NewEncoder(w).Encode(resp)
	return
}
func Patch(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, item := range database {
		if item.Id == params["id"] {
			item.Progress++
			return
		}
	}
}

type Resp struct {
	Id       string `json:"id,omitemoty"`
	Progress int    `json:"progress,omitempty"`
}

var database []Resp

// função principal
func main() {
	router := mux.NewRouter()
	router.HandleFunc("/player/progress/{id}", Get).Methods("GET")
	router.HandleFunc("/player/progress/{id}", Patch).Methods("PATCH")
	log.Fatal(http.ListenAndServe(":8091", router))
}
