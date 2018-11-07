package dadata_test

import (
	"testing"

	"github.com/go-test/deep"
	. "github.com/webdeskltd/dadata/v2"
)

func TestCleanAddresses(t *testing.T) {
	var addressesCleaner AddressesCleaner = newCleaner()

	sourceAddresses := []string{
		"ул.Правды 26",
		"ул.Расковой 25",
	}

	addresses, err := addressesCleaner.CleanAddresses(sourceAddresses...)

	if err != nil {
		t.Errorf("CleanAddresses return error: %v", err)
	}

	if len(addresses) != len(sourceAddresses) {
		t.Errorf("CleanAddresses return %d addresses, but expect %d", len(addresses), len(sourceAddresses))
	}

	emptyAddress := Address{}
	for index, address := range addresses {
		if len(deep.Equal(address, emptyAddress)) == 0 {
			t.Errorf(`CleanAddresses return empty address for "%s"`, sourceAddresses[index])
		}
	}
}

func TestCleanPhones(t *testing.T) {
	var phonesCleaner PhonesCleaner = newCleaner()

	sourcePhones := []string{
		"+7 499 123 45 67",
		"+7(812)123-45-67",
	}

	phones, err := phonesCleaner.CleanPhones(sourcePhones...)

	if nil != err {
		t.Errorf(`CleanPhones return error "%s"`, err)
	}

	if len(sourcePhones) != len(phones) {
		t.Errorf(`CleanPhones return %d phones, but expect %d`, len(phones), len(sourcePhones))
	}

	emptyPhone := Phone{}
	for index, phone := range phones {
		if phone == emptyPhone {
			t.Errorf(`CleanPhones return empty phone %v for "%s"`, phone, sourcePhones[index])
		}
	}
}

func TestCleanNames(t *testing.T) {
	var namesCleaner NamesCleaner = newCleaner()

	sourceNames := []string{
		"Иван Алексеев",
		"Алексей Иванов",
	}

	names, err := namesCleaner.CleanNames(sourceNames...)

	if nil != err {
		t.Errorf(`CleanNames return error: %s`, err.Error())
	}

	if len(sourceNames) != len(names) {
		t.Errorf(`CleanNames return %d names, but expect %d`, len(names), len(sourceNames))
	}

	emptyName := Name{}
	for index, name := range names {
		if name == emptyName {
			t.Errorf(`CleanPones return empty name %v for "%s"`, name, sourceNames[index])
		}
	}
}

func TestCleanEmails(t *testing.T) {
	var emailsCleaner EmailsCleaner = newCleaner()
	sourceEmails := []string{
		"test@test/ru",
		"test@test.ry",
	}

	emails, err := emailsCleaner.CleanEmails(sourceEmails...)

	if nil != err {
		t.Errorf(`CleanEmails return error: %v`, err)
	}

	if len(emails) != len(sourceEmails) {
		t.Errorf(`CleanEmails return %d emails, but expect %d`, len(emails), len(sourceEmails))
	}

	emptyEmail := Email{}
	for index, email := range emails {
		if email == emptyEmail {
			t.Errorf(`CleanEmails return empty email %v for "%s"`, email, sourceEmails[index])
		}
	}
}

func TestCleanBirthdates(t *testing.T) {
	var birthdatesCleaner BirthdatesCleaner = newCleaner()

	sourceBirthdates := []string{
		"24/3/12",
		"3/24/12",
	}

	birthdates, err := birthdatesCleaner.CleanBirthdates(sourceBirthdates...)

	if nil != err {
		t.Errorf(`CleanBirthdates return error: %v`, err)
	}

	if len(sourceBirthdates) != len(birthdates) {
		t.Errorf(`CleanBirthdates return %d birthdates, but expected %d`, len(birthdates), len(sourceBirthdates))
	}

	emptyBirthdate := Birthdate{}
	for index, birthdate := range birthdates {
		if birthdate == emptyBirthdate {
			t.Errorf(`CleanBirthdates return empty birthdate %v for "%s"`, birthdate, sourceBirthdates[index])
		}
	}
}

func TestCleanVehicles(t *testing.T) {
	var vehiclesCleaner VehicleCleaner = newCleaner()

	sourceVehicles := []string{
		"форд фокус",
		"тойота королла",
	}

	vehicles, err := vehiclesCleaner.CleanVehicles(sourceVehicles...)

	if nil != err {
		t.Errorf(`CleanVehicles return error: %v`, err)
	}

	if len(sourceVehicles) != len(vehicles) {
		t.Errorf(`CleanVehicles return %d vehicles, but expected %d`, len(vehicles), len(sourceVehicles))
	}

	emptyVehicle := Vehicle{}
	for index, vehicle := range vehicles {
		if vehicle == emptyVehicle {
			t.Errorf(`CleanVehicles return empty vehicle %v for "%s"`, vehicle, sourceVehicles[index])
		}
	}
}

func TestCleanPassports(t *testing.T) {
	var passportsCleaner PassportCleaner = newCleaner()

	sourcePassports := []string{
		"4519 235347",
		"4509 235857",
	}

	passports, err := passportsCleaner.CleanPassports(sourcePassports...)

	if nil != err {
		t.Errorf(`CleanPassports return error: %v`, err)
	}

	if len(sourcePassports) != len(passports) {
		t.Errorf(`CleanPassports return %d passports, but expected %d`, len(passports), len(sourcePassports))
	}

	emptyPassport := Passport{}
	for index, passport := range passports {
		if passport == emptyPassport {
			t.Errorf(`CleanPassports return empty passport %v for "%s"`, passport, sourcePassports[index])
		}
	}
}
