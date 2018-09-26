// Golang client library for DaData.ru (https://dadata.ru/).

// Package dadata implemented only cleaning (https://dadata.ru/api/clean/) and suggesting (https://dadata.ru/api/suggest/)
package dadata

import (
	"bytes"
	"context"
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

func (daData *DaData) sendRequestToURL(ctx context.Context, method, url string, source interface{}, result interface{}) error {
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("sendRequestToURL: ctx.Err return err=%v", err)
	}

	buffer := &bytes.Buffer{}

	if err := json.NewEncoder(buffer).Encode(source); err != nil {
		return fmt.Errorf("sendRequestToURL: json.Encode return err = %v", err)
	}

	request, err := http.NewRequest(method, url, buffer)

	if err != nil {
		return fmt.Errorf("sendRequestToURL: http.NewRequest return err = %v", err)
	}

	request = request.WithContext(ctx)

	request.Header.Add("Authorization", fmt.Sprintf("Token %s", daData.apiKey))
	request.Header.Add("X-Secret", daData.secretKey)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Accept", "application/json")
	request.Header.Set("Connection", "close")

	response, err := daData.httpClient.Do(request)
	if err != nil {
		return fmt.Errorf("sendRequestToURL: httpClient.Do return err = %v", err)
	}

	defer response.Body.Close()

	if http.StatusOK != response.StatusCode {
		return fmt.Errorf("sendRequestToURL: Request error %v", response.Status)
	}

	if err := json.NewDecoder(response.Body).Decode(&result); err != nil {
		return fmt.Errorf("sendRequestToURL: json.Decode return err = %v", err)
	}

	return nil
}

// sendRequest
func (daData *DaData) sendRequest(ctx context.Context, lastURLPart string, source interface{}, result interface{}) error {
	return daData.sendRequestToURL(ctx, "POST", baseURL+lastURLPart, source, result)
}
