package server

import (
	"net/http"

	"github.com/Coderx44/StatusChecker/statuschecker"
)

func InitRouter(dep dependencies) {
	http.HandleFunc("/website", statuschecker.HandleWebsites(dep.httpchecker))

}
