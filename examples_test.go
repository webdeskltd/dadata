package dadata_test

import (
	"context"
	"fmt"
	"os"
	"time"

	. "gopkg.in/webdeskltd/dadata.v2"
)

func ExampleDaData_CleanAddresses() {
	daData := NewDaData(os.Getenv("API_KEY"), os.Getenv("SECRET_KEY"))

	addresses, err := daData.CleanAddresses("ул.Правды 26", "пер.Расковой 5")

	if nil != err {
		fmt.Println(err)
	}

	for _, address := range addresses {
		fmt.Println(address.StreetTypeFull)
		fmt.Println(address.Street)
		fmt.Println(address.House)
	}

	// Output:
	// улица
	// Правды
	// 26
	// переулок
	// Расковой
	// 5

}

func ExampleDaData_CleanNames() {
	daData := NewDaData(os.Getenv("API_KEY"), os.Getenv("SECRET_KEY"))

	names, err := daData.CleanNames("Алексей Иванов", "Иван Алексеев")
	if nil != err {
		fmt.Println(err)
	}

	for _, name := range names {
		fmt.Println(name.Surname)
		fmt.Println(name.Name)
	}

	// Output:
	// Иванов
	// Алексей
	// Алексеев
	// Иван

}

func ExampleDaData_SuggestAddresses() {
	daData := NewDaData(os.Getenv("API_KEY"), os.Getenv("SECRET_KEY"))

	addresses, err := daData.SuggestAddresses(SuggestRequestParams{Query: "Преснен", Count: 2})
	if nil != err {
		fmt.Println(err)
	}

	for _, address := range addresses {
		fmt.Println(address.UnrestrictedValue)
		fmt.Println(address.Data.Street)
		fmt.Println(address.Data.FiasLevel)
	}

	// Output:
	// г Москва, Пресненская наб
	// Пресненская
	// 7
	// г Москва, ул Пресненский Вал
	// Пресненский Вал
	// 7

}

func ExampleDaData_SuggestBanks() {
	daData := NewDaData(os.Getenv("API_KEY"), os.Getenv("SECRET_KEY"))

	banks, err := daData.SuggestBanks(SuggestRequestParams{Query: "Кредитный", Count: 3})
	if nil != err {
		fmt.Println(err)
	}

	for _, bank := range banks {
		fmt.Println(bank.Data.Name.Full)
		fmt.Println(bank.Data.Bic)
	}

	// Output:
	// "МОСКОВСКИЙ КРЕДИТНЫЙ БАНК" (ПУБЛИЧНОЕ АКЦИОНЕРНОЕ ОБЩЕСТВО)
	// 044525659
	// КОММЕРЧЕСКИЙ БАНК "РЕСПУБЛИКАНСКИЙ КРЕДИТНЫЙ АЛЬЯНС" (ОБЩЕСТВО С ОГРАНИЧЕННОЙ ОТВЕТСТВЕННОСТЬЮ)
	// 044525860
	// ЖИЛИЩНО-КРЕДИТНЫЙ КОММЕРЧЕСКИЙ БАНК "ЖИЛКРЕДИТ" ОБЩЕСТВО С ОГРАНИЧЕННОЙ ОТВЕТСТВЕННОСТЬЮ
	// 044525325
}

func ExampleDaData_SuggestParties() {
	daData := NewDaData(os.Getenv("API_KEY"), os.Getenv("SECRET_KEY"))

	parties, err := daData.SuggestParties(SuggestRequestParams{Query: "Агрохолд", Count: 3})
	if nil != err {
		fmt.Println(err)
	}

	for _, party := range parties {
		fmt.Println(party.Data.Name.Full)
		fmt.Println(party.Data.Ogrn)
	}

	// Output:
	// АГРОХОЛДИНГ
	// 1057746753327
	// АГРОХОЛДИНГ НОВОТЕХ
	// 5087746090042
	// МАГНУС АГРОХОЛДИНГ
	// 1133123023186
}

func ExampleDaData_AddressByID() {
	daData := NewDaData(os.Getenv("API_KEY"), os.Getenv("SECRET_KEY"))
	addr, err := daData.AddressByID("6300000100000")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("City: %s,\nFiasID: %s,\nKladr: %s\n", addr.Data.City, addr.Data.FiasID, addr.Data.KladrID)

	// Output:
	// City: Самара,
	// FiasID: bb035cc3-1dc2-4627-9d25-a1bf2d4b936b,
	// Kladr: 6300000100000
}

func ExampleDaData_GeoIP() {
	daData := NewDaData(os.Getenv("API_KEY"), os.Getenv("SECRET_KEY"))

	geoIPResponse, err := daData.GeoIP("83.220.54.223")
	if nil != err {
		fmt.Println(err)
		return
	}
	if geoIPResponse.Location == nil {
		fmt.Println("empty result from GeoIP")
		return
	}
	address := geoIPResponse.Location.Data
	fmt.Println(address.Country)
	fmt.Println(address.City)
	fmt.Printf("see on https://www.google.com/maps/@%s,%sf,14z\n", address.GeoLat, address.GeoLon)

	// Output:
	// Россия
	// Москва
	// see on https://www.google.com/maps/@55.7540471,37.620405f,14z
}

func ExampleDaData_SuggestAddresses_granular() {
	daData := NewDaData(os.Getenv("API_KEY"), os.Getenv("SECRET_KEY"))

	var req SuggestRequestParams

	req.Query = "лен"

	req.Locations = append(req.Locations, SuggestRequestParamsLocation{
		RegionFiasID: "df3d7359-afa9-4aaa-8ff9-197e73906b1c",
		CityFiasID:   "e9e684ce-7d60-4480-ba14-ca6da658188b",
	})

	req.FromBound = SuggestBound{SuggestBoundStreet}
	req.ToBound = SuggestBound{SuggestBoundStreet}

	req.RestrictValue = true
	req.Count = 2

	addresses, err := daData.SuggestAddresses(req)
	if nil != err {
		fmt.Println(err)
	}

	for _, address := range addresses {
		fmt.Println(address.UnrestrictedValue)
		fmt.Println(address.Data.Street)
	}

	// Output:
	// Самарская обл, г Сызрань, ул Ленина
	// Ленина
	// Самарская обл, г Сызрань, ул Ленинградская
	// Ленинградская
}

func ExampleDaData_SuggestAddressesWithCtx() {
	daData := NewDaData(os.Getenv("API_KEY"), os.Getenv("SECRET_KEY"))

	var req SuggestRequestParams

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)

	req.Query = "лен"

	req.Locations = append(req.Locations, SuggestRequestParamsLocation{
		RegionFiasID: "df3d7359-afa9-4aaa-8ff9-197e73906b1c",
		CityFiasID:   "e9e684ce-7d60-4480-ba14-ca6da658188b",
	})

	req.FromBound = SuggestBound{SuggestBoundStreet}
	req.ToBound = SuggestBound{SuggestBoundStreet}

	req.RestrictValue = true
	req.Count = 2

	addresses, err := daData.SuggestAddressesWithCtx(ctx, req)
	if nil != err {
		fmt.Println(err)
	}

	for _, address := range addresses {
		fmt.Println(address.UnrestrictedValue)
	}

	cancel()
	// if ctx is exited (by cancel or timeout) we must catch err
	_, err = daData.SuggestAddressesWithCtx(ctx, req)
	fmt.Println(err)

	// Output:
	// Самарская обл, г Сызрань, ул Ленина
	// Самарская обл, г Сызрань, ул Ленинградская
	// sendRequestToURL: ctx.Err return err=context canceled
}
