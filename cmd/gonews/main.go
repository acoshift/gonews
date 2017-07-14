package main

import (
	"log"
	"net/http"
	"os"

	"github.com/acoshift/gonews/pkg/app"
	"github.com/acoshift/gonews/pkg/model"
)

const (
	// mongoURL = "mongodb://127.0.0.1:27017"
	mongoURL = "mongodb://gonews:9NAJHQXu4M9CvKqD@ds147872.mlab.com:47872/gonews"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	mux := http.NewServeMux()
	app.Mount(mux)
	err := model.Init(mongoURL)
	if err != nil {
		log.Fatalf("can not init model; %v", err)
	}
	http.ListenAndServe(":"+port, mux)
}
