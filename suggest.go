package dadata

// SuggestRequestParamsLocation constraints for suggestion
// full documentation https://confluence.hflabs.ru/pages/viewpage.action?pageId=204669108
type SuggestRequestParamsLocation struct {
	FiasID               string `json:"fias_id,omitempty"`
	KladrID              string `json:"kladr_id,omitempty"`
	Region               string `json:"region,omitempty"`
	RegionFiasID         string `json:"region_fias_id,omitempty"`
	RegionKladrID        string `json:"region_kladr_id,omitempty"`
	RegionTypeFull       string `json:"region_type_full,omitempty"`
	City                 string `json:"city,omitempty"`
	CityFiasID           string `json:"city_fias_id,omitempty"` // search only in this area
	CityKladrID          string `json:"city_kladr_id,omitempty"`
	CityTypeFull         string `json:"city_type_full,omitempty"`
	CityDistrictTypeFull string `json:"city_district_type_full,omitempty"`
	Settlement           string `json:"settlement,omitempty"`
	SettlementFiasID     string `json:"settlement_fias_id,omitempty"`
	SettlementKladrID    string `json:"settlement_kladr_id,omitempty"`
	SettlementTypeFull   string `json:"settlement_type_full,omitempty"`
	Street               string `json:"street,omitempty"`
	StreetFiasID         string `json:"street_fias_id,omitempty"`
	StreetKladrID        string `json:"street_kladr_id,omitempty"`
	StreetTypeFull       string `json:"street_type_full,omitempty"`
	AreaTypeFull         string `json:"area_type_full,omitempty"`
}

// SuggestBound for granular sugestion
// full documentation https://confluence.hflabs.ru/pages/viewpage.action?pageId=222888017
type SuggestBound struct {
	Value BoundValue
}

// SuggestRequestParams Request struct
type SuggestRequestParams struct {
	Query         string                         `json:"query"` // user input for suggestion
	Count         int                            `json:"count"` // ligmit for results
	Locations     []SuggestRequestParamsLocation `json:"locations"`
	RestrictValue bool                           `json:"restrict_value"` // don't show restricts (region) on results

	FromBound SuggestBound `json:"from_bound"`
	ToBound   SuggestBound `json:"to_bound"`
}

// SuggestAddressResponse result slice for address suggestions
type SuggestAddressResponse struct {
	Suggestions []ResponseAddress `json:"suggestions"`
}

// SuggestNameResponse result slice for name suggestions
type SuggestNameResponse struct {
	Suggestions []ResponseName `json:"suggestions"`
}

// SuggestBankResponse result slice for bank suggestions
type SuggestBankResponse struct {
	Suggestions []ResponseBank `json:"suggestions"`
}

// SuggestPartyResponse result slice for party suggestions
type SuggestPartyResponse struct {
	Suggestions []ResponseParty `json:"suggestions"`
}

// SuggestEmailResponse result slice for email suggestions
type SuggestEmailResponse struct {
	Suggestions []ResponseEmail `json:"suggestions"`
}

func (daData *DaData) sendSuggestRequest(lastURLPart string, requestParams SuggestRequestParams, result interface{}) error {
	return daData.sendRequest("suggest/"+lastURLPart, requestParams, result)
}

// SuggestAddresses try to return suggest addresses by requestParams
func (daData *DaData) SuggestAddresses(requestParams SuggestRequestParams) ([]ResponseAddress, error) {

	result := SuggestAddressResponse{}
	if err := daData.sendSuggestRequest("address", requestParams, &result); err != nil {
		return nil, err
	}

	return result.Suggestions, nil
}

// SuggestNames try to return suggest names by requestParams
func (daData *DaData) SuggestNames(requestParams SuggestRequestParams) ([]ResponseName, error) {

	result := SuggestNameResponse{}
	if err := daData.sendSuggestRequest("fio", requestParams, &result); err != nil {
		return nil, err
	}

	return result.Suggestions, nil
}

// SuggestBanks try to return suggest banks by requestParams
func (daData *DaData) SuggestBanks(requestParams SuggestRequestParams) ([]ResponseBank, error) {

	result := SuggestBankResponse{}
	if err := daData.sendSuggestRequest("bank", requestParams, &result); err != nil {
		return nil, err
	}

	return result.Suggestions, nil
}

// SuggestParties try to return suggest parties by requestParams
func (daData *DaData) SuggestParties(requestParams SuggestRequestParams) ([]ResponseParty, error) {

	result := SuggestPartyResponse{}
	if err := daData.sendSuggestRequest("party", requestParams, &result); err != nil {
		return nil, err
	}

	return result.Suggestions, nil
}

// SuggestEmails try to return suggest emails by requestParams
func (daData *DaData) SuggestEmails(requestParams SuggestRequestParams) ([]ResponseEmail, error) {

	result := SuggestEmailResponse{}
	if err := daData.sendSuggestRequest("email", requestParams, &result); err != nil {
		return nil, err
	}

	return result.Suggestions, nil
}
