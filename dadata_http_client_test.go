package dadata

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type fakeHandlerOK struct{ t *testing.T }

type fakeHandlerISE struct{ t *testing.T }

// ServeHTTP Fake http server http.Handler response OK
func (fh fakeHandlerOK) ServeHTTP(wr http.ResponseWriter, rq *http.Request) {
	wr.Header().Add("Content-Type", "application/json")
	wr.WriteHeader(http.StatusOK)
	wr.Write([]byte(`{}`))
	fh.t.Logf("Request for fake http server %q, response OK and empty json", rq.RequestURI)
}

// ServeHTTP Fake http server http.Handler response internal server error (RFC 7231, 6.6.1)
func (fh fakeHandlerISE) ServeHTTP(wr http.ResponseWriter, rq *http.Request) {
	wr.WriteHeader(http.StatusInternalServerError)
	wr.Write([]byte(http.StatusText(http.StatusInternalServerError)))
	fh.t.Logf("Request for fake http server %q, response 500", rq.RequestURI)
}

func TestHttpClientRequest(t *testing.T) {
	var err error
	var srvOK, srvISE *httptest.Server

	// Restore baseURL
	defer func() { baseURL = constBaseURL }()
	// Server response OK
	srvOK = httptest.NewServer(&fakeHandlerOK{t: t})
	baseURL = srvOK.URL + "/"
	_, err = NewDaData("", "").ProfileBalance()
	if err != nil {
		t.Fatalf("DaData request error:  %s", err)
	}
	// Server fail
	srvISE = httptest.NewServer(&fakeHandlerISE{t: t})
	baseURL = srvISE.URL + "/"
	_, err = NewDaData("", "").ProfileBalance()
	if err == nil {
		t.Fatalf("Incorrect work a http client. Web server response 500 (internal server error), but request completed without errors")
	}
}

func TestCleanNilResponseAndError(t *testing.T) {
	var err error
	var srvISE *httptest.Server
	var dadata *DaData
	var address []Address
	var phones []Phone
	var names []Name
	var emails []Email
	var birthdates []Birthdate
	var vehicles []Vehicle
	var passports []Passport

	// Restore baseURL
	defer func() { baseURL = constBaseURL }()
	srvISE = httptest.NewServer(&fakeHandlerISE{t: t})
	baseURL = srvISE.URL + "/"
	dadata = NewDaData("", "")
	// Test functions
	if address, err = dadata.CleanAddresses(""); address != nil && err != nil {
		t.Fatalf("Function CleanAddresses() returns not a nil response with an error: %s", err)
	}
	if phones, err = dadata.CleanPhones(""); phones != nil && err != nil {
		t.Fatalf("Function CleanPhones() returns not a nil response with an error: %s", err)
	}
	if names, err = dadata.CleanNames(""); names != nil && err != nil {
		t.Fatalf("Function CleanNames() returns not a nil response with an error: %s", err)
	}
	if emails, err = dadata.CleanEmails(""); emails != nil && err != nil {
		t.Fatalf("Function CleanEmails() returns not a nil response with an error: %s", err)
	}
	if birthdates, err = dadata.CleanBirthdates(""); birthdates != nil && err != nil {
		t.Fatalf("Function CleanBirthdates() returns not a nil response with an error: %s", err)
	}
	if vehicles, err = dadata.CleanVehicles(""); vehicles != nil && err != nil {
		t.Fatalf("Function CleanVehicles() returns not a nil response with an error: %s", err)
	}
	if passports, err = dadata.CleanPassports(""); passports != nil && err != nil {
		t.Fatalf("Function CleanPassports() returns not a nil response with an error: %s", err)
	}
}

func TestSuggestNilResponseAndError(t *testing.T) {
	var err error
	var srvISE *httptest.Server
	var dadata *DaData
	var address []ResponseAddress
	var names []ResponseName
	var banks []ResponseBank
	var partyes []ResponseParty
	var emails []ResponseEmail

	// Restore baseURL
	defer func() { baseURL = constBaseURL }()
	srvISE = httptest.NewServer(&fakeHandlerISE{t: t})
	baseURL = srvISE.URL + "/"
	dadata = NewDaData("", "")
	// Test functions
	if address, err = dadata.SuggestAddresses(SuggestRequestParams{}); address != nil && err != nil {
		t.Fatalf("Function SuggestAddresses() returns not a nil response with an error: %s", err)
	}
	if names, err = dadata.SuggestNames(SuggestRequestParams{}); names != nil && err != nil {
		t.Fatalf("Function SuggestNames() returns not a nil response with an error: %s", err)
	}
	if banks, err = dadata.SuggestBanks(SuggestRequestParams{}); banks != nil && err != nil {
		t.Fatalf("Function SuggestBanks() returns not a nil response with an error: %s", err)
	}
	if partyes, err = dadata.SuggestParties(SuggestRequestParams{}); partyes != nil && err != nil {
		t.Fatalf("Function SuggestParties() returns not a nil response with an error: %s", err)
	}
	if emails, err = dadata.SuggestEmails(SuggestRequestParams{}); emails != nil && err != nil {
		t.Fatalf("Function SuggestEmails() returns not a nil response with an error: %s", err)
	}
}

func TestDailyStatNilResponseAndError(t *testing.T) {
	var err error
	var srvISE *httptest.Server
	var dadata *DaData
	var stat *StatResponse

	// Restore baseURL
	defer func() { baseURL = constBaseURL }()
	srvISE = httptest.NewServer(&fakeHandlerISE{t: t})
	baseURL = srvISE.URL + "/"
	dadata = NewDaData("", "")
	// Test functions
	if stat, err = dadata.DailyStat(time.Now()); stat != nil && err != nil {
		t.Fatalf("Function DailyStat() returns not a nil response with an error: %s", err)
	}
}

func TestAddressByIDNilResponseAndError(t *testing.T) {
	var err error
	var srvISE *httptest.Server
	var dadata *DaData
	var address *ResponseAddress
	var addresses []ResponseAddress

	// Restore baseURL
	defer func() { baseURL = constBaseURL }()
	srvISE = httptest.NewServer(&fakeHandlerISE{t: t})
	baseURL = srvISE.URL + "/"
	dadata = NewDaData("", "")
	// Test functions
	if address, err = dadata.AddressByID(""); address != nil && err != nil {
		t.Fatalf("Function AddressByID() returns not a nil response with an error: %s", err)
	}
	if addresses, err = dadata.AddressesByID(""); addresses != nil && err != nil {
		t.Fatalf("Function AddressesByID() returns not a nil response with an error: %s", err)
	}
}

func TestGeoIPNilResponseAndError(t *testing.T) {
	var err error
	var srvISE *httptest.Server
	var dadata *DaData
	var ip *GeoIPResponse

	// Restore baseURL
	defer func() { baseSuggestURL = constBaseSuggestURL }()
	srvISE = httptest.NewServer(&fakeHandlerISE{t: t})
	baseSuggestURL = srvISE.URL + "/"
	dadata = NewDaData("", "")
	// Test functions
	if ip, err = dadata.GeoIP("127.0.0.1"); ip != nil && err != nil {
		t.Fatalf("Function GeoIP() returns not a nil response with an error: %s", err)
	}
}
