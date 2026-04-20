package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()

	router.GET("/", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		fmt.Fprint(w, "Hello, World")
	})

	server := http.Server{
		Handler: router,
		Addr:    "localhost:8080",
	}

	server.ListenAndServe()

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
