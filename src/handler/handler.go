package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Handler struct {
	router *mux.Router
}

func (h Handler) Init() {
	h.router = mux.NewRouter()
	h.router.HandleFunc("/ping", h.Ping).Methods("GET")
	h.router.HandleFunc("/div/{a}/{b}", h.Div).Methods("GET")
	http.Handle("/", h.router)
}

func (h Handler) Ping(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "pong")
}

func (h Handler) Div(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	a := vars["a"]
	b := vars["b"]

	ai, err := strconv.Atoi(a)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	bi, err := strconv.Atoi(b)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if bi == 0 {
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}
	if ai == 0 {
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}

	fmt.Fprintf(w, "%d", ai/bi)
}
