package service

import (
	"context"
	"encoding/json"
	"net/http"
	"sync"
)

var WebsiteList = make(map[string]string)
var mut sync.Mutex

func HandleWebsites(st StatusChecker) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			HandleGetWebsites(st, w, r)

		case http.MethodPost:
			HandlePostWebsites(st, w, r)

		default:
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		}
	})

}

func HandleGetWebsites(st StatusChecker, w http.ResponseWriter, r *http.Request) {

	url := r.URL.Query().Get("name")
	if url == "" {
		HandleGetAllWebsites(st, w, r)
		return
	}

	HandleGetOneWebsite(st, url, w, r)

}

func HandlePostWebsites(st StatusChecker, w http.ResponseWriter, r *http.Request) {

	request := make(map[string][]string)

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	urlList := request["websites"]
	for _, url := range urlList {
		if _, ok := WebsiteList[url]; !ok {
			mut.Lock()
			WebsiteList[url] = "Unknown"
			mut.Unlock()
		}
	}
	w.WriteHeader(http.StatusOK)

}

func HandleGetAllWebsites(st StatusChecker, w http.ResponseWriter, r *http.Request) {

	statusList := make(map[string]string)
	for url := range WebsiteList {
		statusList[url], _ = st.Check(context.Background(), url)
	}
	json.NewEncoder(w).Encode(statusList)

}

func HandleGetOneWebsite(st StatusChecker, url string, w http.ResponseWriter, r *http.Request) {

	statusList := make(map[string]string)
	status, err := st.Check(context.Background(), url)

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	statusList[url] = status
	json.NewEncoder(w).Encode(statusList)

}
