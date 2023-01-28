package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/Coderx44/StatusChecker/service"
)

const PORT = ":3000"

type StatusChecker interface {
	Check(ctx context.Context, name string) (status bool, err error)
}

type httpChecker struct {
}

func (h httpChecker) Check(ctx context.Context, name string) (status bool, err error) {

	res, err := http.Get("http://" + name)
	if err != nil {
		status = false
		return
	}

	if res.StatusCode == 200 {
		status = true
	} else {
		status = false
	}
	return
}

func checkWebsites(checkHttp *httpChecker) {
	for {
		for url := range service.WebsiteList {
			status, err := checkHttp.Check(context.Background(), url)

			if err != nil || !status {
				service.WebsiteList[url] = "DOWN"
			} else if status {
				service.WebsiteList[url] = "UP"
			}
		}

		time.Sleep(time.Minute)
	}
}

func main() {
	checkHttp := httpChecker{}

	http.HandleFunc("/website", service.HandleWebsites)

	go checkWebsites(&checkHttp)
	log.Fatal(http.ListenAndServe(PORT, nil))

}
