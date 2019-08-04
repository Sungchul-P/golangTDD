package handler

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Handler struct {
	router *mux.Router
}

func (h Handler) Init() {
	h.router = mux.NewRouter()
	h.router.HandleFunc("/ping", h.Ping).Methods("GET")
	http.Handle("/", h.router)
}

func (h Handler) Ping(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "pong")
}

func (h Handler) Calc(w http.ResponseWriter, r *http.Request) {

}
