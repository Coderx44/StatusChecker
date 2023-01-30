package service

import (
	"context"
	"net/http"
	"time"
)

type StatusChecker interface {
	Check(ctx context.Context, name string) (status bool, err error)
}

type HttpChecker struct {
}

func (h HttpChecker) Check(ctx context.Context, name string) (status bool, err error) {

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

func CheckWebsites(checkHttp *HttpChecker) {
	for {

		for url := range WebsiteList {
			go func(url string) {
				status, err := checkHttp.Check(context.Background(), url)

				if err != nil || !status {
					WebsiteList[url] = "DOWN"
				} else if status {
					WebsiteList[url] = "UP"
				}
			}(url)
		}

		time.Sleep(time.Minute)
	}
}
