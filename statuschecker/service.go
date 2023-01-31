package statuschecker

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
)

type StatusChecker interface {
	Check(ctx context.Context, name string) (status string, err error)
}

type HttpChecker struct {
}

func NewHttpChecker() StatusChecker {
	return &HttpChecker{}
}

func (h HttpChecker) Check(ctx context.Context, name string) (status string, err error) {

	status, ok := WebsiteList[name]

	if !ok {
		return "DOWN", fmt.Errorf("%s", "website not found")
	}
	return status, nil

}

func CheckWebsites() {

	for {

		for url := range WebsiteList {
			log.Println(WebsiteList, " url")
			go func(url string) {
				res, err := http.Get("http://" + url)
				mut.Lock()
				if err != nil || res.StatusCode != http.StatusOK {
					WebsiteList[url] = "DOWN"
				} else {
					WebsiteList[url] = "UP"
				}
				mut.Unlock()

			}(url)
		}
		time.Sleep(time.Minute)
	}
}

//maps are not goroutine favorable.
//use syncmaps.

//router -> handler -> service -> repo -> DB{mysql, mongo, etc}
//chronjob
