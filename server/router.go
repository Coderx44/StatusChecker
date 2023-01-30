package server

import (
	"net/http"

	"github.com/Coderx44/StatusChecker/service"
)

func InitRouter(dep dependencies) {
	http.HandleFunc("/website", service.HandleWebsites(dep.httpchecker))

}
