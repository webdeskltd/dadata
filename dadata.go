// Golang client library for DaData.ru (https://dadata.ru/).

// Package dadata implemented only cleaning (https://dadata.ru/api/clean/) and suggesting (https://dadata.ru/api/suggest/)
package dadata

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

const baseURL = "https://dadata.ru/api/v2/"

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

func (daData *DaData) sendRequest(lastURLPart string, source interface{}, result interface{}) error {
	buffer := &bytes.Buffer{}

	if encodeErr := json.NewEncoder(buffer).Encode(source); nil != encodeErr {
		fmt.Printf("encodeErr: %v", encodeErr)
		return encodeErr
	}

	request, requestErr := http.NewRequest("POST", baseURL+lastURLPart, buffer)
	if nil != requestErr {
		fmt.Printf("requestErr: %v", requestErr)
		return requestErr
	}

	request.Header.Add("Authorization", fmt.Sprintf("Token %s", daData.apiKey))
	request.Header.Add("X-Secret", daData.secretKey)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Accept", "application/json")
	request.Header.Set("Connection", "close")

	response, httpErr := daData.httpClient.Do(request)
	if nil != httpErr {
		fmt.Printf("httpErr: %v", httpErr)
		return httpErr
	}

	defer response.Body.Close()

	if http.StatusOK != response.StatusCode {
		return fmt.Errorf("Request error %v", response.Status)
	}

	if decodeErr := json.NewDecoder(response.Body).Decode(&result); nil != decodeErr {
		fmt.Printf("decodeErr: %v", decodeErr)
		return decodeErr
	}

	return nil

}
