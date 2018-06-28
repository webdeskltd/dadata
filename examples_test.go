package dadata_test

import (
	"fmt"
	"os"

	. "github.com/mekegi/dadata"
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
