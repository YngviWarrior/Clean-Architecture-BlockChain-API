package controllers

import (
	"fmt"
	"net/http"
)

type Controller interface {
	GetExchange(w http.ResponseWriter, r *http.Request)
	HomeLink(w http.ResponseWriter, r *http.Request)
}

func HomeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home!")
}
