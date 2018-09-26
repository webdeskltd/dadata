package dadata_test

import (
	"reflect"
	"testing"

	. "github.com/webdeskltd/dadata"
)

func TestDaData_SuggestAddresses(t *testing.T) {
	tt := []struct {
		name        string
		query       string
		wantStreets []string
		wantErr     bool
	}{
		{
			"long random chars",
			"rliurihhjbgfihjvdfdsadakhvf",
			[]string{},
			false,
		},
		{
			"popular address - строите",
			"строите",
			[]string{"Строителей", "Строительный"},
			false,
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			daData := newSuggester()
			got, err := daData.SuggestAddresses(SuggestRequestParams{Query: tc.query, Count: 2})
			if (err != nil) != tc.wantErr {
				t.Errorf("DaData.SuggestAddresses() error = %v, wantErr %v", err, tc.wantErr)
				return
			}
			streets := []string{}
			for i := range got {
				streets = append(streets, got[i].Data.Street)
			}
			if !reflect.DeepEqual(streets, tc.wantStreets) {
				t.Errorf("DaData.SuggestAddresses() = %#v, want %#v", streets, tc.wantStreets)
			}
		})
	}
}

func TestDaData_SuggestAddressesWithConstarint(t *testing.T) {
	tt := []struct {
		name       string
		query      string
		wantRegion []string
		wantErr    bool
	}{
		{
			"long random chars",
			"rliurihhjbgfihjvdfdsadakhvf",
			[]string{},
			false,
		},
		{
			"popular address - ленина",
			"ленина",
			[]string{"Астраханская", "Астраханская"},
			false,
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			daData := newSuggester()
			req := SuggestRequestParams{
				Query:         tc.query,
				Count:         2,
				RestrictValue: true,
			}
			req.Locations = append(req.Locations, SuggestRequestParamsLocation{Region: "астраханская"})
			got, err := daData.SuggestAddresses(req)
			if (err != nil) != tc.wantErr {
				t.Errorf("DaData.SuggestAddresses() error = %v, wantErr %v", err, tc.wantErr)
				return
			}
			regions := []string{}
			for i := range got {
				regions = append(regions, got[i].Data.Region)
			}
			if !reflect.DeepEqual(regions, tc.wantRegion) {
				t.Errorf("DaData.SuggestAddresses() = %#v, want %#v", regions, tc.wantRegion)
			}
		})
	}
}

func TestDaData_SuggestNames(t *testing.T) {
	tt := []struct {
		name      string
		query     string
		wantNames []string
		wantErr   bool
	}{
		{
			"rare name",
			"Юст",
			[]string{"Юстус", "Юст"},
			false,
		},
		{
			"popular - алекс",
			"алекс",
			[]string{"Александр", "Алексей"},
			false,
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			daData := newSuggester()
			got, err := daData.SuggestNames(SuggestRequestParams{Query: tc.query, Count: 2})
			if (err != nil) != tc.wantErr {
				t.Errorf("DaData.SuggestNames() error = %v, wantErr %v", err, tc.wantErr)
				return
			}
			names := []string{}
			for i := range got {
				names = append(names, got[i].UnrestrictedValue)
			}
			if !reflect.DeepEqual(names, tc.wantNames) {
				t.Errorf("DaData.SuggestParties() = %#v, want %#v", names, tc.wantNames)
			}
		})
	}
}

func TestDaData_SuggestBanks(t *testing.T) {
	tt := []struct {
		name      string
		query     string
		wantNames []string
		wantErr   bool
	}{
		{
			"long random chars",
			"rliurihhjbgfihjvdfdsadakhvf",
			[]string{},
			false,
		},
		{
			"popular - сбер",
			"сберба",
			[]string{"ПАО СБЕРБАНК", "БАЙКАЛЬСКИЙ БАНК ПАО СБЕРБАНК"},
			false,
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			daData := newSuggester()
			got, err := daData.SuggestBanks(SuggestRequestParams{Query: tc.query, Count: 2})
			if (err != nil) != tc.wantErr {
				t.Errorf("DaData.SuggestBanks() error = %v, wantErr %v", err, tc.wantErr)
				return
			}
			names := []string{}
			for i := range got {
				names = append(names, got[i].UnrestrictedValue)
			}
			if !reflect.DeepEqual(names, tc.wantNames) {
				t.Errorf("DaData.SuggestParties() = %#v, want %#v", names, tc.wantNames)
			}
		})
	}
}

func TestDaData_SuggestParties(t *testing.T) {
	tt := []struct {
		name      string
		query     string
		wantNames []string
		wantErr   bool
	}{
		{
			"long random chars",
			"rliurihhjbgfihjvdfdsadakhvf",
			[]string{},
			false,
		},
		{
			"popular - хлебзавод",
			"хлебзавод",
			[]string{`ООО "ХЛЕБЗАВОД"`, `ООО "ХЛЕБЗАВОД"`},
			false,
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			daData := newSuggester()
			got, err := daData.SuggestParties(SuggestRequestParams{Query: tc.query, Count: 2})
			if (err != nil) != tc.wantErr {
				t.Errorf("DaData.SuggestParties() error = %v, wantErr %v", err, tc.wantErr)
				return
			}
			names := []string{}
			for i := range got {
				names = append(names, got[i].UnrestrictedValue)
			}
			if !reflect.DeepEqual(names, tc.wantNames) {
				t.Errorf("DaData.SuggestParties() = %#v, want %#v", names, tc.wantNames)
			}
		})
	}
}

func TestDaData_SuggestEmails(t *testing.T) {
	tt := []struct {
		name    string
		query   string
		want    []ResponseEmail
		wantErr bool
	}{
		{
			"long random chars",
			"rliurihhjbgfihjvdfdsadakhvf",
			[]ResponseEmail{},
			false,
		},
		{
			"short random chars",
			"dktrigh",
			[]ResponseEmail{},
			false,
		},
		{
			"popular name - konstantin",
			"konstantin",
			[]ResponseEmail{
				{"konstantin", "konstantin", Email{}},
				{"konstantinov", "konstantinov", Email{}},
			},
			false,
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			daData := newSuggester()
			got, err := daData.SuggestEmails(SuggestRequestParams{Query: tc.query, Count: 2})
			if (err != nil) != tc.wantErr {
				t.Errorf("DaData.SuggestEmails() error = %v, wantErr %v", err, tc.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("DaData.SuggestEmails() = %#v, want %#v", got, tc.want)
			}
		})
	}
}

func TestDaData_SuggestAddressesWithBounds(t *testing.T) {
	tt := []struct {
		query            string
		fromBound        BoundValue
		toBound          BoundValue
		constraintRegion string
		wantValues       []string
		wantErr          bool
	}{
		{
			"Дудин",
			SuggestBoundStreet,
			SuggestBoundStreet,
			"Москва",
			[]string{"ул Дудинка"},
			false,
		},
		{
			"Сызр",
			SuggestBoundCity,
			SuggestBoundCity,
			"Самарская",
			[]string{"г Сызрань"},
			false,
		},
	}
	for _, tc := range tt {
		t.Run(tc.query+" "+string(tc.fromBound), func(t *testing.T) {
			daData := newSuggester()
			req := SuggestRequestParams{
				Query:         tc.query,
				Count:         2,
				FromBound:     SuggestBound{tc.fromBound},
				ToBound:       SuggestBound{tc.toBound},
				RestrictValue: true,
			}
			req.Locations = append(req.Locations, SuggestRequestParamsLocation{Region: tc.constraintRegion})
			got, err := daData.SuggestAddresses(req)

			if (err != nil) != tc.wantErr {
				t.Errorf("DaData.SuggestAddresses() error = %v, wantErr %v", err, tc.wantErr)
				return
			}
			gotValues := []string{}
			for i := range got {
				gotValues = append(gotValues, got[i].Value)
			}
			if !reflect.DeepEqual(gotValues, tc.wantValues) {
				t.Errorf("DaData.SuggestAddresses() = %#v, want %#v", gotValues, tc.wantValues)
			}
		})
	}
}

func TestDaData_RegionSuggest(t *testing.T) {
	daData := newSuggester()
	req := SuggestRequestParams{
		Query:     "Санкт",
		FromBound: SuggestBound{Value: "region"},
		ToBound:   SuggestBound{Value: "region"},
	}

	got, err := daData.SuggestAddresses(req)
	if err != nil {
		t.Error(err)
	}

	if len(got) != 1 {
		t.Errorf("suggest by %v and bound 'region' return > 1 response", req.Query)
	}
}
