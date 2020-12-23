package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// Response Struct
type Response struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}

// NewResponse Create new Response
func NewResponse(status bool, message string) *Response {
	return &Response{
		Status:  status,
		Message: message,
	}
}

// HomeHandler Route
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusBadRequest)

	res := NewResponse(true, "Welcome")
	jsonData, err := json.Marshal(res)
	if err != nil {
		log.Fatal("Error Occured.")
	}

	w.Write(jsonData)
}

// ReactHandler Handler
func ReactHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusBadRequest)

	res := NewResponse(true, vars["name"])
	jsonData, err := json.Marshal(res)
	if err != nil {
		log.Fatal("Error Occured.")
	}

	w.Write(jsonData)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/{name}", ReactHandler)

	srv := &http.Server{
		Handler:      r,
		Addr:         ":80",
		WriteTimeout: 10 * time.Second,
		ReadTimeout:  10 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
