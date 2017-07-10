package main

import (
	"log"
	"net/http"

	"github.com/acoshift/gonews/pkg/app"
	"github.com/acoshift/gonews/pkg/model"
)

const (
	port     = ":8080"
	mongoURL = "mongodb://127.0.0.1:27017"
)

func main() {
	mux := http.NewServeMux()
	app.Mount(mux)
	err := model.Init(mongoURL)
	if err != nil {
		log.Fatalf("can not init model; %v", err)
	}
	http.ListenAndServe(port, mux)
}
