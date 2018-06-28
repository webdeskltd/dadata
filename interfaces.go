package dadata

type AddressesCleaner interface {
	CleanAddresses(addresses ...string) ([]Address, error)
}

type PhonesCleaner interface {
	CleanPhones(phones ...string) ([]Phone, error)
}

type NamesCleaner interface {
	CleanNames(names ...string) ([]Name, error)
}

type EmailsCleaner interface {
	CleanEmails(emails ...string) ([]Email, error)
}

type BirthdatesCleaner interface {
	CleanBirthdates(birthdates ...string) ([]Birthdate, error)
}

type VehicleCleaner interface {
	CleanVehicles(vehicles ...string) ([]Vehicle, error)
}

type PassportCleaner interface {
	CleanPassports(passports ...string) ([]Passport, error)
}

/*
Public interface. Stubs it for tests.
*/
type Cleaner interface {
	AddressesCleaner
	PhonesCleaner
	NamesCleaner
	EmailsCleaner
	BirthdatesCleaner
	VehicleCleaner
	PassportCleaner
}

type AddressSuggester interface {
	SuggestAddresses(requestParams SuggestRequestParams) ([]ResponseAddress, error)
}

type NamesSuggester interface {
	SuggestNames(requestParams SuggestRequestParams) ([]ResponseName, error)
}

type EmailsSuggester interface {
	SuggestEmails(requestParams SuggestRequestParams) ([]ResponseEmail, error)
}

type BanksSuggester interface {
	SuggestBanks(requestParams SuggestRequestParams) ([]ResponseBank, error)
}

type PartiesSuggester interface {
	SuggestParties(requestParams SuggestRequestParams) ([]ResponseParty, error)
}

type Suggester interface {
	AddressSuggester
	BanksSuggester
	EmailsSuggester
	NamesSuggester
	PartiesSuggester
}
