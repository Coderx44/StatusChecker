package main

import (
	"log"
	"net/http"

	"github.com/Coderx44/StatusChecker/server"
	"github.com/Coderx44/StatusChecker/statuschecker"
)

const PORT = ":3000"

func main() {
	dependencies, err := server.InitDependencies()
	if err != nil {
		log.Println(err)
	}
	server.InitRouter(dependencies)

	go statuschecker.CheckWebsites()
	log.Fatal(http.ListenAndServe(PORT, nil))

}
