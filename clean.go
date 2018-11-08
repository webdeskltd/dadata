package dadata // import "gopkg.in/webdeskltd/dadata.v2"

import "context"

func (daData *DaData) sendCleanRequest(ctx context.Context, lastURLPart string, source, result interface{}) error {
	return daData.sendRequest(ctx, "clean/"+lastURLPart, source, result)
}

// CleanAddresses clean all provided addresses
// Call https://dadata.ru/api/v2/clean/address
func (daData *DaData) CleanAddresses(sourceAddresses ...string) ([]Address, error) {
	return daData.CleanAddressesWithCtx(context.Background(), sourceAddresses...)
}

// CleanAddressesWithCtx clean all provided addresses
// Call https://dadata.ru/api/v2/clean/address
func (daData *DaData) CleanAddressesWithCtx(ctx context.Context, sourceAddresses ...string) (addresses []Address, err error) {
	if err = daData.sendCleanRequest(ctx, "address", &sourceAddresses, &addresses); err != nil {
		addresses = nil
	}
	return
}

// CleanPhones clean all provided phones
// Call https://dadata.ru/api/v2/clean/phone
func (daData *DaData) CleanPhones(sourcePhones ...string) ([]Phone, error) {
	return daData.CleanPhonesWithCtx(context.Background(), sourcePhones...)
}

// CleanPhonesWithCtx clean all provided phones
// Call https://dadata.ru/api/v2/clean/phone
func (daData *DaData) CleanPhonesWithCtx(ctx context.Context, sourcePhones ...string) (phones []Phone, err error) {
	if err = daData.sendCleanRequest(ctx, "phone", &sourcePhones, &phones); err != nil {
		phones = nil
	}
	return
}

// CleanNames clean all provided names
// Call https://dadata.ru/api/v2/clean/name
func (daData *DaData) CleanNames(sourceNames ...string) ([]Name, error) {
	return daData.CleanNamesWithCtx(context.Background(), sourceNames...)
}

// CleanNamesWithCtx clean all provided names
// Call https://dadata.ru/api/v2/clean/name
func (daData *DaData) CleanNamesWithCtx(ctx context.Context, sourceNames ...string) (names []Name, err error) {
	if err = daData.sendCleanRequest(ctx, "name", &sourceNames, &names); err != nil {
		names = nil
	}
	return
}

// CleanEmails clean all provided emails
// Call https://dadata.ru/api/v2/clean/email
func (daData *DaData) CleanEmails(sourceEmails ...string) ([]Email, error) {
	return daData.CleanEmailsWithCtx(context.Background(), sourceEmails...)
}

// CleanEmailsWithCtx clean all provided emails
// Call https://dadata.ru/api/v2/clean/email
func (daData *DaData) CleanEmailsWithCtx(ctx context.Context, sourceEmails ...string) (emails []Email, err error) {
	if err = daData.sendCleanRequest(ctx, "email", &sourceEmails, &emails); err != nil {
		emails = nil
	}
	return
}

// CleanBirthdates clean all provided birthdates
// Call https://dadata.ru/api/v2/clean/birthdate
func (daData *DaData) CleanBirthdates(sourceBirthdates ...string) ([]Birthdate, error) {
	return daData.CleanBirthdatesWithCtx(context.Background(), sourceBirthdates...)
}

// CleanBirthdatesWithCtx clean all provided birthdates
// Call https://dadata.ru/api/v2/clean/birthdate
func (daData *DaData) CleanBirthdatesWithCtx(ctx context.Context, sourceBirthdates ...string) (birthdates []Birthdate, err error) {
	if err = daData.sendCleanRequest(ctx, "birthdate", &sourceBirthdates, &birthdates); err != nil {
		birthdates = nil
	}
	return
}

// CleanVehicles clean all provided vehicles
// Call https://dadata.ru/api/v2/clean/vehicle
func (daData *DaData) CleanVehicles(sourceVehicles ...string) ([]Vehicle, error) {
	return daData.CleanVehiclesWithCtx(context.Background(), sourceVehicles...)
}

// CleanVehiclesWithCtx clean all provided vehicles
// Call https://dadata.ru/api/v2/clean/vehicle
func (daData *DaData) CleanVehiclesWithCtx(ctx context.Context, sourceVehicles ...string) (vehicles []Vehicle, err error) {
	if err = daData.sendCleanRequest(ctx, "vehicle", &sourceVehicles, &vehicles); err != nil {
		vehicles = nil
	}
	return
}

// CleanPassports clean all provided passports
// Call https://dadata.ru/api/v2/clean/passport
func (daData *DaData) CleanPassports(sourcePassports ...string) ([]Passport, error) {
	return daData.CleanPassportsWithCtx(context.Background(), sourcePassports...)
}

// CleanPassportsWithCtx clean all provided passports
// Call https://dadata.ru/api/v2/clean/passport
func (daData *DaData) CleanPassportsWithCtx(ctx context.Context, sourcePassports ...string) (passports []Passport, err error) {
	if err = daData.sendCleanRequest(ctx, "passport", &sourcePassports, &passports); err != nil {
		passports = nil
	}
	return
}
