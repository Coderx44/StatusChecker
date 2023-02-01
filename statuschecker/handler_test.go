package statuschecker_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/Coderx44/StatusChecker/mocks"
	"github.com/Coderx44/StatusChecker/statuschecker"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type HandlerTestSuite struct {
	suite.Suite
	service *mocks.StatusChecker
}

func (suite *HandlerTestSuite) SetupTest() {
	suite.service = &mocks.StatusChecker{}
}

func (suite *HandlerTestSuite) TearDownTest() {
	t := suite.T()
	suite.service.AssertExpectations(t)
}

func TestHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(HandlerTestSuite))
}

func (suite *HandlerTestSuite) TestHandleGetOneWebsite() {
	t := suite.T()

	t.Run("when website is present in map", func(t *testing.T) {

		res := "UP"

		queryParams := url.Values{}
		queryParams.Add("name", "google.com")

		getOneurl := fmt.Sprintf("/website?%s", queryParams.Encode())

		r := httptest.NewRequest(http.MethodGet, getOneurl, nil)
		w := httptest.NewRecorder()

		suite.service.On("Check", mock.Anything, "google.com").Return(res, nil)

		statuschecker.HandleGetOneWebsite(suite.service, w, r)
		goRes := make(map[string]string)
		json.NewDecoder(w.Body).Decode(&goRes)
		assert.Equal(t, http.StatusOK, w.Result().StatusCode)
		assert.Equal(t, res, goRes["google.com"])

	})

	t.Run("when website is not present in map", func(t *testing.T) {

		res := ""
		inputUrl := "fakewebsite.com"
		queryParams := url.Values{}
		queryParams.Add("name", inputUrl)

		getOneurl := fmt.Sprintf("/website?%s", queryParams.Encode())

		r := httptest.NewRequest(http.MethodGet, getOneurl, nil)
		w := httptest.NewRecorder()

		suite.service.On("Check", mock.Anything, inputUrl).Return(res, errors.New("website not found"))

		statuschecker.HandleGetOneWebsite(suite.service, w, r)
		var goRes string
		json.NewDecoder(w.Body).Decode(&goRes)
		assert.Equal(t, http.StatusNotFound, w.Result().StatusCode)
		assert.Equal(t, res, goRes)

	})
}

func (suite *HandlerTestSuite) TestaddWebsiteHandler() {
	t := suite.T()
	t.Run("when post is successful", func(t *testing.T) {
		body := []byte(`{"websites":["www.google.com", "www.facebook.com"]}`)
		r := httptest.NewRequest(http.MethodPost, "/website", bytes.NewBuffer(body))

		w := httptest.NewRecorder()

		statuschecker.AddWebsiteHandler(suite.service, w, r)
		assert.Equal(t, http.StatusOK, w.Result().StatusCode)

		if len(statuschecker.WebsiteList) != 2 {
			t.Errorf("unexpected number of websites in WebsiteList: got %d, want 2", len(statuschecker.WebsiteList))
		}
		for k, v := range statuschecker.WebsiteList {
			if v != "Unknown" {
				t.Errorf("unexpected value for %s in WebsiteList: got %v, want 'Unknown'", k, v)
			}
		}

	})

	t.Run("when post is not valid", func(t *testing.T) {
		r := httptest.NewRequest(http.MethodPost, "/website", strings.NewReader(""))

		w := httptest.NewRecorder()

		statuschecker.AddWebsiteHandler(suite.service, w, r)
		assert.Equal(t, http.StatusBadRequest, w.Result().StatusCode)

	})

}

func (suite *HandlerTestSuite) TestHandleWebsites() {

	t := suite.T()

	t.Run("when get all website is called successfully", func(t *testing.T) {
		r := httptest.NewRequest(http.MethodGet, "/website", nil)
		w := httptest.NewRecorder()

		inputUrl := "www.google.com"
		statuschecker.WebsiteList[inputUrl] = "Unknown"
		res := "Unknown"
		expctedRes := map[string]string{
			"www.google.com": "Unknown",
		}

		// suite.service.On("GetWebsiteHandler", suite.service, w, r).Return()
		suite.service.On("Check", mock.Anything, inputUrl).Return(res, nil)

		statuschecker.HandleWebsites(suite.service)(w, r)
		gotRes := make(map[string]string)
		json.NewDecoder(w.Body).Decode(&gotRes)

		assert.Equal(t, expctedRes, gotRes)

	})

	// t.Run("when get one website is called successfully", func(t *testing.T) {
	// 	inputUrl := "www.twitter.com"
	// 	queryParams := url.Values{}
	// 	queryParams.Add("name", inputUrl)
	// 	getOneurl := fmt.Sprintf("/website?%s", queryParams.Encode())
	// 	r := httptest.NewRequest(http.MethodGet, getOneurl, nil)
	// 	w := httptest.NewRecorder()

	// 	statuschecker.WebsiteList[inputUrl] = "Unknown"
	// 	res := "Unknown"
	// 	expctedRes := map[string]string{
	// 		"www.twitter.com": "Unknown",
	// 	}

	// 	// suite.service.On("GetWebsiteHandler", suite.service, w, r).Return()
	// 	suite.service.On("Check", mock.Anything, inputUrl).Return(res, nil)

	// 	statuschecker.HandleWebsites(suite.service)(w, r)
	// 	gotRes := make(map[string]string)
	// 	json.NewDecoder(w.Body).Decode(&gotRes)

	// 	assert.Equal(t, expctedRes, gotRes)

	// })

	t.Run("when an invalid request is sent", func(t *testing.T) {
		r := httptest.NewRequest(http.MethodGet, "/websites", nil)
		w := httptest.NewRecorder()

		router := http.NewServeMux()

		router.HandleFunc("/website", statuschecker.HandleWebsites(suite.service))

		router.ServeHTTP(w, r)

		assert.Equal(t, http.StatusNotFound, w.Code)
	})
}
