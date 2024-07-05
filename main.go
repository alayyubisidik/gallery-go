package main

import (
	"gallery_go/helper"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func NewServer(router *httprouter.Router) *http.Server {
	return &http.Server{
		Addr:    "localhost:3000",
		Handler: router,
	}
}

func main() {
	server := initializedServer()
	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
