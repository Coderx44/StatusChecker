package main

import (
	"log"
	"net/http"

	"github.com/Coderx44/StatusChecker/service"
)

const PORT = ":3000"

func main() {
	checkHttp := service.HttpChecker{}

	http.HandleFunc("/website", service.HandleWebsites)

	go service.CheckWebsites(&checkHttp)
	log.Fatal(http.ListenAndServe(PORT, nil))

}
