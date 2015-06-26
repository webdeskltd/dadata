package dadata_test

import (
	"fmt"
	"os"

	// "github.com/probiizonka/dadata"
	dadata "."
)

func ExampleDaData_CleanAddresses() {
	daData := dadata.NewDaData(os.Getenv("API_KEY"), os.Getenv("SECRET_KEY"))

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
	daData := dadata.NewDaData(os.Getenv("API_KEY"), os.Getenv("SECRET_KEY"))

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
