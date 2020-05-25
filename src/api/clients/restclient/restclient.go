package restclient

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

var (
	enableMock bool = false
	mocks           = make(map[string]*Mock)
)

type Mock struct {
	Response   *http.Response
	Err        error
	Url        string
	HttpMethod string
}

func StartMockUp() {
	enableMock = true
}

func StopMockUp() {
	enableMock = false
}

func AddMocks(mock Mock) {
	mocks[getMockId(mock.Url, mock.HttpMethod)] = &mock
}

func getMockId(url string, httpMethod string) string {
	return fmt.Sprintf("%s_%s", url, httpMethod)
}

func FlushMocks() {
	mocks = make(map[string]*Mock)
}

func Post(url string, headers http.Header, body interface{}) (*http.Response, error) {

	if enableMock {

		mock := mocks[getMockId(url, http.MethodPost)]

		if mock == nil {
			return nil, errors.New("Cannot find mock for this url")
		}
		return mock.Response, mock.Err
	}

	//TODO(
	// STEP 1 - CONVERT BODY TO JSON
	// STEP 2 - CREATE NEW REQUEST
	// STEP 3 - SET HEADERS TO NEWLY CREATED REQUEST
	// STEP 4 - CREATE CLIENT AND DO THE REQUEST
	// )

	jsonBody, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(jsonBody))
	if err != nil {
		return nil, err
	}
	request.Header = headers

	client := http.Client{}
	return client.Do(request)
}
