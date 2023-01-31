package statuschecker_test

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
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
