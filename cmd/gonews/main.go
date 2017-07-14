package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/acoshift/gonews/pkg/app"
	"github.com/acoshift/gonews/pkg/model"
)

func main() {
	var config struct {
		MongoURL string `json:"mongo_url"`
	}
	b, err := ioutil.ReadFile("config.json")
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(b, &config)
	if err != nil {
		log.Fatal(err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	mux := http.NewServeMux()
	app.Mount(mux)
	err = model.Init(config.MongoURL)
	if err != nil {
		log.Fatalf("can not init model; %v", err)
	}
	http.ListenAndServe(":"+port, mux)
}
