package main

import (
	"github.com/gonozov0/go-musthave-shortener/internal/app"
	"log"
	"net/http"
)

func main() {
	rootHandler := app.NewRootHandler()
	//fmt.Println(rootHandler)
	http.Handle("/", &rootHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
