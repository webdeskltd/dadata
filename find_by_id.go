package dadata

import (
	"context"
	"fmt"
)

// AddressByID find address by Fias or Kladr
// see full documentation https://confluence.hflabs.ru/pages/viewpage.action?pageId=312016944
func (daData *DaData) AddressByID(id string) (*ResponseAddress, error) {
	return daData.AddressByIDWithCtx(context.Background(), id)
}

// AddressByIDWithCtx find address by Fias or Kladr
// see full documentation https://confluence.hflabs.ru/pages/viewpage.action?pageId=312016944
func (daData *DaData) AddressByIDWithCtx(ctx context.Context, id string) (*ResponseAddress, error) {
	result := SuggestAddressResponse{}
	req := SuggestRequestParams{Query: id}
	if err := daData.sendRequestToURL(ctx, "POST", baseSuggestURL+"findById/address", req, &result); err != nil {
		return nil, err
	}

	if len(result.Suggestions) == 0 {
		return nil, fmt.Errorf("dadata.AddressByID: cannot detect address by id %v", id)
	}

	return &result.Suggestions[0], nil
}
