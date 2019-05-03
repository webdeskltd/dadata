// Golang client library for DaData.ru (https://dadata.ru/).

// Package dadata implemented cleaning (https://dadata.ru/api/clean/) and suggesting (https://dadata.ru/api/suggest/)
package dadata // import "gopkg.in/webdeskltd/dadata.v2"

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const constBaseURL = "https://dadata.ru/api/v2/"

var baseURL = constBaseURL

// DaData client for DaData.ru (https://dadata.ru/)
type DaData struct {
	apiKey     string
	secretKey  string
	httpClient *http.Client
}

//NewDaData Create new client of DaData.
//Api and secret keys see on profile page (https://dadata.ru/profile/).
func NewDaData(apiKey, secretKey string) *DaData {
	return NewDaDataCustomClient(apiKey, secretKey, &http.Client{})
}

// NewDaDataCustomClient Create new custom client of DaData. By example, this option should be used to Google AppEngine:
//    ctx := appengine.NewContext(request)
//    appEngineClient := urlfetch.Client(ctx)
//    daData:= NewDaDataCustomClient(apiKey, secretKey, appEngineClient)
func NewDaDataCustomClient(apiKey, secretKey string, httpClient *http.Client) *DaData {
	return &DaData{
		apiKey:     apiKey,
		secretKey:  secretKey,
		httpClient: httpClient,
	}
}

func (daData *DaData) sendRequestToURL(ctx context.Context, method, url string, source interface{}, result interface{}) (err error) {
	var buffer *bytes.Buffer
	var request *http.Request
	var response *http.Response
	var buf *bytes.Buffer

	if err = ctx.Err(); err != nil {
		err = fmt.Errorf("sendRequestToURL: ctx.Err return error: %s", err)
		return
	}
	buffer = &bytes.Buffer{}
	if err = json.NewEncoder(buffer).Encode(source); err != nil {
		err = fmt.Errorf("sendRequestToURL: json.Encode return error: %s", err)
		return
	}
	request, err = http.NewRequest(method, url, buffer)
	if err != nil {
		err = fmt.Errorf("sendRequestToURL: http.NewRequest return error: %s", err)
		return
	}
	request = request.WithContext(ctx)
	request.Header.Add("Authorization", fmt.Sprintf("Token %s", daData.apiKey))
	request.Header.Add("X-Secret", daData.secretKey)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Accept", "application/json")
	response, err = daData.httpClient.Do(request)
	if err != nil {
		err = fmt.Errorf("sendRequestToURL: httpClient.Do return error: %s", err)
		return
	}
	defer func() { _ = response.Body.Close() }()

	if http.StatusOK != response.StatusCode {
		err = fmt.Errorf("sendRequestToURL: Request error: %s", response.Status)
		return
	}
	buf = &bytes.Buffer{}
	if _, err = io.Copy(buf, response.Body); err != nil {
		err = fmt.Errorf("sendRequestToURL: reading response body error: %s", err)
		return
	}
	if err = json.Unmarshal(buf.Bytes(), &result); err != nil {
		err = fmt.Errorf("sendRequestToURL: json.Decode return error: %s\nServer return:\n%s",
			err, buf.String(),
		)
		return
	}

	return
}

// sendRequest
func (daData *DaData) sendRequest(ctx context.Context, lastURLPart string, source interface{}, result interface{}) error {
	return daData.sendRequestToURL(ctx, "POST", baseURL+lastURLPart, source, result)
}
