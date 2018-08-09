package dadata

import "fmt"

const baseSuggestURL = "https://suggestions.dadata.ru/suggestions/api/4_1/rs/"

// GeoIPRequestParams Request struct
type GeoIPRequestParams struct {
	IP string `json:"ip"`
}

type GeoIPResponse struct {
	Location *ResponseAddress `json:"location"`
}

// GeoIP try to return suggest parties by requestParams
func (daData *DaData) GeoIP(ip string) (*GeoIPResponse, error) {

	var result GeoIPResponse

	if err := daData.sendRequestToURL("GET", baseSuggestURL+"detectAddressByIp?ip="+ip, nil, &result); err != nil {
		return nil, err
	}
	if result.Location == nil {
		return nil, fmt.Errorf("Cannot detect address by ip %v", ip)
	}
	return &result, nil
}
