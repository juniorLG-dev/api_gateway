package api

import (
	"net/http"
	"bytes"
	"io/ioutil"
)

type api struct {
	method 		string
	apiURL 		string
	bodyBytes []byte
}

func NewAPIReq(
	method 		string,
	apiURL 		string,
	bodyBytes []byte,
) *api {
	return &api{
		method: method,
		apiURL: apiURL,
		bodyBytes: bodyBytes,
	}
}

type ResponseInfoOutput struct {
	Response string
	Status 	 int
	Header 	 string
}

func (a *api) SendRequest() (ResponseInfoOutput, error) {
	req, err := http.NewRequest(a.method, a.apiURL, bytes.NewBuffer(a.bodyBytes))
	if err != nil {
		return ResponseInfoOutput{}, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return ResponseInfoOutput{}, err
	}

	backendResponseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return ResponseInfoOutput{}, err
	}

	return ResponseInfoOutput{
		Response: string(backendResponseBody),
		Status: resp.StatusCode,
		Header: resp.Header.Get("Content-Type"),
	}, nil
}