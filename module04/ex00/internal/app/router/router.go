package router

import (
	"encoding/json"
	"log"
	"net/http"

	candy "github.com/Arclight-V/Golang/tree/module04-ex00/module04/ex00/internal/service/routs/models"
	"github.com/golang/gddo/httputil/header"
)

type Router struct {
}

func NewRouter() *Router {

	return &Router{}
}

func (rout *Router) handler(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Content-Type") != "" {
		value, _ := header.ParseValueAndParams(r.Header, "Content-Type")
		if value != "application/json" {
			msg := "Content-Type header is not application/json"
			http.Error(w, msg, http.StatusUnsupportedMediaType)
			return
		}
	}

	r.Body = http.MaxBytesReader(w, r.Body, 1048576)

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	var c candy.CandyObjectRequest
	err := dec.Decode(&c)
	if err != nil {

	}
}

func (rout *Router) Start() {
	http.HandleFunc("/", rout.handler)
	log.Fatal(http.ListenAndServe("localhost:3333", nil))
}
