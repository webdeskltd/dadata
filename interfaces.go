package dadata

import (
	"context"
	"time"
)

// AddressesCleaner is the interface for cleaning Addresses
type AddressesCleaner interface {
	CleanAddresses(addresses ...string) ([]Address, error)
}

// PhonesCleaner is the interface for cleaning phones
type PhonesCleaner interface {
	CleanPhones(phones ...string) ([]Phone, error)
}

// NamesCleaner is the interface for cleaning Names
type NamesCleaner interface {
	CleanNames(names ...string) ([]Name, error)
}

// EmailsCleaner is the interface for cleaning Emails
type EmailsCleaner interface {
	CleanEmails(emails ...string) ([]Email, error)
}

// BirthdatesCleaner is the interface for cleaning Birthdates
type BirthdatesCleaner interface {
	CleanBirthdates(birthdates ...string) ([]Birthdate, error)
}

// VehicleCleaner is the interface for cleaning Vehicle
type VehicleCleaner interface {
	CleanVehicles(vehicles ...string) ([]Vehicle, error)
}

// PassportCleaner is the interface for cleaning Passport
type PassportCleaner interface {
	CleanPassports(passports ...string) ([]Passport, error)
}

// Cleaner combine all xxxCleaner interfaces
// Stubs it for tests
type Cleaner interface {
	AddressesCleaner
	PhonesCleaner
	NamesCleaner
	EmailsCleaner
	BirthdatesCleaner
	VehicleCleaner
	PassportCleaner
}

// AddressSuggester is the interface for suggest Address
type AddressSuggester interface {
	SuggestAddresses(requestParams SuggestRequestParams) ([]ResponseAddress, error)
}

// NamesSuggester is the interface for suggest Names
type NamesSuggester interface {
	SuggestNames(requestParams SuggestRequestParams) ([]ResponseName, error)
}

// EmailsSuggester is the interface for suggest Emails
type EmailsSuggester interface {
	SuggestEmails(requestParams SuggestRequestParams) ([]ResponseEmail, error)
}

// BanksSuggester is the interface for suggest Banks
type BanksSuggester interface {
	SuggestBanks(requestParams SuggestRequestParams) ([]ResponseBank, error)
}

// PartiesSuggester is the interface for suggest Parties
type PartiesSuggester interface {
	SuggestParties(requestParams SuggestRequestParams) ([]ResponseParty, error)
}

// Suggester combine all xxxSuggester interfaces
// Stubs it for tests
type Suggester interface {
	AddressSuggester
	BanksSuggester
	EmailsSuggester
	NamesSuggester
	PartiesSuggester
}

// GeoIPDetector is the interface for detect address by client IP
type GeoIPDetector interface {
	GeoIP(ip string) (*GeoIPResponse, error)
}

// ByIDFinder interface for return data by id
type ByIDFinder interface {
	AddressByID(id string) (*ResponseAddress, error)
}

// Stater interface for return daily statistics
type Stater interface {
	DailyStat(ctx context.Context, date time.Time) (*StatResponse, error)
}

type Balancer interface {
	ProfileBalance(ctx context.Context) (*BalanceResponse, error)
}
