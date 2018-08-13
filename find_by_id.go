package dadata

import "fmt"

// AddressByID find address by Fias or Kladr
// see full documentation https://confluence.hflabs.ru/pages/viewpage.action?pageId=312016944
func (daData *DaData) AddressByID(id string) (*ResponseAddress, error) {
	result := SuggestAddressResponse{}
	req := SuggestRequestParams{Query: id}
	if err := daData.sendRequestToURL("POST", baseSuggestURL+"findById/address", req, &result); err != nil {
		return nil, err
	}

	if len(result.Suggestions) == 0 {
		return nil, fmt.Errorf("dadata.AddressByID: cannot detect address by id %v", id)
	}

	return &result.Suggestions[0], nil
}
