package main

import (
	"context"
	"net/http"

	"github.com/Coderx44/StatusChecker/service"
)

type StatusChecker interface {
	Check(ctx context.Context, name string) (status bool, err error)
}

type httpChecker struct {
}

func main() {
	// checkHttp := httpChecker{}

	http.HandleFunc("/website", service.HandleWebsites)

	// go checkWebsites(&checkHttp)

}
