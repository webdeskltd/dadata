package dadata // import "gopkg.in/webdeskltd/dadata.v2"

import "context"

func (daData *DaData) sendSuggestRequest(ctx context.Context, lastURLPart string, requestParams SuggestRequestParams, result interface{}) error {
	return daData.sendRequest(ctx, "suggest/"+lastURLPart, requestParams, result)
}

// SuggestAddresses try to return suggest addresses by requestParams
func (daData *DaData) SuggestAddresses(requestParams SuggestRequestParams) ([]ResponseAddress, error) {
	return daData.SuggestAddressesWithCtx(context.Background(), requestParams)
}

// SuggestAddressesWithCtx try to return suggest addresses by requestParams
func (daData *DaData) SuggestAddressesWithCtx(ctx context.Context, requestParams SuggestRequestParams) (ret []ResponseAddress, err error) {
	var result = &SuggestAddressResponse{}
	if err = daData.sendSuggestRequest(ctx, "address", requestParams, result); err != nil {
		return
	}
	ret = result.Suggestions
	return
}

// SuggestNames try to return suggest names by requestParams
func (daData *DaData) SuggestNames(requestParams SuggestRequestParams) ([]ResponseName, error) {
	return daData.SuggestNamesWithCtx(context.Background(), requestParams)
}

// SuggestNamesWithCtx try to return suggest names by requestParams
func (daData *DaData) SuggestNamesWithCtx(ctx context.Context, requestParams SuggestRequestParams) (ret []ResponseName, err error) {
	var result = &SuggestNameResponse{}
	if err = daData.sendSuggestRequest(ctx, "fio", requestParams, result); err != nil {
		return
	}
	ret = result.Suggestions
	return
}

// SuggestBanks try to return suggest banks by requestParams
func (daData *DaData) SuggestBanks(requestParams SuggestRequestParams) ([]ResponseBank, error) {
	return daData.SuggestBanksWithCtx(context.Background(), requestParams)
}

// SuggestBanksWithCtx try to return suggest banks by requestParams
func (daData *DaData) SuggestBanksWithCtx(ctx context.Context, requestParams SuggestRequestParams) (ret []ResponseBank, err error) {
	var result = &SuggestBankResponse{}
	if err = daData.sendSuggestRequest(ctx, "bank", requestParams, result); err != nil {
		return
	}
	ret = result.Suggestions
	return
}

// SuggestParties try to return suggest parties by requestParams
func (daData *DaData) SuggestParties(requestParams SuggestRequestParams) ([]ResponseParty, error) {
	return daData.SuggestPartiesWithCtx(context.Background(), requestParams)
}

// SuggestPartiesWithCtx try to return suggest parties by requestParams
func (daData *DaData) SuggestPartiesWithCtx(ctx context.Context, requestParams SuggestRequestParams) (ret []ResponseParty, err error) {
	var result = &SuggestPartyResponse{}
	if err = daData.sendSuggestRequest(ctx, "party", requestParams, result); err != nil {
		return
	}
	ret = result.Suggestions
	return
}

// SuggestEmails try to return suggest emails by requestParams
func (daData *DaData) SuggestEmails(requestParams SuggestRequestParams) ([]ResponseEmail, error) {
	return daData.SuggestEmailsWithCtx(context.Background(), requestParams)
}

// SuggestEmailsWithCtx try to return suggest emails by requestParams
func (daData *DaData) SuggestEmailsWithCtx(ctx context.Context, requestParams SuggestRequestParams) (ret []ResponseEmail, err error) {
	var result = &SuggestEmailResponse{}
	if err = daData.sendSuggestRequest(ctx, "email", requestParams, result); err != nil {
		return
	}
	ret = result.Suggestions
	return
}
