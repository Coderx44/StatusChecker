package service

import (
	"encoding/json"
	"net/http"
)

var WebsiteList = make(map[string]string)

func HandleWebsites(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:
		HandleGetWebsites(w, r)

	case http.MethodPost:
		HandlePostWebsites(w, r)

	default:
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}

}

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

func HandleGetAllWebsites(w http.ResponseWriter, r *http.Request) {

	json.NewEncoder(w).Encode(WebsiteList)
}

func HandleGetOneWebsite(w http.ResponseWriter, r *http.Request, url string) {

	status, ok := WebsiteList[url]

	if !ok {
		http.Error(w, "Website not found", http.StatusNotFound)
		return
	}

	website := map[string]string{}
	website[url] = status
	json.NewEncoder(w).Encode(website)
}

func HandleGetWebsites(w http.ResponseWriter, r *http.Request) {

	url := r.URL.Query().Get("name")

	if url == "" {
		HandleGetAllWebsites(w, r)
		return
	}

	HandleGetOneWebsite(w, r, url)

}
