/*
Unofficial client for DaData.ru (https://dadata.ru/).

While implemented only cleaning (https://dadata.ru/api/clean/).
*/
package dadata

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

const BASE_URL = "https://dadata.ru/api/v2/"

type DaData struct {
	ApiKey     string
	SecretKey  string
	httpClient *http.Client
}

/*
Create new client of DaData.

Api and secret keys see on profile page (https://dadata.ru/profile/).
*/
func NewDaData(apiKey, secretKey string) *DaData {
	return NewDaDataCustomClient(apiKey, secretKey, &http.Client{})
}

/*
Create new custom client of DaData. By example, this option should be used to Google AppEngine:
    ctx := appengine.NewContext(request)
    appEngineClient := urlfetch.Client(ctx)
    daData:= NewDaDataCustomClient(apiKey, secretKey, client)
*/
func NewDaDataCustomClient(apiKey, secretKey string, httpClient *http.Client) *DaData {
	return &DaData{
		ApiKey:     apiKey,
		SecretKey:  secretKey,
		httpClient: httpClient,
	}
}

func (daData *DaData) sendRequest(lastUrlPart string, source interface{}, result interface{}) error {
	buffer := &bytes.Buffer{}

	if encodeErr := json.NewEncoder(buffer).Encode(source); nil != encodeErr {
		fmt.Printf("encodeErr: %v", encodeErr)
		return encodeErr
	}

	request, requestErr := http.NewRequest("POST", BASE_URL+lastUrlPart, buffer)
	if nil != requestErr {
		fmt.Printf("requestErr: %v", requestErr)
		return requestErr
	}

	request.Header.Add("Authorization", fmt.Sprintf("Token %s", daData.ApiKey))
	request.Header.Add("X-Secret", daData.SecretKey)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Accept", "application/json")

	response, httpErr := daData.httpClient.Do(request)
	if nil != httpErr {
		fmt.Printf("httpErr: %v", httpErr)
		return httpErr
	}

	if http.StatusOK != response.StatusCode {
		return fmt.Errorf("Request error %v", response.Status)
	}

	if decodeErr := json.NewDecoder(response.Body).Decode(&result); nil != decodeErr {
		fmt.Printf("decodeErr: %v", decodeErr)
		return decodeErr
	}

	return nil

}
