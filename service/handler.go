package service

import (
	"encoding/json"
	"net/http"
)

var WebsiteList = make(map[string]string)

func HandlePostWebsites(w http.ResponseWriter, r *http.Request) {
	request := make(map[string][]string)

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	urlList := request["websites"]
	for _, url := range urlList {

		WebsiteList[url] = "Unknown"
	}

	w.WriteHeader(http.StatusOK)

}
